package object // import "github.com/severecloud/vksdk/object"

type databaseCity struct {
}

type databaseFaculty struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type databaseRegion struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type databaseSchool struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

type databaseUniversity struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}