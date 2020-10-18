// 自动生成模板BusArticle
package model

import (
	"gorm.io/gorm"
)

// 如果含有time.Time 请自行import time包
type BusArticle struct {
      gorm.Model
      Title  string `json:"title" form:"title" gorm:"column:title;comment:文章标题;type:varchar(100);size:100;"`
      Desc  string `json:"desc" form:"desc" gorm:"column:desc;comment:文章描述;type:varchar(200);size:200;"`
      Author  int `json:"author" form:"author" gorm:"column:author;comment:作者ID;type:int(10);size:10;"`
      AuthorName  string `json:"authorName" form:"authorName" gorm:"column:authorName;comment:作者名字;type:varchar(100);size:100;"`
      Content  string `json:"content" form:"content" gorm:"column:content;comment:文章内容;type:text;"`
      Tag  string `json:"tag" form:"tag" gorm:"column:tag;comment:文章标签;type:varchar(100);size:100;"`
      TagNames  string `json:"tagNames" form:"tagNames" gorm:"column:tagNames;comment:标签回显;type:varchar(100);size:100;"` 
}


func (BusArticle) TableName() string {
  return "bus_article"
}
