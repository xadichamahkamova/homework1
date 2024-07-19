package api

import (
	"service/api/handler"
	"service/pkg"
	"service/producer"
	"service/storage"

	"github.com/gin-gonic/gin"
)

func NewGin(m *pkg.Mongo, pro producer.Producer) *gin.Engine {

	r := gin.Default()

	handler := handler.HandlerSt{
		Task: storage.NewTaskRepo(m),
		Producer: pro,
	}

	r.POST("/task", handler.CreateTask)
	r.GET("/task/:id", handler.GetTaskById)
	r.GET("/task", handler.ListOfTask)
	r.PUT("/task", handler.UpdateTask)
	r.DELETE("/task/:id", handler.DeleteTask)

	return r
}