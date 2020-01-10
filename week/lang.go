package week
import (
	"strings"
)

var langMap = map[string]string {
	"c": "C",
	"cpp": "C++",
	"css": "CSS",
	"go": "Go",
	"h": "C",
	"html": "HTML",
	"java": "Java",
	"js": "JavaScript",
	"jsx": "JavaScript",
	"json": "Json",
	"makefile": "Makefile",
	"md": "Markdown",
	"php": "PHP",
	"py": "Python",
	"sh": "Shell",
	"ts": "TypeScript",
	"tsx": "TypeScript",
	"yml": "Yaml",
}

func GetFileType(allFileName string) string {
	fileNameList := strings.Split(allFileName, "/")
	fileName := fileNameList[len(fileNameList) - 1]
	fileNameExtList := strings.Split(fileName, ".")
	fileNameExtStr := fileNameExtList[len(fileNameExtList) - 1]
	fileNameExt := langMap[strings.ToLower(fileNameExtStr)]
	if fileNameExt != "" {
		return fileNameExt
	}
	return "Other"
}