package connection

type EmailJob_MessPayload struct {
	Email   string
	Content string
}

func GetEmailChan() chan EmailJob_MessPayload {
	return emailChan
}
