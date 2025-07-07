package helper

import "gorm.io/gorm"

type Employ struct {
	gorm.Model
	Name   string `json:"name"`
	Salary int    `json:"salary"`
	Post   string `json:"post"`
}
