package main

import (
	"github.com/LgoLgo/cqupt-grabber/cqupt"
)

func main() {
	tool := cqupt.NewForSecXk()

	str3 := []string{
		"匹配词",
	}
	cookie := "你的cookie"
	loads := tool.Queryer.BlockSearch(cookie, str3)

	tool.Grabber.LoopRob(cookie, loads)
}
