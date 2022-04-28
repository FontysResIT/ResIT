package handler

import (
	"net/http"

	"github.com/FontysResIT/datetime_api/logic"
	"github.com/FontysResIT/datetime_api/model"
	"github.com/gin-gonic/gin"
)

var _TSlogic logic.ITimeSlotLogic

type TShandler struct{}

func NewTimeSlotHandler(logic logic.ITimeSlotLogic) *TShandler {
	_TSlogic = logic
	return &TShandler{}
}

// @Description Get all timeslots
// @Accept  json
// @Produce  json
// @Success 200 {array} []model.TimeSlot "ok"
// @Router /timeslots [get]
func (*TShandler) GetAllTimeSlots(c *gin.Context) {
	var timeSlot = _TSlogic.GetAllTimeSlots()
	c.JSON(http.StatusOK, timeSlot)
}

// @Description Create a timeslot
// @Accept  json
// @Produce  json
// @Param	data	body	model.TimeSlot	true	"The timeslot to create"
// @Success 200 {object} model.TimeSlot "ok"
// @Router /timeslots [post]
func (*TShandler) CreateTimeSlot(c *gin.Context) {
	var input model.TimeSlot
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	timeslot, err := _TSlogic.CreateTimeSlot(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Error creating timeslot"})
		return
	}
	c.JSON(http.StatusOK, timeslot)
}
