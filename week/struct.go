package week

type WeekData struct {
	typeMap map[string]TypeMap
	UserName string
	token string
}

type EventData struct {
	filename string
	additions int
	deletions int
	date string
}

type TypeMap struct {
	additions int
	deletions int
	events []EventData
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