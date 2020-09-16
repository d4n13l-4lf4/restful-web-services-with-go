package model

type Repo struct {
	ID int `json:"id"`
	Name string `json:"name"`
	FullName string `json:"full_name"`
	Forks int `json:"forks"`
	Private bool `json:"private"`
}

type File struct {
	Content string `json:"content"`
}

type Gist struct {
	Description string    `json:"description"`
	Public bool           `json:"public"`
	Files map[string]File `json:"files"`
}