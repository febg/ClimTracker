package api

import (
	"log"

	"github.com/febg/Climbtracker/data"
)

// Control is
type Control struct {
	Config   ControlConfig
	DataBase data.Data
}

// ControlConfig configures the settings of the server controller
type ControlConfig struct {
	LocalHost bool
}

// NewControl creates and initialzes a new server controller using conifg settings
func NewControl(config ControlConfig) (*Control, error) {
	c := Control{
		Config: config,
	}
	defer log.Printf("[LOG] Started server controller { [Localhost: %v] }", c.Config.LocalHost)
	var err error
	if config.LocalHost {
		c.DataBase, err = data.NewMySQL()
		if err != nil {
			log.Printf("[ERROR] Could not initialized data storage sytem %v", err)
		}
	}
	c.DataBase, err = data.NewLocalHost()
	if err != nil {
		log.Printf("[ERROR] Could not initialized data storage sytem %v", err)
	}
}
