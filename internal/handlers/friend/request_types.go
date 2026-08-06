package friend

// sendRequestBody friend request bhejte waqt client se aane wala JSON body
type sendRequestBody struct {
    ReceiverID int `json:"receiver_id" binding:"required"`
}