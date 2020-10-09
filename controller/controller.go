package controller

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/my-Sakura/SMService/model/mysql"
	"github.com/my-Sakura/SMService/utils"
)

type SMSController struct {
	DB  *sql.DB
	Con *utils.Config
}

type Message struct {
	Mobile string `json:"mobile"`
	Code   string `json:"code"`
}

func NewSMSController(db *sql.DB, con *utils.Config) *SMSController {
	return &SMSController{
		DB:  db,
		Con: con,
	}
}

func (s *SMSController) RegistRouter(r gin.IRouter) {
	err := mysql.CreateDatabase(s.DB)
	if err != nil {
		log.Println(err)
	}

	mysql.CreateTable(s.DB)

	r.POST("/send", s.Send)
}

func (s *SMSController) Send(c *gin.Context) {
	var msg Message
	msg.Code = utils.Rand(s.Con.Length)
	c.ShouldBind(&msg)

	mysql.Insert(s.DB, msg.Mobile, msg.Code)

	url := fmt.Sprintf(s.Con.Url, msg.Mobile, msg.Code, s.Con.TemplateCode)
	if err := utils.Send(s.Con.Method, url, s.Con.AppCode); err != nil {
		c.Error(err)
		c.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest})
		return
	}

	c.JSON(http.StatusOK, gin.H{"statusCode": http.StatusOK})
}
