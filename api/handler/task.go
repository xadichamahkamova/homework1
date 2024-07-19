package handler

import (
	"encoding/json"
	"service/models"
	"service/producer"
	"service/storage"

	"github.com/gin-gonic/gin"
	amqp "github.com/rabbitmq/amqp091-go"
)

type HandlerSt struct {
	Task *storage.TaskRepo
	Producer producer.Producer
}

func(h *HandlerSt) CreateTask(c *gin.Context) {

	req := models.Task{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message":err.Error()})
		return
	}
	
	byteData, err := json.Marshal(&req)
	if err != nil {
		c.JSON(500, gin.H{"message":err.Error()})
		return
	}
	err = h.Producer.Channel.Publish(
		"",                    
		h.Producer.Queue.Name, 
		false,                
		false,                 
		amqp.Publishing{
			ContentType: "application/json",
			Body:        byteData,
		},
	)
	if err != nil {
		c.JSON(500, gin.H{"message":err.Error()})
		return
	}

	resp, err := h.Task.CreateTask(&req)
	if err != nil {
		c.JSON(500, gin.H{"message":err.Error()})
		return
	}
	c.JSON(200, resp)
}

func(h *HandlerSt) GetTaskById(c *gin.Context) {

	id := c.Param("id")
	resp, err := h.Task.GetTaskById(id)
	if err != nil {
		c.JSON(500, gin.H{"message":err.Error()})
		return
	}

	byteData, err := json.Marshal(&resp)
	if err != nil {
		c.JSON(500, gin.H{"message":err.Error()})
		return
	}
	err = h.Producer.Channel.Publish(
		"",                    
		h.Producer.Queue.Name, 
		false,                 
		false,                 
		amqp.Publishing{
			ContentType: "application/json",
			Body:        byteData,
		},
	)
	if err != nil {
		c.JSON(500, gin.H{"message":err.Error()})
		return
	}

	c.JSON(200, resp)
}

func(h *HandlerSt) ListOfTask(c *gin.Context) {

	resp, err := h.Task.ListOfTask()
	if err != nil {
		c.JSON(500, gin.H{"message":err.Error()})
		return
	}

	byteData, err := json.Marshal(&resp)
	if err != nil {
		c.JSON(500, gin.H{"message":err.Error()})
		return
	}
	err = h.Producer.Channel.Publish(
		"",                   
		h.Producer.Queue.Name, 
		false,                 
		false,                
		amqp.Publishing{
			ContentType: "application/json",
			Body:        byteData,
		},
	)
	if err != nil {
		c.JSON(500, gin.H{"message":err.Error()})
		return
	}

	c.JSON(200, resp)
}

func(h *HandlerSt) UpdateTask(c *gin.Context) {

	req := models.Task{}
	if err := c.BindJSON(&req); err != nil {
		c.JSON(400, gin.H{"message":err.Error()})
		return
	}

	resp, err := h.Task.UpdateTask(&req)
	if err != nil {
		c.JSON(500, gin.H{"message":err.Error()})
		return
	}

	byteData, err := json.Marshal(&resp)
	if err != nil {
		c.JSON(500, gin.H{"message":err.Error()})
		return
	}
	err = h.Producer.Channel.Publish(
		"",                    
		h.Producer.Queue.Name, 
		false,                 
		false,                 
		amqp.Publishing{
			ContentType: "application/json",
			Body:        byteData,
		},
	)
	if err != nil {
		c.JSON(500, gin.H{"message":err.Error()})
		return
	}

	c.JSON(200, resp)
}

func(h *HandlerSt) DeleteTask(c *gin.Context) {

	id := c.Param("id")
	resp, err := h.Task.DeleteTask(id)
	if err != nil {
		c.JSON(500, gin.H{"message":err.Error()})
		return
	}

	byteData, err := json.Marshal(&resp)
	if err != nil {
		c.JSON(500, gin.H{"message":err.Error()})
		return
	}
	err = h.Producer.Channel.Publish(
		"",                    
		h.Producer.Queue.Name, 
		false,                 
		false,                 
		amqp.Publishing{
			ContentType: "application/json",
			Body:        byteData,
		},
	)
	if err != nil {
		c.JSON(500, gin.H{"message":err.Error()})
		return
	}

	c.JSON(200, resp)
}