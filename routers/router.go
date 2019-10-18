package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	_ "github.com/iamlockon/gorestemplate/docs" //must import this package for swagger to load docs.
	"github.com/iamlockon/gorestemplate/pkg/setting"
	"github.com/iamlockon/gorestemplate/routers/api"
	"github.com/iamlockon/gorestemplate/routers/api/v1"
)

// @title <Title>
// @version <API version>
// @description <Description>
// @termsOfService <API terms of service>

// @contact.name <contact infomation>
// @contact.url <contact url>
// @contact.email <contact email>

// @license.name <license name>
// @license.url <license url>

// @host <API host>
// @BasePath <base path : /api/v1>

// @securityDefinitions.basic BasicAuth

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationurl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information

// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationurl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information

// @x-extension-openapi {"example": "value on a json format"}
func InitRouter() *gin.Engine {
	// Init Gin engine
	r := gin.New()

	// Add middlewares as you wish
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Set run mode
	gin.SetMode(setting.ServerSetting.RunMode)

	// Add auth route
	r.GET("/auth", api.GetAuth)

	// Swagger API docs
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// Example routes
	apiv1 := r.Group("api/v1")

	{
		apiv1.GET("/examples", v1.GetExamples)
		apiv1.POST("/examples", v1.AddExample)
		apiv1.PUT("/examples/:id", v1.EditExample)
		apiv1.DELETE("/examples/:id", v1.DeleteExample)
	}

	return r
}
