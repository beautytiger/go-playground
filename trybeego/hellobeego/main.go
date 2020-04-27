package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"playground/trybeego/hellobeego/models"
	_ "playground/trybeego/hellobeego/routers"
)

func insertUser() {
	o := orm.NewOrm()
	user := models.User{}
	user.Name = "wang"
	id, err := o.Insert(&user)
	if err != nil {
		logs.Error("insert error")
		return
	}
	logs.Info("insert success, id=", id)
}

func queryUser() {
	o := orm.NewOrm()
	user := models.User{Id: 1}
	err := o.Read(&user)
	if err != nil {
		logs.Error("query is error")
	}
	logs.Info("read user info: ", user.Name)
}

func userUpdate() {
	o := orm.NewOrm()
	user := models.User{Id: 1, Name: "ming"}
	_, err := o.Update(&user)
	if err != nil {
		logs.Error("update user error")
	}
	logs.Info("update user failed")
}

func deleteUser() {
	o := orm.NewOrm()
	user := models.User{Id: 1}
	_, err := o.Delete(&user)
	if err != nil {
		logs.Error("delete error")
	}
	logs.Info("delete user success")
}

func insertOrder() {
	o := orm.NewOrm()
	order := models.UserOrder{}
	order.Order_data = "this is an order"
	user := models.User{Id: 1}
	order.User = &user
	id, err := o.Insert(&order)
	if err != nil {
		logs.Error("order insert error")
		return
	}
	logs.Info("order insert success, id=", id)
}

func orderQuery() {
	var orders []*models.UserOrder
	o := orm.NewOrm()
	qs := o.QueryTable("UserOrder")
	_, err := qs.Filter("user__id", 1).All(&orders)
	if err != nil {
		logs.Error("order query error")
	}
	for _, order := range orders {
		logs.Info(order.Order_data)
	}
}

func main() {
	//insertUser()
	//queryUser()
	//userUpdate()
	//deleteUser()
	//insertOrder()
	//orderQuery()
	beego.SetStaticPath("download1", "down")
	beego.Run()
}
