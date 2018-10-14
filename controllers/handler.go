package controllers

import "github.com/jinzhu/gorm"

type (
	Handler struct {
		DB  *gorm.DB
	}

	Response struct {
		Status string `json:"status"`
		Message string `json:"message,omitempty"`
		Error string `json:"error,omitempty"`
	}
)
