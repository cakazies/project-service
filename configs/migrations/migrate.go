package migrations

import (
	"log"

	"github.com/cakazies/project-service/configs"
	"github.com/cakazies/project-service/models"
)

func main() {
	log.Println("Loading ...")
	configs.InitASCII()
	configs.Connect()
	configs.GetDB.AutoMigrate(&models.Project{}, &models.ProjectDetail{})
	configs.GetDB.AutoMigrate(&models.ProjectTImeline{}, &models.ProjectTimelineGallery{})
	configs.GetDB.AutoMigrate(&models.ProjectDocument{}, &models.CategoryTypes{})
	configs.GetDB.AutoMigrate(&models.ProjectGallery{}, &models.Category{})
	// configs.GetDB.Model(&bnimodels.BniLogResponse{}).ModifyColumn("data", "text")

	defer configs.GetDB.Close()
	log.Println("Done ...")
}
