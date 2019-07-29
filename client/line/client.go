package line

import (
	"fmt"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/pkg/errors"
	"rootship.co.jp/merci/config"
)

// Client lineクライアント
type Client struct {
	Bot *linebot.Client
}

// NewClient lineクライアント生成
func NewClient(conf *config.Config) (*Client, error) {
	line := conf.Client.LINE

	me := new(Client)

	bot, err := linebot.New(line.ChannelSecret, line.AccessToken)
	if err != nil {
		return nil, errors.Wrapf(err, fmt.Sprintf("Unable to connect line. ChannelSecret: %v AccessToken: %v", line.ChannelSecret, line.AccessToken))
	}

	me.Bot = bot

	return me, nil
}
