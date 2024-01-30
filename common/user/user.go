package user

type User struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Pwd     string `json:"pwd"`
	Address string `json:"address"`
}
