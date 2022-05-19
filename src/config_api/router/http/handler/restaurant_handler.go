package handler

import (
	"net/http"

	"github.com/FontysResIT/config_api/logic"
	"github.com/FontysResIT/config_api/model"
	"github.com/gin-gonic/gin"
)

var _restaurantLogic logic.IRestaurantLogic

type RestaurantHandler struct{}

func NewRestaurantHandler(logic logic.IRestaurantLogic) *RestaurantHandler {
	_restaurantLogic = logic
	return &RestaurantHandler{}
}

func (*RestaurantHandler) GetAll(c *gin.Context) {
	var timeSlot = _restaurantLogic.GetAllRestaurants()
	c.JSON(http.StatusOK, timeSlot)
}

func (*RestaurantHandler) Create(c *gin.Context) {
	var input model.Restaurant
	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	reservation, err := _restaurantLogic.CreateRestaurant(input)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": "Oops, something went wrong"})
		return
	}
	c.JSON(http.StatusOK, reservation)
}
