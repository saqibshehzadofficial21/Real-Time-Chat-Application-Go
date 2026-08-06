package message


import msgsvc "chat-app/internal/service/message"

// MessageHandler HTTP requests ko MessageService tak forward karta hai
type MessageHandler struct {
    service msgsvc.MessageService
}

// NewMessageHandler ek naya MessageHandler banata hai
func NewMessageHandler(s msgsvc.MessageService) *MessageHandler {
    return &MessageHandler{service: s}
}