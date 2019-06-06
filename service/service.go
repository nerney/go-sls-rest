package service

import (
	"errors"
	"go-sls-rest/data"
	"go-sls-rest/models"

	"github.com/teris-io/shortid"
)

// GetAll items or return an empty array.
func GetAll() []models.Item {
	return data.GetAll()
}

// GetOne item by id, returns an error if not found.
func GetOne(id string) (models.Item, error) {
	if item := data.GetOne(id); item != nil {
		return *item, nil
	}
	return models.Item{}, errors.New("not found")
}

// Create a new item and return it with the newly generated id.
func Create(txt models.ItemTxt) models.Item {
	return *data.Put(models.Item{ID: shortid.MustGenerate(), Txt: txt})
}

// Update an existing item, returns an error if not found.
func Update(item models.Item) error {
	if _, err := GetOne(item.ID); err != nil {
		return err
	}
	data.Put(item)
	return nil
}

// Delete an item, returns an error if not found.
func Delete(id string) error {
	if _, err := GetOne(id); err != nil {
		return err
	}
	data.Delete(id)
	return nil
}
