package tool

import (
	"bufio"
	"encoding/json"
	"os"
)

type Config struct {
	AppName  string         `json:"app_name"`
	AppMode  string         `json:"app_mode"`
	AppHost  string         `json:"app_host"`
	AppPort  string         `json:"app_port"`
	Sms      SmsConfig      `json:"sms"`
	DataBase DataBaseConfig `json:"database"`
}

type SmsConfig struct {
	SignName     string `json:"sign_name"`
	TemplateCode string `json:"template_code"`
	RegionId     string `json:"region_id"`
	AppKey       string `json:"app_key"`
	AppSecret    string `json:"app_secret"`
	Schema       string `json:"schema"`
}

type DataBaseConfig struct {
	Driver   string `json:"driver"`
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DBName   string `json:"db_name"`
	Charset  string `json:"charset"`
	ShowSQL  bool   `json:"show_sql"`
}

var _config *Config = nil

func GetConfig() *Config {
	return _config
}

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&_config); err != nil {
		return nil, err
	}
	return _config, nil
}
