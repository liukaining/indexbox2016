package record

import (
	"github.com/astaxie/beego"
)

type RecordController struct {
	beego.Controller
}

// TODO
func (record *RecordController) AddRecord()  {
	record.Data["record"] = "record"
	record.TplName = "record/add_record"
}

// TODO
func (record *RecordController) UpdateRecord()  {
	record.Data["record"] = "record"
	record.TplName = "record/update_record"
}

// TODO
func (record *RecordController) GetRecord()  {
	record.Data["record"] = "record"
	record.TplName = "record/article_record"
}

// TODO
func (record *RecordController) GetRecordAll()  {
	record.Data["record"] = "record"
	record.TplName = "record/all_record"
}
