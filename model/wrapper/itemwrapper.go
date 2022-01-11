package wrapper

import (
	"shoppinglist/model"

	"gorm.io/gorm"
)

// ItemWrapper using the items
type ItemWrapper struct {
	DB *gorm.DB
}

//CreateItemWrapper using the items
func CreateItemWrapper(db *gorm.DB) *ItemWrapper {
	return &ItemWrapper{
		DB: db,
	}
}

//Insert the items
func (s *ItemWrapper) Insert(data interface{}) error {
	item := data.(*model.Item)

	return s.DB.Create(&item).Error
}

//Update the items
func (s *ItemWrapper) Update(data interface{}) error {
	item := data.(*model.Item)

	return s.DB.Save(&item).Error
}

//Delete the items
func (s *ItemWrapper) Delete(data interface{}) error {
	item := data.(*model.Item)

	return s.DB.Delete(&item).Error
}
