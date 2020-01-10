package main

import (
    "github.com/echosoar/gitevent/week"
)

func main() {
    var week *week.WeekData = new(week.WeekData)
    week.UserName = "echosoar"
    week.Run()
}