package main

import (
	//"database/sql"
	//"fmt"
	//"html/template"
	//"log"
	//"net/http"
	"github.com/BeHappych/beta_2.0/controller"
	//"example.com/m/go/pkg/mod/github.com/swaggo/swag/example/celler@v0.0.0-20230830154040-e9d0aa572ea5/controller"
	//"example.com/m/go/pkg/mod/github.com/swaggo/swag/example/celler@v0.0.0-20230830154040-e9d0aa572ea5/controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	r := gin.Default()

	c := controller.NewController()

	v1 := r.Group("/api/v1")
	{
		accounts := v1.Group("/accounts")
		{
			//accounts.GET(":id", c.ShowAccount)
			accounts.GET("", c.ListAccounts)
			//accounts.POST("", c.AddAccount)
			//accounts.DELETE(":id", c.DeleteAccount)
			//accounts.PATCH(":id", c.UpdateAccount)
			//accounts.POST(":id/images", c.UploadAccountImage)
		}

	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
