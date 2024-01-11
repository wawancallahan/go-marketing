package util

type ResultNone struct{}

type ResultList struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
