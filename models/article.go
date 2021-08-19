package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"` //声明为index字段
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

func ExistArticleByID(id int) bool {
	var article Article
	Db.Select("id").Where("id = ?", id).First(&article)

	return article.ID > 0
}

func GetArticleTotal(maps interface{}) (count int) {

	Db.Model(&Article{}).Where(maps).Count(&count)

	return
}

func GetArticles(pageNum, pageSize int, maps interface{}) (aretices []Article) {
	Db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&aretices)

	return
}

func GetArticle(id int) (article Article) {
	Db.Where("id = ?", id).First(&article)
	Db.Model(&article).Related(&article.Tag) //关联到tag, tagId就是外键 gorm会通过类名+ID的方式去查找这两个类之前的关联关系,结构体中内嵌的Tag,通过Related进行关联查询

	return
}

func EditArticle(id int, data interface{}) bool {
	Db.Model(&Article{}).Where("id = ?", id).Update(data)

	return true
}

func AddArticle(data map[string]interface{}) bool {
	Db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})

	return true
}

func DeleteArticle(id int) bool {
	Db.Where("id = ?", id).Delete(&Article{})

	return true
}

func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix()) //当前时间戳

	return nil
}

func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}
