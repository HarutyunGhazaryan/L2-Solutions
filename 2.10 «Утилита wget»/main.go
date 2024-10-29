package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"sync"

	"github.com/PuerkitoBio/goquery"
)

// downloadFile скачивает файл по URL и сохраняет его в заданный путь
func downloadFile(fileURL, filePath string, wg *sync.WaitGroup) {
	defer wg.Done() // Уменьшаем счетчик горутин по завершении функции

	// Создаем файл
	out, err := os.Create(filePath)
	if err != nil {
		log.Printf("Failed to create file %s: %v", filePath, err)
		return
	}
	defer out.Close()

	// Выполняем HTTP-запрос
	resp, err := http.Get(fileURL)
	if err != nil {
		log.Printf("Failed to get URL %s: %v", fileURL, err)
		return
	}
	defer resp.Body.Close()

	// Проверяем статус код
	if resp.StatusCode != http.StatusOK {
		log.Printf("Failed to download file: %s", resp.Status)
		return
	}

	// Копируем данные из ответа в файл
	if _, err = io.Copy(out, resp.Body); err != nil {
		log.Printf("Failed to write to file %s: %v", filePath, err)
	}
}

// downloadPage скачивает страницу по URL и парсит ее
func downloadPage(pageURL, basePath string) error {
	// Загружаем страницу
	resp, err := http.Get(pageURL)
	if err != nil {
		return fmt.Errorf("failed to get URL %s: %w", pageURL, err)
	}
	defer resp.Body.Close()

	// Проверяем статус код
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download page: %s", resp.Status)
	}

	// Создаем директорию для сохранения страницы
	pageName := filepath.Base(pageURL)
	pagePath := filepath.Join(basePath, pageName)
	if err := os.MkdirAll(pagePath, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory %s: %w", pagePath, err)
	}

	// Загружаем HTML-документ
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to parse HTML document: %w", err)
	}

	var wg sync.WaitGroup // Группа ожидания для горутин

	// Обрабатываем все ссылки на странице
	doc.Find("a").Each(func(i int, s *goquery.Selection) {
		link, exists := s.Attr("href")
		if !exists {
			return // Пропускаем, если атрибут отсутствует
		}
		absLink := resolveLink(pageURL, link)
		fmt.Println("Found link:", absLink)
		wg.Add(1)
		go downloadFile(absLink, filepath.Join(pagePath, filepath.Base(link)), &wg)
	})

	// Обрабатываем изображения
	doc.Find("img").Each(func(i int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if exists {
			absSrc := resolveLink(pageURL, src)
			fmt.Println("Found image:", absSrc)
			wg.Add(1)
			go downloadFile(absSrc, filepath.Join(pagePath, filepath.Base(src)), &wg)
		}
	})

	// Обрабатываем CSS
	doc.Find("link[rel='stylesheet']").Each(func(i int, s *goquery.Selection) {
		href, exists := s.Attr("href")
		if exists {
			absHref := resolveLink(pageURL, href)
			fmt.Println("Found stylesheet:", absHref)
			wg.Add(1)
			go downloadFile(absHref, filepath.Join(pagePath, filepath.Base(href)), &wg)
		}
	})

	// Обрабатываем JS
	doc.Find("script").Each(func(i int, s *goquery.Selection) {
		src, exists := s.Attr("src")
		if exists {
			absSrc := resolveLink(pageURL, src)
			fmt.Println("Found script:", absSrc)
			wg.Add(1)
			go downloadFile(absSrc, filepath.Join(pagePath, filepath.Base(src)), &wg)
		}
	})

	wg.Wait() // Ждем завершения всех горутин
	return nil
}

// resolveLink преобразует относительную ссылку в абсолютную
func resolveLink(base, link string) string {
	if _, err := url.Parse(link); err == nil {
		return link // Возвращаем абсолютный URL без изменений
	}

	baseURL, err := url.Parse(base)
	if err != nil {
		log.Println("Error parsing base URL:", err)
		return link
	}

	if link[0] == '/' {
		return baseURL.Scheme + "://" + baseURL.Host + link
	}

	return baseURL.ResolveReference(&url.URL{Path: link}).String()
}

func main() {
	// URL страницы для загрузки
	urlToDownload := "http://example.com"
	basePath := "./downloaded"

	if err := downloadPage(urlToDownload, basePath); err != nil {
		log.Fatal(err)
	}
}
