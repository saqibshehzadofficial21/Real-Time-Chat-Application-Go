package friend

import "errors"

// RejectRequest friend request ko reject kar deta hai
func (s *friendService) RejectRequest(requestID, userID int) error {
    req, err := s.repo.GetRequestByID(requestID)
    if err != nil {
        return errors.New("friend request not found")
    }

    if req.ReceiverID != userID {
        return errors.New("you are not authorized to reject this request")
    }

    return s.repo.UpdateStatus(requestID, "rejected")
}