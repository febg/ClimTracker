package api

import (
	"database/sql"
	"log"

	"github.com/jasonlvhit/gocron"

	"github.com/febg/Climbtracker/data"
	//"../data"
)

// Control is
type Control struct {
	Config   ControlConfig
	DataBase *sql.DB
	Cache    *data.CachedUsers
}

// ControlConfig configures the settings of the server controller
type ControlConfig struct {
	LocalMySQL bool
}

// NewControl creates and initialzes a new server controller using conifg settings
func NewControl(config ControlConfig) (*Control, error) {
	c := Control{
		Config: config,
	}
	log.Printf("[LOG] Starting server controller... { [LocalMySQL: %v] }", c.Config.LocalMySQL)
	defer log.Printf("[LOG] Started server controller")
	c.Cache = data.InitializeCache()
	log.Printf("[LOG] Created local cache on heap")
	var err error
	if !config.LocalMySQL {
		c.DataBase, err = data.NewMySQL()
		if err != nil {
			log.Printf("[FATAL] Could not initialized data storage sytem %v", err)
			return nil, err
		}
		log.Printf("[LOG] Stablished Connection to remote MySQL server")
		return &c, nil
	}
	c.DataBase, err = data.NewLocalMySQL()
	if err != nil {
		log.Printf("[FATAL] Could not initialized data storage sytem %v", err)
	}
	log.Printf("[LOG] Stablished Connection to local MySQL server")

	return &c, nil
}

// InitCron Initializes cron to schedulle tasks
func initCron(time int, t string) error {
	var task interface{} = t
	log.Printf("[LOG] Initialized cron for task: %v, every %v minutes", task, time)
	gocron.Every(2).Seconds().Do(tasktest)
	<-gocron.Start()
	return nil
}

//StartUpdateTimer sdfhdfg dfgh  dfgh
func UpdateData() {
	log.Printf("[LOG] Update Timmer Initialized")
	return
}

func tasktest() {
	log.Printf("Cron")
}

func test() {
	log.Printf("test")
}
