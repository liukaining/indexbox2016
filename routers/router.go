package routers

import (
	"indexbox/controllers"
	"indexbox/controllers/article"
	"indexbox/controllers/category"
	"indexbox/controllers/inner_process"
	"indexbox/controllers/record"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/indexbox", &controllers.MainController{})
	beego.Router("/", &controllers.MainController{})

	// article
	beego.Router("/add_article",&article.ArticleController{},"post:AddArticle")
	beego.Router("/list_article/:category_id:int",&article.ArticleController{},"get:ListArticle")
	beego.Router("/search/:query:string",&article.ArticleController{},"post:SearchArticle")
	beego.Router("/del_article/:article_id:int",&article.ArticleController{},"get:DelArticle")

	// category
	beego.Router("/add_category",&category.CategoryController{},"post:AddCategory")
	beego.Router("/list_category_all",&category.CategoryController{},"get:ListAllCategory")
	beego.Router("/list_category_sub/:category_id:int",&category.CategoryController{},"get:ListSubCategory")
	beego.Router("/list_category_level/:level_id:int",&category.CategoryController{},"get:ListLevelCategory")
	beego.Router("/del_category/:category_id:int",&category.CategoryController{},"get:DelCategory")

	// record
	beego.Router("/add_record",&record.RecordController{},"post:AddRecord")
	beego.Router("/update_record",&record.RecordController{},"post:UpdateRecord")
	beego.Router("/get_record/:article_id:int",&record.RecordController{},"get:GetRecord")
	beego.Router("/get_record_all",&record.RecordController{},"get:GetRecordAll")

	//inner_process
	beego.Router("/count/:article_id:int/:article_url:string",&inner_process.ProcessController{},"get:Count")





}
