package models

import "time"

type RequestStatus string

const (
    StatusPending  RequestStatus = "pending"
    StatusAccepted RequestStatus = "accepted"
    StatusRejected RequestStatus = "rejected"
)

type FriendRequest struct {
    ID         int           `json:"id" gorm:"primaryKey"`
    SenderID   int           `json:"sender_id" gorm:"not null"`
    ReceiverID int           `json:"receiver_id" gorm:"not null"`
    Status     RequestStatus `json:"status" gorm:"type:varchar(20);default:'pending'"`
    CreatedAt  time.Time     `json:"created_at"`
    UpdatedAt  time.Time     `json:"updated_at"`
}

func (FriendRequest) TableName() string {
    return "friend_requests"
}