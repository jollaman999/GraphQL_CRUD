package types

type User struct {
	Id 		string `json: "id"`
	Name	string `json: "name"`
}

type Users struct {
	Users []User `json: "user"`
}