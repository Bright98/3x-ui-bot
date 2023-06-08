package main

import (
	"3x-ui-bot/tools"
)

func init() {
	err := tools.Login()
	if err != nil {
		panic(err)
	}
}
func main() {
	tools.Start()
}
