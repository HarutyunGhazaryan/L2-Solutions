package main

import (
	"fmt"
	"reflect"
	"time"
)

// orReflect объединяет один или несколько done-каналов в один выходной канал.
// Когда любой из входящих каналов закрывается, выходной канал также закрывается.
// Эта реализация использует пакет reflect для динамического выбора канала.
// Если не переданы каналы, выходной канал будет немедленно закрыт.
var orReflect func(channels ...<-chan interface{}) <-chan interface{}

func init() {
	orReflect = func(channels ...<-chan interface{}) <-chan interface{} {
		done := make(chan interface{})

		go func() {
			defer close(done)

			switch len(channels) {
			case 0:
				return
			case 1:
				<-channels[0]
				return
			}

			selectCases := make([]reflect.SelectCase, len(channels))
			for i, ch := range channels {
				selectCases[i] = reflect.SelectCase{
					Dir:  reflect.SelectRecv,
					Chan: reflect.ValueOf(ch),
				}
			}

			_, _, _ = reflect.Select(selectCases)
		}()

		return done
	}
}

// orSimple объединяет один или несколько done-каналов в один выходной канал.
// Когда любой из входящих каналов закрывается, выходной канал также закрывается.
// Эта реализация использует горутины и оператор select для обработки завершения каналов.
// Если не переданы каналы, выходной канал будет немедленно закрыт.
var orSimple func(channels ...<-chan interface{}) <-chan interface{}

func init() {
	orSimple = func(channels ...<-chan interface{}) <-chan interface{} {
		done := make(chan interface{})

		go func() {
			defer close(done)

			for _, ch := range channels {
				go func(c <-chan interface{}) {
					<-c
					done <- struct{}{}
				}(ch)
			}
		}()

		return done
	}
}

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func main() {
	start := time.Now()

	<-orReflect(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Println(time.Since(start))

	start = time.Now()

	<-orSimple(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Println(time.Since(start))
}
