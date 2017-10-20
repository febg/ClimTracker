package api

import "github.com/febg/Climbtracker/data"

// Control is
type Control struct {
	DataBase data.MySQL
}

type ControlConfig struct {
	LocalHost bool
}

func NewControl(config ControlConfig) {

}
