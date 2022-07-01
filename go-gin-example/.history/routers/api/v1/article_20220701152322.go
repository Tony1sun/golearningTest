package v1

import (
	"go-gin-example/models"
	"go-gin-example/pkg/e"
	"log"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// 获取单个文章
func GetArticle(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	code := e.INVALID_PARAMS
	var data interface{}
	if !valid.HasErrors() {
		if models.ExistArticleByID(id) {
			data = models.GetArticle(id)
			code = e.SUCCESS
		} else {
			code = e.ERROR_NOT_EXIST_ARTICLE
		}
	} else {
		for _, err := range valid.Errors {
			log.Printf("err.key: %s, err.message: %s", err.Key, err.Message)
		}
	}
}

// 获取多个文章
func GetArticles(c *gin.Context) {

}

//新增文章
func AddArticle(c *gin.Context) {
}

// 修改文章
func EditArticle(c *gin.Context) {

}

// 删除文章
func DeleteArticle(c *gin.Context) {

}
