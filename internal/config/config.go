package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

func getConfigFilePath ()(string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return  "", err
	}
	filePath := home + "/" + configFileName
	return filePath, nil
}

type Config struct {
	DbUrl string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func (c *Config) SetUser(user string) error {
	c.CurrentUserName = user
	file, err := getConfigFilePath()
	if err != nil {
		return err
	}
	data, err:= json.Marshal(c)
	if err != nil {
		return err
	}
	err = os.WriteFile(file, data, 0644)
	if err != nil {
		return err
	}
	return nil
}

func Read() (Config, error) {
	c := Config{}
	file, err := getConfigFilePath()
	if err != nil {
		return c, err
	}
	data, err := os.ReadFile(file)
	if err != nil {
		return c, err
	}
	err = json.Unmarshal(data, &c)
	if err != nil {
		return Config{}, err
	}
	return c, nil
}
