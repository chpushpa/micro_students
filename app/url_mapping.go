package app

import (
	"github.com/chpushpa/micro_student/controllers"
)

func mapUrls() {
	router.POST("/student", controllers.Create)
	//router.GET("/student/:id", controllers.Get)

}
