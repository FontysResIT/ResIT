package handler

import (
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/RealSnowKid/ResIT/logic"
	"github.com/RealSnowKid/ResIT/model"
	"github.com/gin-gonic/gin"
)

var _logic logic.IReservationLogic

type handler struct{}

func NewReservationHandler(logic logic.IReservationLogic) *handler {
	_logic = logic
	return &handler{}
}

// @Description Get all reservations
// @Accept  json
// @Produce  json
// @Success 200 {array} []model.Reservation "ok"
// @Router /reservation [get]
func (*handler) GetAllReservations(c *gin.Context) {
	var reservation = _logic.GetAllReservations()
	c.JSON(http.StatusOK, reservation)
}

func (*handler) GetAllReservationsByDate(c *gin.Context) {
	dateParams := strings.Split(c.Param("date"), "-")
	var intDateParams = []int{}
	for _, v := range dateParams {
		j, err := strconv.Atoi(v)
		if err != nil {
			panic(err)
		}
		intDateParams = append(intDateParams, j)
	}
	date := time.Date(intDateParams[0], time.Month(intDateParams[1]), intDateParams[2], 0, 0, 0, 0, time.Local)
	log.Println("DATE", date)
	var reservations = _logic.GetAllReservationsByDate(date)
	c.JSON(http.StatusOK, reservations)
}

// @Description Create reservation
// @Accept  json
// @Produce  json
// @Success 200 mongo.InsertOneResult "ok"
// @Router /reservation [post]
func (*handler) CreateReservation(c *gin.Context) {
	log.Println("first log")
	var input model.Reservation
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		log.Println(input)
		return
	}
	log.Println(input)
	var reservation = _logic.CreateReservation(input)
	c.JSON(http.StatusOK, reservation)

	// var reservationoutcome = _logic.CreateReservation(model.Reservation)
	// c.JSON(http.StatusOK, reservation)
}
