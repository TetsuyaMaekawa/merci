package line

import (
	"testing"

	"rootship.co.jp/merci/config"
)

func TestNewClient(t *testing.T) {
	conf := config.Config{
		Client: &config.Client{
			LINE: config.LINE{
				AccessToken:   "testAccessToken",
				ChannelSecret: "testChannelSecret",
			},
		},
	}

	line, err := NewClient(&conf)

	if err != nil {
		t.Errorf("Test failed. err must be nil. Detail: %v", err)
	}

	if line == nil {
		t.Error("Test failed. line must not be nil.")
	}
}
