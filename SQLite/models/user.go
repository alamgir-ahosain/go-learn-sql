package models

type User struct {
	ID   int     `json:"Id"`
	SID  string  `json:"sid"`
	Name string  `json:"name"`
	CGPA float32 `json:"cgpa"`
}
