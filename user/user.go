package user

import "github.com/febg/Climbtracker/gym"

// UserData represents the user climbing data from data base
type UserData struct {
	Name     string
	Email    string
	Password string
	UserID   string
	QrCode   string
}

type PrivateUser struct {
	FInfo []PublicUser
}

type FriendList struct {
	Friends []string
}

type PublicUser struct {
	Name     string
	Email    string
	Public   string
	Climbing gym.OverallData
	//TODO Friends/Related friends implementation
}

// NewCheckIn contains checkin information
type NewCheckIn struct {
	Level  string
	UserID string
}

func (pu *PublicUser) PublicUserCongif() {

}

func (privU *PrivateUser) Append(pubU *PublicUser) {
	//privU.FInfo = append(privU.FInfo, pubU)
}
