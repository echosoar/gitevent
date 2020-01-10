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
    week := new(week.WeekData)
    week.UserName = "echosoar"
		week.Run()
		typeMap := week.GetTypeMap()
		for lang, value := range typeMap {
			fmt.Printf("%s files:%d add:%d del:%d\n", lang, value.Files, value.Additions, value.Deletions)
		}
}
```

```shell
$ go run main.go

Yaml files:3 add:42 del:12
TypeScript files:88 add:845 del:1055
Json files:22 add:119 del:37
JavaScript files:7 add:70 del:52
unknow files:1 add:1 del:1
```
---
MIT @echosoar