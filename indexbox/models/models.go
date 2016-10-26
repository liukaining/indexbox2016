package models

import (
	"github.com/astaxie/beego/orm"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type ArticleInfo struct {
	Id int
	Abstract string
	Title string
	Url string
	InsertTime int
	InsertUser string
	IsDel int
	Ext string
}

type CategoryBaseInfo struct {
	Id int
	Level int
	Name string
	FatherCategory int
	InsertTime int
	InsertUser string
}

type CategoryArticleRecord struct {
	Id int
	Article *ArticleInfo `orm:"rel(fk)"`
	Category *CategoryBaseInfo `orm:"rel(fk)"`
	InsertTime int
	IsEffective int

}
