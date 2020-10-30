package config

import (
	"errors"
	"fmt"
	"os"
	"strconv"

	"github.com/Moletastic/utem-gsp/utils"
)

type DBConfig struct {
	Engine  string `json:"engine"`
	Name    string `json:"name"`
	Host    string `json:"host"`
	Port    uint   `json:"port"`
	User    string `json:"user"`
	Pass    string `json:"pass"`
	Refresh bool   `json:"refresh"`
	LogMode bool   `json:"log_mode"`
}

type GSPConfig struct {
	Port     uint      `json:"port"`
	Pop      bool      `json:"pop"`
	DBConfig *DBConfig `json:"db_config"`
}

func (c *GSPConfig) GetAddress() string {
	return fmt.Sprintf(":%d", c.Port)
}

func (c *GSPConfig) ToString() string {
	return utils.Pretty(c)
}

func GetConfig() (*GSPConfig, error) {
	config := new(GSPConfig)
	config.DBConfig = new(DBConfig)
	pop, err := strconv.ParseBool(os.Getenv("POP"))
	if err != nil {
		fmt.Println("Error reading POP env variable. Setting variable as false")
		config.Pop = false
	} else {
		config.Pop = pop
	}
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		fmt.Println("Error reading PORT env variable. Ignoring")
	} else {
		config.Port = uint(port)
	}
	dbport, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		fmt.Println("Error reading DB_PORT env variable. Ignoring")
	} else {
		config.DBConfig.Port = uint(dbport)
	}
	config.DBConfig.Engine = os.Getenv("DB_ENGINE")
	config.DBConfig.Name = os.Getenv("DB_NAME")
	config.DBConfig.User = os.Getenv("DB_USER")
	config.DBConfig.Pass = os.Getenv("DB_PASS")
	config.DBConfig.Host = os.Getenv("DB_HOST")
	refresh, err := strconv.ParseBool(os.Getenv("DB_REFRESH"))
	if err != nil {
		fmt.Println("Error reading DB_REFRESH env variable. Ignoring")
	} else {
		config.DBConfig.Refresh = refresh
	}
	logmode, err := strconv.ParseBool(os.Getenv("DB_LOG"))
	if err != nil {
		fmt.Println("Error reading DB_LOG env variable. Ignoring")
	} else {
		config.DBConfig.LogMode = logmode
	}
	if config.DBConfig.Engine == "" || config.DBConfig.Name == "" {
		return nil, errors.New("DB_NAME and DB_ENGINE env variables needs to be setted")
	}
	return config, nil
}
