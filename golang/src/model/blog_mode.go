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
	result := Db.Find(&data)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func GetOne(id int) (data []Blog) {
	result := Db.First(&data, id)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

func (b *Blog) Create(){
	result := Db.Create(b)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (b *Blog) Update() {
	result := Db.Save(b)
	if result.Error != nil {
		panic(result.Error)
	}
}

func (b *Blog) Delete() {
	result := Db.Delete(b)
	if result.Error != nil {
		panic(result.Error)
	}
}