package models

import _ "github.com/go-sql-driver/mysql"

// Model Struct
type User struct {
	Id    int          //default primary key, incr
	Name  string       `orm:"size(100)"`
	Order []*UserOrder `orm:"reverse(many)"`
}

type UserOrder struct {
	Id         int
	Order_data string `orm:size(100)`
	User       *User  `orm:ref(fk)`
}

func init() {
	// set default database
	//orm.RegisterDataBase("default", "mysql", "root:password@tcp(127.0.0.1:3306)/class27?charset=utf8", 30)

	// register model
	//orm.RegisterModel(new(User))

	// create table
	//orm.RunSyncdb("default", false, true)
}
