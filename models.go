package RestAPI_go

type Data struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var db []Data
