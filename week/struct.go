package week

type WeekData struct {
	typeMap map[string]TypeInfo
	UserName string
	token string
}

type TypeInfo struct {
	Files int
	Additions int
	Deletions int
}

type OriginEventData struct {
	Type string
	Payload struct {	
		Commits [] struct {
			Url string 
		}
	}
}

type OriginCommitsJson struct {
	Files [] struct {
		Filename string
		Additions int
		Deletions int
	}
}