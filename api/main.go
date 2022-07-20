package api

import (
	"api_gateway/config"
	"api_gateway/pkg/logger"
	"api_gateway/services"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	// @Summary 登录
	// @Description 登录
	// @Produce json
	// @Param body body controllers.LoginParams true "body参数"
	// @Success 200 {string} string "ok" "返回用户信息"
	// @Failure 400 {string} string "err_code：10002 参数错误； err_code：10003 校验错误"
	// @Failure 401 {string} string "err_code：10001 登录失败"
	// @Failure 500 {string} string "err_code：20001 服务错误；err_code：20002 接口错误；err_code：20003 无数据错误；err_code：20004 数据库异常；err_code：20005 缓存异常"
	// @Router /user/person/login [post]
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "api_gateway/api/docs"
	v1 "api_gateway/api/handlers/v1"
)

type RouterOptions struct {
	Log      logger.Logger
	Cfg      config.Config
	Services services.ServiceManager
}

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func New(opt *RouterOptions) *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowCredentials = true
	config.AllowHeaders = append(config.AllowHeaders, "*")

	router.Use(cors.New(config))
	// router.Use(MaxAllowed(100))

	handlerV1 := v1.New(&v1.HandlerV1Options{
		Log:      opt.Log,
		Cfg:      opt.Cfg,
		Services: opt.Services,
	})

	router.GET("/config", handlerV1.GetConfig)

	apiV1 := router.Group("/v1")
	apiV1.GET("/ping", handlerV1.Ping)

	// profession
	apiV1.POST("/profession", handlerV1.CreateProfession)
	apiV1.GET("/profession", handlerV1.GetAllProfession)
	apiV1.GET("/profession/:profession_id", handlerV1.GetProfession)
	apiV1.PUT("/profession/:profession_id", handlerV1.UpdateProfession)
	apiV1.DELETE("/profession/:profession_id", handlerV1.DeleteProfession)

	// attribute
	apiV1.POST("/attribute", handlerV1.CreateAttribute)
	apiV1.GET("/attribute/:attribute_id", handlerV1.GetAttribute)
	apiV1.GET("/attribute", handlerV1.GetAllAttributes)
	apiV1.PUT("/attribute/:attribute_id", handlerV1.UpdateAttribute)
	apiV1.DELETE("/attribute/:attribute_id", handlerV1.DeleteAttribute)

	// company
	apiV1.POST("/company", handlerV1.CreateCompany)
	apiV1.GET("/company", handlerV1.GetAllCompanies)
	apiV1.GET("/company/:company_id", handlerV1.GetCompany)
	apiV1.PUT("/company/:company_id", handlerV1.UpdateCompany)
	apiV1.DELETE("/company/:company_id", handlerV1.DeleteCompany)

	// position
	apiV1.POST("/position", handlerV1.CreatePosition)
	apiV1.GET("/position", handlerV1.GetAllPositions)
	apiV1.GET("/position/:position_id", handlerV1.GetPosition)
	apiV1.PUT("/position/:position_id", handlerV1.UpdatePosition)
	apiV1.DELETE("/position/:position_id", handlerV1.DeletePosition)

	// swagger
	url := ginSwagger.URL("swagger/doc.json") // The url pointing to API definition
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return router
}

func MaxAllowed(n int) gin.HandlerFunc {
	sem := make(chan struct{}, n)
	acquire := func() { sem <- struct{}{} }
	release := func() { <-sem }
	return func(c *gin.Context) {
		acquire()       // before request
		defer release() // after request
		c.Next()

	}
}
