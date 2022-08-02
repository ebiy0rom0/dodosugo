package main

import (
	"fmt"
	"os"
	"strings"

	"dodosugo/interface/controller"
	"dodosugo/interface/presenter"
	"dodosugo/usecase/interactor"
)

func main() {
	c := make(chan string)
	defer close(c)

	mp := presenter.NewMatcherPresenter(os.Stdout, c)
	m := interactor.NewMatcher(mp)
	mc := controller.NewMatcherContoller(m)

	if err := mc.Dodosugo(
		[]string{"ドド", "スコ"},
		strings.Repeat("ドド"+strings.Repeat("スコ", 3), 3),
		"ラブ注入♡",
	); err != nil {
		fmt.Printf("error: %s", err.Error())
	}
}
