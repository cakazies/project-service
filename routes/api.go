package routes

import (
	"os"
	"project-service/configs"
	ctr "project-service/controllers"

	"github.com/gin-gonic/gin"
)

type ProjectService struct{}

func (p *ProjectService) Run() {
	// connect ke DB
	configs.InitAscii() // declare ASCII for beauty command line
	configs.Connect()
	p.routing()
}

func (p *ProjectService) routing() {
	r := gin.Default()
	// r.Use(mdw.CheckMiddleware())
	v1 := r.Group("api/v1")
	{
		cg := v1.Group("/categories")
		{
			catCtrl := ctr.CategoryCtrl{}
			cg.GET("/", catCtrl.GetCategories)
			cg.GET("/:id", catCtrl.GetCategory)
		}

	}
	// run this router
	r.Run(os.Getenv("APPS_PORT"))
}
