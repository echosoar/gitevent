package week
import (
	"strings"
)

var langMap = map[string]string {
	"ts": "TypeScript",
	"js": "JavaScript",
	"json": "Json",
	"yml": "Yaml",
	"go": "Golang",
}

func GetFileType(allFileName string) string {
	fileNameList := strings.Split(allFileName, "/")
	fileName := fileNameList[len(fileNameList) - 1]
	fileNameExtList := strings.Split(fileName, ".")
	fileNameExt := langMap[fileNameExtList[len(fileNameExtList) - 1]]
	if fileNameExt != "" {
		return fileNameExt
	}
	return "unknow"
}