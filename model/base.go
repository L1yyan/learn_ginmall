package model
//分页查询
type BasePage struct {
	PageNum int `form:"pageNum"`
	PageSize int `form:"pageSize"`
}