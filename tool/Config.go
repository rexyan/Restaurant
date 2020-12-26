package tool

import (
	"bufio"
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	AppName  string         `json:"app_name"`
	AppMode  string         `json:"app_mode"`
	AppHost  string         `json:"app_host"`
	AppPort  string         `json:"app_port"`
	Sms      SmsConfig      `json:"sms"`
	DataBase DataBaseConfig `json:"database"`
	Redis    RedisConfig    `json:"redis"`
}

type SmsConfig struct {
	TemplateID string `json:"template_id"`
	AppCode    string `json:"app_code"`
	URL        string `json:"url"`
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

type RedisConfig struct {
	Addr     string `json:"addr"`
	Password string `json:"password"`
	DB       int	`json:"db"`
	Port     string `json:"port"`
	Prefix   string `json:"prefix"`
}

var config *Config = nil

func GetConfig() *Config {
	return config
}

func ParseConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	decoder := json.NewDecoder(reader)
	if err := decoder.Decode(&config); err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return config, nil
}
