package handler

import (
	"net/http"

	"github.com/FontysResIT/ResIT/logic"
	"github.com/FontysResIT/ResIT/model"
	"github.com/gin-gonic/gin"
)

var _TSlogic logic.ITimeSlotLogic

type TShandler struct{}

func NewTimeSlotHandler(logic logic.ITimeSlotLogic) *handler {
	_TSlogic = logic
	return &handler{}
}

// @Description Get all timeslots
// @Accept  json
// @Produce  json
// @Success 200 {array} []model.TimeSlot "ok"
// @Router /timeslots [get]
func (*handler) GetAllTimeSlots(c *gin.Context) {
	var timeSlot = _TSlogic.GetAllTimeSlots()
	c.JSON(http.StatusOK, timeSlot)
}

// @Description Create a timeslot
// @Accept  json
// @Produce  json
// @Param	data	body	model.TimeSlot	true	"The timeslot to create"
// @Success 200 {object} model.TimeSlot "ok"
// @Router /timeslots [post]
func (*handler) CreateTimeSlot(c *gin.Context) {
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
