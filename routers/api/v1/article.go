package v1

import (
	"github.com/astaxie/beego/validation"
	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"github.com/zero-dora/go-gin-example/pkg/app"
	"github.com/zero-dora/go-gin-example/pkg/e"
	"github.com/zero-dora/go-gin-example/pkg/logging"
	"github.com/zero-dora/go-gin-example/pkg/qrcode"
	"github.com/zero-dora/go-gin-example/pkg/setting"
	"github.com/zero-dora/go-gin-example/pkg/util"
	"github.com/zero-dora/go-gin-example/service/article_service"
	"github.com/zero-dora/go-gin-example/service/tag_service"
	"net/http"
)

//@Summary 获取单个文章
//@Produce json
//@Param id query int true "ID"
//@Success 200 {string} json "{"code":200,"msg":"ok","data":{}}"
//@Router  /api/v1/article/{id} [get]
func GetArticle(c *gin.Context) {
	appG := app.Gin{c}
	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
	}

	articleService := article_service.Article{ID: id}
	exist, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}

	if !exist {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	article, err := articleService.Get()
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR_GET_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, article)
}

//@Summary 获取多个文章
//@Produce json
//@Param state query int false "State"
//@Param tag_id query int false "标签id"
//@Success 200 {string} json "{"code":200,"msg":"ok","data":{}}"
//@Router  /api/v1/articles [get]
func GetArticles(c *gin.Context) {
	appG := app.Gin{c}

	valid := validation.Validation{}

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}
	var tagId int = -1
	if arg := c.Query("tag_id"); arg != "" {
		tagId = com.StrTo(arg).MustInt()
		valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	}

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	articleService := article_service.Article{
		TagID:    tagId,
		State:    state,
		PageNum:  util.GetPage(c),
		PageSize: setting.AppSetting.PageSize,
	}
	total, err := articleService.Count()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_COUNT_ARTICLE_FAIL, nil)
		return
	}

	articles, err := articleService.GetAll()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}

	data := make(map[string]interface{})
	data["list"] = articles
	data["total"] = total
	appG.Response(http.StatusOK, e.SUCCESS, data)
}

//
//type AddArticleForm struct {
//	TagID         int    `form:"tag_id" valid:"Required;Min(1)"`
//	Title         string `form:"title" valid:"Required;MaxSize(100)"`
//	Desc          string `form:"desc" valid:"Required;MaxSize(255)"`
//	Content       string `form:"content" valid:"Required;MaxSize(65535)"`
//	CreatedBy     string `form:"created_by" valid:"Required;MaxSize(100)"`
//	CoverImageUrl string `form:"cover_image_url" valid:"Required;MaxSize(255)"`
//	State         int    `form:"state" valid:"Range(0,1)"`
//}

//@Summary 新增文章
//@Produce json
//@Param tag_id query int true "标签id"
//@Param title query string true "标题"
//@Param desc query string true "简述"
//@Param content query string true "内容"
//@Param state query int false "State:状态只允许0或1"
//@Param created_by query string true "创建人"
//@Success 200 {string} json "{"code":200,"msg":"ok","data":{}}"
//@Router  /api/v1/article [post]
func AddArticle(c *gin.Context) {
	appG := app.Gin{c}

	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	coverImageUrl := c.Query("cover_image_url")

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.Required(content, "content").Message("内容不能为空")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	valid.Required(coverImageUrl, "cover_image_url").Message("图片不能为空")
	valid.MaxSize(coverImageUrl, 255, "cover_image_url").Message("图片地址长度不能大于255")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	articleService := article_service.Article{
		TagID:         tagId,
		Title:         title,
		Desc:          desc,
		Content:       content,
		CoverImageUrl: coverImageUrl,
		State:         state,
		CreatedBy:     createdBy,
	}

	err := articleService.Add()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_GET_ARTICLES_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

//@Summary 编辑文章
//@Produce json
//@Param id path int true "文章id"
//@Param tag_id query int true "标签id"
//@Param title query string true "标题"
//@Param desc query string true "简述"
//@Param content query string true "内容"
//@Param state query int false "State:状态只允许0或1"
//@Param modified_by query string true "修改人"
//@Success 200 {string} json "{"code":200,"msg":"ok","data":{}}"
//@Router  /api/v1/article/{id} [put]
func EditArticle(c *gin.Context) {
	appG := app.Gin{c}

	valid := validation.Validation{}

	id := com.StrTo(c.Param("id")).MustInt()
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")
	coverImageUrl := c.Query("cover_image_url")

	var state int = -1
	if arg := c.Query("state"); arg != "" {
		state = com.StrTo(arg).MustInt()
		valid.Range(state, 0, 1, "state").Message("状态只允许0或1")
	}

	valid.Min(id, 1, "id").Message("ID必须大于0")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人最长为100字符")
	valid.MaxSize(coverImageUrl, 255, "cover_image_url").Message("图片地址长度不能大于255")

	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	articleService := article_service.Article{
		ID:            id,
		TagID:         tagId,
		Title:         title,
		Desc:          desc,
		Content:       content,
		CoverImageUrl: coverImageUrl,
		ModifiedBy:    modifiedBy,
	}
	exists, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	tagService := tag_service.Tag{ID: tagId}
	exists, err = tagService.ExistByID()

	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EXIST_TAG_FAIL, nil)
		return
	}

	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_TAG, nil)
		return
	}

	err = articleService.Edit()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_EDIT_ARTICLE_FAIL, nil)
		return
	}
	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

//@Summary 删除文章
//@Produce json
//@Param id path int true "文章id"
//@Success 200 {string} json "{"code":200,"msg":"ok","data":{}}"
//@Router  /api/v1/article/{id} [delete]
func DeleteArticle(c *gin.Context) {
	appG := app.Gin{C: c}

	id := com.StrTo(c.Param("id")).MustInt()

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("ID必须大于0")
	if valid.HasErrors() {
		app.MarkErrors(valid.Errors)
		appG.Response(http.StatusOK, e.INVALID_PARAMS, nil)
		return
	}
	articleService := article_service.Article{ID: id}
	exists, err := articleService.ExistByID()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_CHECK_EXIST_ARTICLE_FAIL, nil)
		return
	}
	if !exists {
		appG.Response(http.StatusOK, e.ERROR_NOT_EXIST_ARTICLE, nil)
		return
	}

	err = articleService.Delete()
	if err != nil {
		appG.Response(http.StatusInternalServerError, e.ERROR_DELETE_ARTICLE_FAIL, nil)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)

}

func GenerateArticlePoster(c *gin.Context) {
	appG := app.Gin{C: c}

	article := &article_service.Article{}
	qr := qrcode.NewQrCode(c.Query("qrcode_url"), 300, 300, qr.M, qr.Auto)
	posterName := article_service.GetPosterFlag() + "-" + qrcode.GetQrCodeFileName(qr.URL) + qr.GetQrCodeExt()
	articlePoster := article_service.NewArticlePoster(posterName, article, qr)
	articlePosterBgService := article_service.NewArticlePosterBg(
		"bg.jpg",
		articlePoster,
		&article_service.Rect{
			X0: 0,
			Y0: 0,
			X1: 550,
			Y1: 700,
		},
		&article_service.Pt{
			X: 125,
			Y: 298,
		},
	)
	_,filePath,err := articlePosterBgService.Generate()
	if err != nil{
		logging.Info(err)
		appG.Response(http.StatusOK,e.ERROR_GEN_ARTICLE_POSTER_FAIL,nil)
		return
	}
	appG.Response(http.StatusOK,e.SUCCESS,map[string]string{
		"poster_url" : qrcode.GetQrCodeFullUrl(posterName),
		"poster_save_url" : filePath + posterName,
	})

}
