package main

import (
	_ "embed"
	"strings"
)

//go:embed bot_information.md
var BotInformation string

func UpdateTest(str string, user string) string {
	if strings.HasPrefix(str, "✅ ") {
		s, _ := strings.CutPrefix(str, "✅ ")
		return s
	} else {
		return "✅ <del>" + str + "</del>\nDone by: <i>" + user + "</i>"
	}
}

func JoinMessagesText(msgs any) string {
	return ""
}
