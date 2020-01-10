# gitevent

分析过去一周的时间中某用户在Github上提交的代码

包括代码类型(语言)，每种类型提交了多少个文件、增加了多少行代码、删除了多少行代码。

## Usage

```shell
$ go get github.com/echosoar/gitevent/week
```

```golang
package main

import "github.com/echosoar/gitevent/week"

func main() {
    var week *week.WeekData = new(week.WeekData)
    week.UserName = "echosoar"
    week.Run()
}
```

---
MIT @echosoar