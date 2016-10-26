package inner_process

import (
	"github.com/astaxie/beego"
)

type ProcessController struct {
	beego.Controller
}

//TODO
func (process *ProcessController) Count (){
	article_id := process.Data["article_id"]
	article_url := process.Data["article_url"]

	process.Data["article_id"] = article_id

	process.Data["article_url"] = article_url
	process.TplName = "categoty/add_category.tpl"
}