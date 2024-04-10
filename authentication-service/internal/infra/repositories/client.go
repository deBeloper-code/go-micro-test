package repositories

import (
	"fmt"

	"gorm.io/gorm"
)

type client struct {
	db *gorm.DB
}

// NewClient returns a new instance to use postgres.
func NewClient() *client {
	return &client{
		db: connect(),
	}
}

// GetAll returns a slice of all users, sorted by last name
func (c *client) GetAll(dest interface{}, conds ...interface{}) error {
	// Find results
	if err := c.db.Find(dest).Error; err != nil {
		return err
	}
	return nil
}

func (c *client) GetByEmail(dest interface{}, email string) error {
	// Fist results
	if err := c.db.First(dest, "email = ?", email).Error; err != nil {
		return err
	}
	return nil
}

func (c *client) GetById(dest interface{}, id int) error {
	// Fist results
	if err := c.db.First(dest, "id = ?", id).Error; err != nil {
		return err
	}
	return nil
}
func (c *client) UpdateInfo(dest interface{}, newUpdate interface{}, id int) error {

	result := c.db.Model(dest).Where("id = ?", id).Updates(newUpdate)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("ID didn't find %d", id)
	}

	return nil
}

func (c *client) Delete(dest interface{}, id int) error {

	result := c.db.Where("id = ?", id).Delete(dest)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("ID didn't find %d", id)
	}

	return nil
}

func (c *client) Create(dest interface{}) error {
	// Fist results
	if err := c.db.Create(dest).Error; err != nil {
		return err
	}
	return nil
}

func (c *client) ResetPassword(dest interface{}, id int, newPassword string) error {

	result := c.db.Model(dest).Where("id = ?", id).Update("password", newPassword)
	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("ID didn't find %d", id)
	}

	return nil
}
