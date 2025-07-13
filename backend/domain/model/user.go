package model

import (
	"errors"
	"strings"
	"time"
)

type User struct {
	ID        int
	Name      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// NewUser creates a new user with validation
func NewUser(name string) (*User, error) {
	user := &User{
		Name:      name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	
	if err := user.Validate(); err != nil {
		return nil, err
	}
	
	return user, nil
}

// Validate validates user data according to business rules
func (u *User) Validate() error {
	if err := u.validateName(); err != nil {
		return err
	}
	return nil
}

// validateName validates the user name according to business rules
func (u *User) validateName() error {
	if u.Name == "" {
		return errors.New("user name cannot be empty")
	}
	
	// Trim whitespace
	u.Name = strings.TrimSpace(u.Name)
	
	if len(u.Name) == 0 {
		return errors.New("user name cannot be only whitespace")
	}
	
	if len(u.Name) < 2 {
		return errors.New("user name must be at least 2 characters long")
	}
	
	if len(u.Name) > 50 {
		return errors.New("user name must be at most 50 characters long")
	}
	
	return nil
}

// UpdateName updates the user name with validation
func (u *User) UpdateName(newName string) error {
	oldName := u.Name
	u.Name = newName
	
	if err := u.validateName(); err != nil {
		u.Name = oldName // revert on error
		return err
	}
	
	u.UpdatedAt = time.Now()
	return nil
}

// IsNewUser checks if the user is newly created (within last 24 hours)
func (u *User) IsNewUser() bool {
	return time.Since(u.CreatedAt) < 24*time.Hour
}