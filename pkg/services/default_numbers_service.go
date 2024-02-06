package services

import (
	"context"
	"errors"
	"strconv"

	"multipliers/pkg/core"
	"multipliers/pkg/models"
)

type DefaultNumbersService struct {
	storage map[int]string
}

func NewDefaultNumbersService() *DefaultNumbersService {
	return &DefaultNumbersService{
		storage: make(map[int]string),
	}
}

func (d DefaultNumbersService) Save(_ context.Context, number *models.Number) error {

	numberToFind, err := strconv.Atoi(*number.Number)
	if err != nil {
		return errors.New("number not valid")
	}

	if _, ok := d.storage[numberToFind]; !ok {
		d.storage[numberToFind] = core.GetMultiplierType(numberToFind)
	}

	return nil
}

func (d DefaultNumbersService) Find(_ context.Context, number *models.Number) (*models.Number, error) {

	numberToFind, err := strconv.Atoi(*number.Number)
	if err != nil {
		return nil, errors.New("number not valid")
	}

	if numberFound, ok := d.storage[numberToFind]; ok {
		return &models.Number{Number: &numberFound}, nil
	}

	return nil, errors.New("number not found")
}

func (d DefaultNumbersService) FindAll(_ context.Context) ([]string, error) {

	var numbers []string
	for _, value := range d.storage {
		numbers = append(numbers, value)
	}

	return numbers, nil
}
