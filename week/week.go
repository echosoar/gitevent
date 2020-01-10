package week
import (
	"strconv"
	"net/http"
	"encoding/json"
	"io/ioutil"
)

func (week *WeekData) Run() {
	week.typeMap = make(map[string]TypeInfo)
	week.getAllEvent()
}

func (week *WeekData) getAllEvent() {
	week.getEventByPage(1);
}

func (week *WeekData) getEventByPage(page int) {
	url := "https://api.github.com/users/" + week.UserName + "/events?page=" + strconv.Itoa(page)
	body := getGithub(url)
	if body == nil {
		return
	}
	var arr []OriginEventData
	json.Unmarshal(body, &arr)
	
	for _, originEvent := range arr {
		if originEvent.Type == "PushEvent" {
			week.getEventDataList(originEvent.Payload.Commits[0].Url)
		}
	}
	
}

func getGithub(url string) []byte {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func (week *WeekData) getEventDataList(url string) {
	body := getGithub(url)
	var cmtJson OriginCommitsJson
	json.Unmarshal(body, &cmtJson)
	for _, file := range cmtJson.Files {
		fileType := GetFileType(file.Filename);
		tmp := week.typeMap[fileType]
		tmp.Additions += file.Additions
		tmp.Deletions += file.Deletions
		tmp.Files ++
		week.typeMap[fileType] = tmp
	}
}

func (week *WeekData) GetTypeMap() map[string]TypeInfo {
	return week.typeMap;
}