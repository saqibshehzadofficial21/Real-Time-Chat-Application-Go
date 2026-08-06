package message

// sendMessageRequest message bhejte waqt client se aane wala JSON body
type sendMessageRequest struct {
    ConversationID int    `json:"conversation_id" binding:"required"`
    Content        string `json:"content" binding:"required"`
}