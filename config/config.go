package config

import (
	"fmt"

	"github.com/spf13/viper"
)

// Config アプリケーション設定情報
type Config struct {
	AppName  string `mapstructure:"app_name"`
	LogLevel string `mapstructure:"log_level"`
	LogPath  string `mapstructure:"log_path"`
	Client   *Client
}

// Client クライアント設定情報
type Client struct {
	RDS  *RDS
	LINE LINE
}

// RDS データベース設定情報
type RDS struct {
	Dialect  string `mapstructure:"dialect"`
	Dsn      string `mapstructure:"dsn"`
	IdleConn int    `mapstructure:"idle_conn"`
	MaxConn  int    `mapstructure:"max_conn"`
	Debug    bool   `mapstructure:"debug"`
}

// LINE line設定情報
type LINE struct {
	ChannelSecret string `mapstructure:"channelSecret"`
	AccessToken   string `mapstructure:"accessToken"`
}

// NewConfig 各種設定読み込み
func NewConfig(path string) (*Config, error) {
	me := new(Config)

	viper.SetConfigFile(path)
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("Unable to read config. %v", err)
	}

	if err := viper.Unmarshal(&me); err != nil {
		return nil, fmt.Errorf("Unable to unmarshal. %v", err)
	}

	return me, nil
}
