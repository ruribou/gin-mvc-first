package model

import (
	"gorm.io/gorm"
)

type Blog struct {
	gorm.Model
	Title   string
	Content string
}

func GetAll() (data []Blog) {
	result := db.Find(&data)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func GetOne(id int) (data Blog) {
	result := db.First(&data, id)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func (b *Blog) Create(){
	result := db.Create(b)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (b *Blog) Update() {
	result := db.Save(b)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (b *Blog) Delete() {
	result := db.Delete(b)
	if result.Error != nil {
		panic(result.Error)
	}
}