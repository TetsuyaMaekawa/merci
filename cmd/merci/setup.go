package main

import (
	"net/http"

	"rootship.co.jp/merci/service/merci"

	"rootship.co.jp/merci/system"

	"github.com/zenazn/goji"
	"github.com/zenazn/goji/web"
)

func setup() {
	const confPath = "../../merci.toml"

	context, err := system.NewContext(confPath)
	if err != nil {
		return
	}

	handler := merci.Handler{Rds: context.Rds, Line: context.Line}

	// ルーティング
	goji.Post("/merci", func(c web.C, w http.ResponseWriter, r *http.Request) {
		handler.DistributeEvent(r)
	})
}
