package logic

import (
	"testing"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/FontysResIT/config_api/model"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type RestaurantRepository struct {
	mock.Mock
}
type MockProducer struct {
	mock.Mock
}

var restaurant = model.Restaurant{Id: primitive.NewObjectID(), Name: "DIM"}

func (mock *RestaurantRepository) All() []model.Restaurant {
	args := mock.Called()
	result := args.Get(0)
	return result.([]model.Restaurant)
}
func (mock *RestaurantRepository) Create(Restaurant model.Restaurant) (model.Restaurant, error) {
	mock.Called()
	return Restaurant, nil
}

func TestGetAll(t *testing.T) {
	mockRepo := new(RestaurantRepository)

	restaurant := model.Restaurant{Id: primitive.NewObjectID(), Name: "DIM"}

	mockRepo.On("All").Return([]model.Restaurant{restaurant})

	testService := NewRestaurantLogic(mockRepo)

	testService.GetAllRestaurants()

	mockRepo.AssertExpectations(t)
}

func TestCreate(t *testing.T) {
	mockRepo := new(RestaurantRepository)

	restaurant := model.Restaurant{Id: primitive.NewObjectID(), Name: "DIM"}

	// Setup expectations
	mockRepo.On("Create").Return(restaurant, nil)
	testService := NewRestaurantLogic(mockRepo)

	result, err := testService.CreateRestaurant(restaurant)
	mockRepo.AssertExpectations(t)

	//Data Assertion
	assert.Equal(t, restaurant.Id, result.Id)
	assert.Equal(t, err, nil)
}
