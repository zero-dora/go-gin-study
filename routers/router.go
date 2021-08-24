package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "github.com/zero-dora/go-gin-example/docs"
	"github.com/zero-dora/go-gin-example/middleware/jwt"
	"github.com/zero-dora/go-gin-example/pkg/export"
	"github.com/zero-dora/go-gin-example/pkg/qrcode"
	"github.com/zero-dora/go-gin-example/pkg/setting"
	"github.com/zero-dora/go-gin-example/pkg/upload"
	"github.com/zero-dora/go-gin-example/routers/api"
	v1 "github.com/zero-dora/go-gin-example/routers/api/v1"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.ServerSetting.RunMode)
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))
	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.GET("/auth", api.GetAuth)
	r.POST("/upload", api.UploadImage)
	apiV1 := r.Group("/api/v1")
	apiV1.Use(jwt.JWT())
	{
		apiV1.GET("/tags", v1.GetTags)           //获取文章标签列表
		apiV1.POST("/tag", v1.AddTag)            //添加文章标签
		apiV1.PUT("/tag/:id", v1.EditTag)        //修改指定文章标签
		apiV1.DELETE("/tag/:id", v1.DeleteTag)   //删除指定文章标签
		apiV1.POST("/tags/export", v1.ExportTag) //导出标签
		apiV1.POST("/tags/import", v1.ImportTag) //导入标签

		apiV1.GET("/articles", v1.GetArticles)                           //获取文章列表
		apiV1.GET("/article/:id", v1.GetArticle)                         //获取指定文章
		apiV1.POST("/article", v1.AddArticle)                            //新建文章
		apiV1.PUT("/article/:id", v1.EditArticle)                        //更新指定文章
		apiV1.DELETE("/article/:id", v1.DeleteArticle)                   //删除指定文章
		apiV1.POST("/article/poster/generate", v1.GenerateArticlePoster) //生成二维码海报
	}
	return r
}
