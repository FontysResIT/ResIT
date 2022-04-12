package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/RealSnowKid/ResIT/logic"
	"github.com/gin-gonic/gin"
)

var _DTSlogic logic.IDateTimeslotLogic

type DTShandler struct{}

func NewDateTimeslotHandler(logic logic.IDateTimeslotLogic) *handler {
	_DTSlogic = logic
	return &handler{}
}

// @Description Get all date time slots
// @Accept  json
// @Produce  json
// @Success 200 {array} []model.DateTimeSlot "ok"
// @Router /dateTimeSlots [get]
func (*handler) GetAllDateTimeslots(c *gin.Context) {
	var dateTimeSlot = _DTSlogic.GetAllDateTimeslots()
	c.JSON(http.StatusOK, dateTimeSlot)
}

func (*handler) GetDateTimeslotByParam(c *gin.Context) {
	query := c.Param("query")
	param := c.Param("param")
	log.Println("Query:", query)
	switch query {
	case "date":
		var params = strings.Split(strings.Split(param, "/")[1], "-")
		var dates = []int{}

		for _, v := range params {
			j, err := strconv.Atoi(v)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			dates = append(dates, j)
		}
		date := time.Date(dates[0], time.Month(dates[1]), dates[2], 0, 0, 0, 0, time.Local)
		log.Println("Date param:", date)
		dateTimeSlots := _DTSlogic.GetDateTimeslotsByDate(date)
		c.JSON(http.StatusOK, dateTimeSlots)
	case "dateId":
		var params = strings.Split(strings.Split(param, "/")[1], "-")
		var dates = []int{}

		for _, v := range params {
			j, err := strconv.Atoi(v)
			if err != nil {
				c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				return
			}
			dates = append(dates, j)
		}
		date := time.Date(dates[0], time.Month(dates[1]), dates[2], 0, 0, 0, 0, time.Local)
		log.Println("Date param:", date)
		dateTimeSlot := _DTSlogic.GetDateTimeslotByDate(date)
		c.JSON(http.StatusOK, dateTimeSlot)
	}
}
