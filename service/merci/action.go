package merci

import (
	"github.com/line/line-bot-sdk-go/linebot"
)

// responseToFolloEvent フォローイベントに対してレスポンスする
func (h *Handler) responseToFolloEvent(event *linebot.Event) {
	if _, err := h.Line.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("友達追加ありがとう！たくさんのメルシーを仲間に伝えよう！\n使い方を知りたい時は「使い方」とメッセージを送ってみてね！")).Do(); err != nil {
		h.responseErrMessage(event, "エラーが発生したよ！詳細は和知か前川に確認してね！")
		// logにエラー内容を出力する
	}
}

func (h *Handler) responseToMessageEvent(event *linebot.Event, splitMessage []string) {
	if len(splitMessage) == 1 {
		if splitMessage[0] == "使い方" {
			if _, err := h.Line.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("使い方を教えるね！\n送る人の氏名 メルシーメッセージ 送りたい人の氏名\n↑この形で送ってね！")).Do(); err != nil {
				h.responseErrMessage(event, "エラーが発生したよ！詳細は和知か前川に確認してね！")
				// logにエラー内容を出力する
			}
		}

	} else {
		// 送りたい相手の氏名が正しいかチェック間違っている場合は一覧を出力する

		if _, err := h.Line.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("メルシーを送ってくれてありがとう！あなたのメルシーの気持ちが相手届けられるよ！\nまたメルシー書いてね！")).Do(); err != nil {
			h.responseErrMessage(event, "エラーが発生したよ！詳細は和知か前川に確認してね！")
			// logにエラー内容を出力する
		}
	}
}

func (h *Handler) responseErrMessage(event *linebot.Event, message string) {
	if _, err := h.Line.Bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message)).Do(); err != nil {
		// logにエラーを出力する
		// ここでもエラーが発生した場合は和知と前川にエラーが発生した旨をlineする
	}
}
