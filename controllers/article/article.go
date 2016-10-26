package article

import (
	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

//TODO
func (article *ArticleController) Get(){
	article.Data["name"] = "lkn"
	article.TplName = "article/add_article.tpl"
}
//TODO add
func (article *ArticleController) AddArticle(){
	article.Data["name"] = "lkn"
	article.TplName = "article/add_article.tpl"
}
//TODO list
func (article *ArticleController) ListArticle(){
	category_id  := article.Ctx.Input.Param(":category_id");
	pn := article.Ctx.Input.Param(":pn")
	article.Data["name"] = "lkn"
	article.Data["category_id"] = category_id
	article.Data["pn"] = pn

	article.TplName = "article/add_article.tpl"
}

//TODO search
func (article *ArticleController) SearchArticle(){
	article.Data["name"] = "lkn"
	article.TplName = "article/add_article.tpl"
}

// TODO del

func (article *ArticleController) DelArticle(){
	article.Data["name"] = "lkn"
	article.TplName = "article/add_article.tpl"
}

