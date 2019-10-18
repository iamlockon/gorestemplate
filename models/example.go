package models

import (
	"github.com/jinzhu/gorm"
)

type Example struct {
	Model
	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}
// GetExamples gets a list of examples based on paging and constraints
func GetExamples(pageNum int, pageSize int, maps interface{}) ([]Example, error) {
	var (
		examples []Example
		err error
	)
	
	if pageSize > 0 && pageNum > 0 {
		err = db.Where(maps).Find(&examples).Offset(pageNum).Limit(pageSize).Error
	} else {
		err = db.Where(maps).Find(&examples).Error
	}

	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}

	return examples, nil
}
// GetExampleTotal counts the total number of examples based on the constraint
func GetExampleTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Example{}).Where(maps).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// ExistExampleByID determines whether a Example exists based on the ID
func ExistExampleByID(id int) (bool, error) {
	var example Example
	err := db.Select("id").Where("id = ? AND deleted_on = ? ", id, 0).First(&example).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}
	if example.ID > 0 {
		return true, nil
	}

	return false, nil
}

// ExistExampleByName checks if there is a tag with the same name
func ExistExamplByName(name string) (bool, error) {
	var example Example
	err := db.Select("id").Where("name = ? AND deleted_on = ? ", name, 0).First(&example).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return false, err
	}

	if example.ID > 0 {
		return true, nil
	}

	return false, nil
}

// DeleteExample deletes an example
func DeleteExample(id int) error {
	if err := db.Where("id = ?", id).Delete(&Example{}).Error; err != nil {
		return err
	}

	return nil
}

// EditExample modify a single example
func EditExample(id int, data interface{}) error {
	if err := db.Model(&Example{}).Where("id = ? AND deleted_on = ? ", id, 0).Updates(data).Error; err != nil {
		return err
	}

	return nil
}

// AddExample Add an example
func AddExample(name string, state int, createdBy string) error {
	example := Example{
		Name:      name,
		State:     state,
		CreatedBy: createdBy,
	}
	if err := db.Create(&example).Error; err != nil {
		return err
	}

	return nil
}

// CleanAllExample will hard delete the records, and gorm require using of Unscoped() when doing hard deletes.
func CleanAllExample() (bool, error) {
	if err := db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Example{}).Error; err != nil {
		return false, err
	}

	return true, nil
}