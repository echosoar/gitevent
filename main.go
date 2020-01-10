package main

import (
	"fmt"
  "github.com/echosoar/gitevent/week"
)

func main() {
    week := new(week.WeekData)
    week.UserName = "echosoar"
		week.Run()
		typeMap := week.GetTypeMap()
		for lang, value := range typeMap {
			fmt.Printf("%s files:%d add:%d del:%d\n", lang, value.Files, value.Additions, value.Deletions)
		}
}