package models

type Role = int

type UserModel struct {
	tableName struct{} `pg:"users"`

	ID       uint   `pg:",pk"`
	Username string `pg:",unique,notnull"`
	Password string `pg:",notnull"`
	Email    string `pg:",unique,notnull"`
	Roles    []int  `pg:",default:'1'"`
	Avatar   string `pg:",default:'default.jpg'"`
}

type User struct {
	Id       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Avatar   string `json:"avatar"`
	Password string `json:"password"`
	Roles    []int  `pg:",array" json:"roles"`
}

type Userfor struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Roles    []Role `json:"role"`
	Avatar   string `json:"avatar"`
}

type UserLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
