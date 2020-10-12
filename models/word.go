package models

import (
	"github.com/jinzhu/gorm"
)

type Word struct {
	Model

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
}

func ExistWordByID(id int) (bool, error) {
	var word Word
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&word).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if word.ID > 0 {
		return true, nil
	}

	return false, nil
}

func GetWords(maps interface{}) ([]*Word, error) {
	var words []*Word
	err := db.Where(maps).Find(&words).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return words, nil
}

func GetWord(id int) (*Word, error) {
	var word Word
	err := db.Where("id = ? AND deleted_on = ? ", id, 0).First(&word).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return &word, nil
}

func EditWord(id int, data interface{}) error {
	if err := db.Model(&Word{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

func AddWord(data map[string]interface{}) (int, error) {
	word := Word{
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		CreatedBy: data["created_by"].(string),
	}
	if err := db.Create(&word).Error; err != nil {
		return 0, err
	}

	return word.ID, nil
}

func DeleteWord(id int) error {
	if err := db.Where("id = ?", id).Delete(Word{}).Error; err != nil {
		return err
	}

	return nil
}

func CleanAllWord() error {
	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Word{}).Error; err != nil {
		return err
	}

	return nil
}
