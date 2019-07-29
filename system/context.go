package system

import (
	"rootship.co.jp/merci/client/line"
	"rootship.co.jp/merci/client/rds"
	"rootship.co.jp/merci/config"
)

// Context コンテキスト
type Context struct {
	Conf *config.Config
	Rds  *rds.Client
	Line *line.Client
}

// NewContext コンテキストの生成
func NewContext(confPath string) (*Context, error) {
	me := new(Context)

	conf, err := config.NewConfig(confPath)
	if err != nil {
		return nil, err
	}

	rds, err := rds.NewClient(conf)
	if err != nil {
		return nil, err
	}

	Bot, err := line.NewClient(conf)
	if err != nil {
		return nil, err
	}

	me.Conf = conf
	me.Rds = rds
	me.Line = Bot

	return me, nil
}
