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
	Friends []string
}

type PublicUser struct {
	Name     string
	Email    string
	Qrcode   string
	Climbing gym.PublicClimbingData
	//TODO Friends/Related friends implementation
}

type PrivateUserConfig struct {
	Public string
}

type PublicUserConifg struct {
}

// NewCheckIn contains checkin information
type NewCheckIn struct {
	Level  string
	UserID string
}

func (pu *PublicUser) PublicUserCongif() {

}
