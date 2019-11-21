package routes

import (
	"os"
	"project-service/configs"
	ctr "project-service/controllers"
	mdw "project-service/middleware"

	"github.com/gin-gonic/gin"
)

type (
	// ProjectService struct for
	ProjectService struct{}
)

func init() {
	configs.InitASCII() // declare ASCII for beauty command line
	configs.Connect()
}

// Run function for running called in main.go
func (p *ProjectService) Run() {
	// connect ke DB
	p.routing()
}

func (p *ProjectService) routing() {
	r := gin.Default()
	r.Use(mdw.CheckMiddleware)
	v1 := r.Group("api/v1")
	{
		v1.GET("/token", ctr.GetToken)

		cg := v1.Group("/categories")
		{
			catCtrl := ctr.CategoryCtrl{}
			cg.GET("/", catCtrl.GetCategories)
			cg.GET("/:id", catCtrl.GetCategory)
		}
		project := v1.Group("/project")
		{
			projectCtr := ctr.ProjectCtr{}
			project.POST("/save", projectCtr.Create)
			project.GET("/", projectCtr.GetProjects)
			project.GET("/:id", projectCtr.GetProject)
			project.POST("/edit/:id", projectCtr.Update)

			gallery := project.Group("/gallery")
			{

				gallery.POST("/create", projectCtr.Create)
			}
		}

	}
	// run this router
	r.Run(os.Getenv("APPS_PORT"))
}
