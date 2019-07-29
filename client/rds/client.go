package rds

import (
	"fmt"

	"rootship.co.jp/merci/config"

	_ "github.com/go-sql-driver/mysql" // for mysql
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

// Client DBクライアント
type Client struct {
	DB *gorm.DB
}

// NewClient RDSへの接続を確立
func NewClient(conf *config.Config) (*Client, error) {
	rds := conf.Client.RDS
	me := new(Client)

	// コネクション生成
	db, err := gorm.Open(rds.Dialect, rds.Dsn)
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("Unable to connect databese. Dialect: %v DSN: %v", "Dialect", "DSN"))
	}

	// プール設定
	db.DB().SetMaxIdleConns(rds.IdleConn)
	db.DB().SetMaxOpenConns(rds.MaxConn)

	// ログモード設定
	db.LogMode(rds.Debug)

	me.DB = db
	return me, nil
}

// Close RDSのコネクションクローズ
func (c *Client) Close() {
	if c.DB != nil {
		c.DB.Close()
	}
}
