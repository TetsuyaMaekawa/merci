package merci

import (
	"net/http"

	"github.com/line/line-bot-sdk-go/linebot"
	"rootship.co.jp/merci/client/line"
	"rootship.co.jp/merci/client/rds"
	"rootship.co.jp/merci/util"
)

// Handler ...
type Handler struct {
	Rds  *rds.Client
	Line *line.Client
}

// DistributeEvent 受け付けたリクエストをハンドリング
func (h *Handler) DistributeEvent(r *http.Request) {
	events, err := h.Line.Bot.ParseRequest(r)
	if err == nil {
		for _, event := range events {
			switch event.Type {
			case linebot.EventTypeFollow:
				h.responseToFolloEvent(event)
			case linebot.EventTypeMessage:

				h.distributeMessage(event)
			default:
			}
		}
	} else {
	}
}

// distributeMessage メッセージによって処理を振り分ける
func (h *Handler) distributeMessage(event *linebot.Event) {
	switch message := event.Message.(type) {
	case *linebot.TextMessage:
		textMessage := message.Text
		splitMessage := util.SplitString(textMessage)
		if util.IsArg(splitMessage) {
			h.responseToMessageEvent(event, splitMessage)
		} else {
			h.responseErrMessage(event, "メッセージの送り方が違うみたい...\n送る人の氏名 メルシーメッセージ 送りたい人の氏名\n↑この形で送ってね！")
		}

	default:
		h.responseErrMessage(event, "画像やスタンプは読み込めないよ...\nそんなに高性能じゃないんだ....\n文字列だけ送ってくれ！\n使い方がわからない場合は「使い方」って入力してね！")
	}
}
