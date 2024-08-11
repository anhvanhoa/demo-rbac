package tmpl

type itemCategory struct {
	Id   uint8
	Name string
	Link string
}

var HeaderData [4]itemCategory = [4]itemCategory{
	{Id: 1, Name: "Todo Groups", Link: "/todo-groups"},
	{Id: 2, Name: "Account", Link: "/account"},
}
