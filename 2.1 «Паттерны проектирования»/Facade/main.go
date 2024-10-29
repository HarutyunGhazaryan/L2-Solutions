package main

import "fmt"

type RegistrationSystem struct{}

func (rs *RegistrationSystem) RegisterStudent(name, course string) {
	fmt.Println("Registering student:", name)
}

type CourseSystem struct{}

func (cs *CourseSystem) EnrollStudentInCourse(name, course string) {
	fmt.Println("Enrolling student", name, "in course:", course)
}

type NotificationSystem struct{}

func (ns *NotificationSystem) SendNotification(name string) {
	fmt.Println("Sending notification to:", name)
}

type OnlineSchoolFacade struct {
	registration *RegistrationSystem
	course       *CourseSystem
	notification *NotificationSystem
}

func NewOnlineSchoolFacade() *OnlineSchoolFacade {
	return &OnlineSchoolFacade{
		registration: &RegistrationSystem{},
		course:       &CourseSystem{},
		notification: &NotificationSystem{},
	}
}

func (os *OnlineSchoolFacade) RegisterStudentForCourse(name, course string) {
	os.registration.RegisterStudent(name, course)
	os.course.EnrollStudentInCourse(name, course)
	os.notification.SendNotification(name)
}

func main() {
	onlineSchoolFacade := NewOnlineSchoolFacade()
	onlineSchoolFacade.RegisterStudentForCourse("Name1", "Go")
}

// Паттерн Фасад используется здесь для предоставления упрощённого интерфейса для выполнения сложных операций
// регистрации студента, зачисления его на курс и отправки уведомлений. Вместо того, чтобы взаимодействовать
// с несколькими подсистемами напрямую, клиент может использовать единый интерфейс.
//
// Плюсы:
// - Упрощает сложные операции, скрывая детали работы подсистем за единой оболочкой.
// - Снижает сцепление между подсистемами и клиентом, что облегчает поддержку и делает систему более гибкой.
//
// Минусы:
// - Может добавить лишнюю сложность, если система слишком простая, создавая дополнительный уровень абстракции.
// - Фасад может стать слишком сложным, если будет пытаться скрыть слишком много функций, что приведёт к
// нарушению принципа единой ответственности.
