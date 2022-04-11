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

func (*handler) GetDateTimeslotByDate(c *gin.Context) {
	var input time.Time
	param := c.Param("date")
	var year int
	if s, err := strconv.Atoi(strings.Split(param, "-")[0]); err == nil {
		year = s
	}
	var month int
	if s, err := strconv.Atoi(strings.Split(param, "-")[1]); err == nil {
		month = s
	}
	var date int
	if s, err := strconv.Atoi(strings.Split(param, "-")[2]); err == nil {
		date = s
	}
	log.Println("Params: ", param)
	log.Println(year, time.Month(month), date)
	input = time.Date(year, time.Month(month), date, 0, 0, 0, 0, time.Local)
	var dateTimeSlots = _DTSlogic.GetDateTimeslotByDate(input)
	c.JSON(http.StatusOK, dateTimeSlots)
}
