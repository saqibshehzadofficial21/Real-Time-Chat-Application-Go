package repository

import (
    "chat-app/internal/models"
    "gorm.io/gorm"
)

type FriendRepository interface {
    CreateRequest(req *models.FriendRequest) error
    GetRequestByID(id int) (*models.FriendRequest, error)
    UpdateStatus(id int, status models.RequestStatus) error
    GetPendingForUser(userID int) ([]models.FriendRequest, error)
    ExistsBetween(senderID, receiverID int) (bool, error)
    CreateConversation(isGroup bool) (*models.Conversation, error)
    AddParticipants(conversationID int, userIDs []int) error
}

type friendRepo struct {
    db *gorm.DB
}

func NewFriendRepository(db *gorm.DB) FriendRepository {
    return &friendRepo{db: db}
}

func (r *friendRepo) CreateRequest(req *models.FriendRequest) error {
    return r.db.Create(req).Error
}

func (r *friendRepo) GetRequestByID(id int) (*models.FriendRequest, error) {
    var req models.FriendRequest
    err := r.db.First(&req, id).Error
    if err != nil {
        return nil, err
    }
    return &req, nil
}

func (r *friendRepo) UpdateStatus(id int, status models.RequestStatus) error {
    return r.db.Model(&models.FriendRequest{}).Where("id = ?", id).Update("status", status).Error
}

func (r *friendRepo) GetPendingForUser(userID int) ([]models.FriendRequest, error) {
    var requests []models.FriendRequest
    err := r.db.Where("receiver_id = ? AND status = ?", userID, models.StatusPending).Find(&requests).Error
    return requests, err
}

func (r *friendRepo) ExistsBetween(senderID, receiverID int) (bool, error) {
    var count int64
    err := r.db.Model(&models.FriendRequest{}).
        Where("(sender_id = ? AND receiver_id = ?) OR (sender_id = ? AND receiver_id = ?)",
            senderID, receiverID, receiverID, senderID).
        Count(&count).Error
    return count > 0, err
}

func (r *friendRepo) CreateConversation(isGroup bool) (*models.Conversation, error) {
    conv := &models.Conversation{IsGroup: isGroup}
    err := r.db.Create(conv).Error
    return conv, err
}

func (r *friendRepo) AddParticipants(conversationID int, userIDs []int) error {
    var participants []models.ConversationParticipant
    for _, uid := range userIDs {
        participants = append(participants, models.ConversationParticipant{
            ConversationID: conversationID,
            UserID:         uid,
        })
    }
    return r.db.Create(&participants).Error
}