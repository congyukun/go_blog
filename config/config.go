package config

// import "github.com/spf13/viper"

import (
	"os"

	"github.com/BurntSushi/toml"
)

type conf struct {
	Viewer       Viewer
	SystemConfig SystemConfig
}

type Viewer struct {
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Logo        string   `json:"logo"`
	Navigation  []string `json:"navigations"`
	Bilibili    string   `json:"bilibili"`
	Avatar      string   `json:"avatar"`
	UserName    string   `json:"user_name"`
	UserDesc    string   `json:"user_desc"`
}

type SystemConfig struct {
	AppName        string `json:"app_name"`
	Version        string `json:"version"`
	CurrentDir     string
	CdnUrl         string `json:"cdn_url"`
	QiniuAccessKey string
	QiniuSecretKey string
	Valine         bool
	ValineAppId    string
	ValineAppKey   string
}

var Config *conf

func init() {
	Config = new(conf)
	Config.SystemConfig.AppName = "blog"
	Config.SystemConfig.Version = "1.0.0"
	Config.SystemConfig.CurrentDir, _ = os.Getwd()
	_, err := toml.DecodeFile("./config/conf.toml", &Config)
	if err != nil {
		panic(err)
	}
}
