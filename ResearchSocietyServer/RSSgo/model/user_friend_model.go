package model

type UserFriend struct {
	UserId            uint            `gorm:"not null" json:"userId"`
	FriendId          uint            `gorm:"not null" json:"friendId"`
	State             bool            `json:"state"`
	UserInformation   UserInformation `gorm:"ForeignKey:UserId"`
	FriendInformation UserInformation `gorm:"ForeignKey:FriendId"`
}
