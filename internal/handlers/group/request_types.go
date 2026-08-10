package group

// createGroupRequest naya group banate waqt client se aane wala JSON body
type createGroupRequest struct {
    Name      string `json:"name" binding:"required"`
    MemberIDs []int  `json:"member_ids" binding:"required"`
}