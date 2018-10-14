package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/sushi86/timemanagement/controllers"
	"github.com/sushi86/timemanagement/models"
)

func main() {
	e := echo.New()

	db, err := gorm.Open("mysql", "root:secret@/timetracker?charset=utf8&parseTime=true")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	db.AutoMigrate(models.Project{})
	db.AutoMigrate(models.Team{})
	db.AutoMigrate(models.User{})

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	h := &controllers.Handler{DB: db}

	auth := e.Group("/auth")
	auth.POST("/token", h.Login)

	users := e.Group("/users")
	users.Use(middleware.JWT([]byte("test")))
	users.GET("/:id", h.GetUser)
	users.POST("/", h.CreateUser)
	users.PUT("/:id", h.UpdateUser)
	users.DELETE("/:id", h.DeleteUser)

	projects := e.Group("/projects")
	projects.GET("/", h.GetProjects)
	projects.GET("/:id", h.GetProject)
	projects.POST("/", h.CreateProject)
	projects.PUT("/:id", h.UpdateProject)
	projects.DELETE("/:id", h.DeleteProject)

	teams := e.Group("/teams")
	teams.GET("/", h.GetTeams)
	teams.GET("/:id", h.GetTeam)
	teams.POST("/", h.CreateTeam)
	teams.PUT("/:id", h.UpdateTeam)
	teams.DELETE("/:id", h.DeleteTeam)

	//times:= e.Group("/times")
	//times.GET("/:project", h.GetTime)
	//times.POST("/:project", h.AddTime)
	//times.PUT("/:id", h.UpdateTime)
	//times.DELETE("///:id", h.DeleteTime)

	e.Logger.Fatal(e.Start(":1323"))
}