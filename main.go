package main

import (
	"3x-ui-bot/tools"
)

func init() {
	requirements := tools.RequirementsValue
	for _, requirement := range requirements.Servers {
		err := tools.Login(&requirement)
		if err != nil {
			panic(err)
		}
	}
}
func main() {
	tools.Start()
}
