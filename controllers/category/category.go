package category

import (
	"github.com/astaxie/beego"
)

type CategoryController struct {
	beego.Controller
}

//TODO
func (category *CategoryController) AddCategory (){
	category.Data["category"] = "lll"
	category.TplName = "categoty/add_category.tpl"
}

//TODO
func (category *CategoryController) DelCategory (){
	category.Data["category"] = "lll"
	category.TplName = "categoty/del_category.tpl"
}

//TODO
func (category *CategoryController) ListAllCategory (){
	category.Data["category"] = "lll"
	category.TplName = "categoty/list_all_category.tpl"
}

//TODO
func (category *CategoryController) ListSubCategory (){
	category.Data["category"] = "lll"
	category.TplName = "categoty/list_sub_category.tpl"
}

//TODO
func (category *CategoryController) ListLevelCategory (){
	category.Data["category"] = "lll"
	category.TplName = "categoty/list_level_category.tpl"
}

