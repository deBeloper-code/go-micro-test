package entity

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID        int       `json:"id"`
	Email     string    `json:"email" gorm:"unique"`
	FirstName string    `json:"first_name,omitempty"`
	LastName  string    `json:"last_name,omitempty"`
	Password  string    `json:"password" gorm:"not null"`
	Active    int       `json:"active"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = int(uuid.New().ID())
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()

	return
}

func (u *User) BeforeUpdate(tx *gorm.DB) (err error) {
	// Update time
	u.UpdatedAt = time.Now()
	return
}
