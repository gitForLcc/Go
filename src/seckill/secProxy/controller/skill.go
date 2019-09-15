package controller

import (
	"github.com/astaxie/beego"
)

type SkillController struct {
	beego.Controller
}

func (p *SkillController) SecKill() {
	p.Data["json"] = "seckill"
	p.ServeJSON()
}

func (p *SkillController) SecInfo() {
	p.Data["json"] = "secinfo"
	p.ServeJSON()
}
