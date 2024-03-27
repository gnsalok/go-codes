package main

type Notification interface {
	GetRecipient() string
	GetContent() string
}

type Email struct {
	To      string
	Subject string
	Body    string
}

type SMS struct {
	To      string
	Message string
}

func (e Email) GetRecipient() string {
	return e.To
}

func (e Email) GetContent() string {
	return e.Subject + "\n" + e.Body
}

func (s SMS) GetRecipient() string {
	return s.To
}

func (s SMS) GetContent() string {
	return s.Message
}

func SendNotification(notification Notification) {
	// Implement generic notification sending logic based on interface methods
	notification.GetRecipient()
	notification.GetContent()
}

func main() {

	email := &Email{
		To:      "gnsalk@gmail.com",
		Subject: "Hello",
		Body:    "Hi there",
	}

	sms := &SMS{
		To:      "099999",
		Message: "Hey Alok",
	}

	SendNotification(email)
	SendNotification(sms)

}
