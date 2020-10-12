package v1

import (
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/jalilbengoufa/go-search/controllers/word_controller"
	"github.com/jalilbengoufa/go-search/pkg/app"
	"github.com/jalilbengoufa/go-search/pkg/e"
	"github.com/jalilbengoufa/go-search/pkg/redis"
	"github.com/unknwon/com"
)

type AddWordForm struct {
	Title     string `form:"title" valid:"Required;MaxSize(100)"`
	Desc      string `form:"desc" valid:"Required;MaxSize(255)"`
	CreatedBy string `form:"created_by" valid:"Required;MaxSize(100)"`
}

func GetWord(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	valid := validation.Validation{}
	valid.Min(id, 1, "id")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		c.JSON(http.StatusBadRequest, gin.H{"status": e.INVALID_PARAMS})
		return
	}

	wordController := word_controller.Word{ID: id}
	exists, err := wordController.ExistByID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": e.ERROR_CHECK_EXIST_ARTICLE_FAIL})
		return
	}
	if !exists {
		c.JSON(http.StatusOK, gin.H{"status": e.ERROR_NOT_EXIST_ARTICLE})
		return
	}

	word, err := wordController.Get()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": e.ERROR_GET_ARTICLE_FAIL})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": e.SUCCESS, "data": word})
}

func AddWord(c *gin.Context) {
	var (
		form AddWordForm
	)

	httpCode, errCode := app.BindAndValid(c, &form)
	if errCode != e.SUCCESS {
		c.JSON(httpCode, gin.H{"status code": errCode})
		return
	}

	wordController := word_controller.Word{
		Title:     form.Title,
		Desc:      form.Desc,
		CreatedBy: form.CreatedBy,
	}
	if err := wordController.Add(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": e.ERROR_ADD_ARTICLE_FAIL})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": e.SUCCESS})
}

func GetWords(c *gin.Context) {
	valid := validation.Validation{}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		c.JSON(http.StatusBadRequest, gin.H{"status": e.INVALID_PARAMS})
		return
	}

	wordController := word_controller.Word{}

	words, err := wordController.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"status": e.ERROR_GET_ARTICLES_FAIL})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": e.SUCCESS, "data": words})
}

func FindWord(c *gin.Context) {

	word := com.StrTo(c.Query("word")).String()
	valid := validation.Validation{}
	valid.Required(word, "word").Message("there must be a word to search for")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		c.JSON(http.StatusBadRequest, gin.H{"status": e.INVALID_PARAMS, "error": valid.Errors})
		return
	}

	docs, total, err := redis.Find(word)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": e.SUCCESS, "data": docs, "total": total})

}

func Autocomplete(c *gin.Context) {

	word := com.StrTo(c.Query("word")).String()
	valid := validation.Validation{}
	valid.Required(word, "word").Message("there must be a word to search for")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		c.JSON(http.StatusBadRequest, gin.H{"status": e.INVALID_PARAMS, "error": valid.Errors})
		return
	}

	suggestions, err := redis.Autocomplete(word)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": e.SUCCESS, "data": suggestions})

}
