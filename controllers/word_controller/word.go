package word_controller

import (
	"github.com/jalilbengoufa/go-search/models"
	"github.com/jalilbengoufa/go-search/pkg/redis"
)

type Word struct {
	ID         int
	Title      string
	Desc       string
	CreatedBy  string
	ModifiedBy string
}

func (a *Word) Add() error {
	word := map[string]interface{}{
		"title":      a.Title,
		"desc":       a.Desc,
		"created_by": a.CreatedBy,
	}
	id, err := models.AddWord(word)
	if err != nil {
		return err
	}

	err = redis.Insert(a.Title, a.Desc, id)

	if err != nil {
		return err
	}

	return nil
}

func (a *Word) Edit() error {
	return models.EditWord(a.ID, map[string]interface{}{
		"title":       a.Title,
		"desc":        a.Desc,
		"modified_by": a.ModifiedBy,
	})
}

func (a *Word) Get() (*models.Word, error) {

	word, err := models.GetWord(a.ID)
	if err != nil {
		return nil, err
	}

	return word, nil
}

func (a *Word) GetAll() ([]*models.Word, error) {

	words, err := models.GetWords(a.getMaps())
	if err != nil {
		return nil, err
	}
	return words, nil
}

func (a *Word) Delete() error {
	return models.DeleteWord(a.ID)
}

func (a *Word) ExistByID() (bool, error) {
	return models.ExistWordByID(a.ID)
}

func (a *Word) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0

	return maps
}
