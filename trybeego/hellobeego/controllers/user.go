package controllers

import (
	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	c.Ctx.WriteString("get from pure get")
}

func (c *UserController) Post() {
	c.Ctx.WriteString("This is a post data")
}

func (c *UserController) GetInfo() {
	id := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("get info data, id = "+id)
}

func (c *UserController) GetNum() {
	id := c.Ctx.Input.Param(":id")
	c.Ctx.WriteString("get number, id = "+id)
}

func (c *UserController) GetFile() {
	splat := c.Ctx.Input.Param(":splat")
	path := c.Ctx.Input.Param(":path")
	ext := c.Ctx.Input.Param(":ext")
	c.Ctx.WriteString("get file, path = "+path+"; ext= "+ext+"; splat= "+splat)
}