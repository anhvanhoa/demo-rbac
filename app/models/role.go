package models

type RoleModel struct {
	tableName struct{} `pg:"roles"`
	Id        int      `pg:",pk"`
	Name      string
}

type RoleItem struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
