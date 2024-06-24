package servicies

import (
	"gin-fleamarket/dto"
	"gin-fleamarket/models"
	"gin-fleamarket/repositries"
)

type IItemservice interface {
	FindAll() (*[]models.Item, error)
	FindById(itemId uint) (*models.Item, error)
	Create(CreateItemInput dto.CreateItemInput) (*models.Item, error)
	Update(itemId uint, updateItemInput dto.UpdateItemInput) (*models.Item, error)
	Delete(itemid uint) error
}

type ItemService struct {
	repository repositries.IItemRepository
}

func NewItemService(repository repositries.IItemRepository) IItemservice {
	return &ItemService{repository: repository}
}

func (s *ItemService) FindAll() (*[]models.Item, error) {
	return s.repository.FindAll()
}

func (s *ItemService) FindById(itemId uint) (*models.Item, error) {
	return s.repository.FindById(itemId)
}

func (s *ItemService) Create(createItemInput dto.CreateItemInput) (*models.Item, error) {
	newItem := models.Item{
		Name:        createItemInput.Name,
		Price:       createItemInput.Price,
		Description: createItemInput.Description,
		SoldOut:     false,
	}
	return s.repository.Create(newItem)
}

func (s *ItemService) Update(itemId uint, updateItemInput dto.UpdateItemInput) (*models.Item, error) {
	targetItem, err := s.FindById(itemId)
	if err != nil {
		return nil, err
	}

	if updateItemInput.Name != nil {
		targetItem.Name = *updateItemInput.Name
	}
	if updateItemInput.Price != nil {
		targetItem.Price = *updateItemInput.Price
	}
	if updateItemInput.Description != nil {
		targetItem.Description = *updateItemInput.Description
	}
	if updateItemInput.SoldOut != nil {
		targetItem.SoldOut = *updateItemInput.SoldOut
	}

	return s.repository.Update(*targetItem)
}

func (s *ItemService) Delete(itemId uint) error {
	return s.repository.Delete(itemId)
}
