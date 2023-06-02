package main

// @title gin_api_frame API
// @version 1.0
// @description The api docs of gin_api_frame project
// @termsOfService http://swagger.io/terms/

// @contact.name Zhao Haihang
// @contact.url http://www.swagger.io/support
// @contact.email 1932859223@qq.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:4000
// @BasePath /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	app, err := CreateApp()
	if err != nil {
		panic(err)
	}
	if err = app.Start() ;err != nil {
		panic(err)
	}
	app.AwaitSignal()
}
