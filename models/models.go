package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	First_Name string `json:"firstname" validate:"required,min=2,max=30"`
	Last_Name  string `json:"lasttname" validate:"required,min=2,max=30"`
	Username   string `json:"username" validate:"required,unique,min=4"`
	Email      string `json:"email" validate:"email,required"`
	Password   string `json:"password" validate:"required,min=6"`
	User_Id    int    `json:"user_id"`
}

type Owner struct {
	gorm.Model
	First_Name    string `json:"firstname" validate:"required,min=2,max=30"`
	Last_Name     string `json:"lasttname" validate:"required,min=2,max=30"`
	Username      string `json:"username" validate:"required,unique,min=4"`
	Email         string `json:"email" validate:"email,required"`
	Password      string `json:"password" validate:"required,min=6"`
	Owner_Id      int    `json:"user_id"`
	StoreCategory Category
	Balance       int `json:"balance"`
}

type Admin struct {
	gorm.Model
	First_Name string `json:"firstname" validate:"required,min=2,max=30"`
	Last_Name  string `json:"lasttname" validate:"required,min=2,max=30"`
	Username   string `json:"username" validate:"required,unique,min=4"`
	Email      string `json:"email" validate:"email,required"`
	Password   string `json:"password" validate:"required,min=6"`
}

type Category struct {
	gorm.Model
	Clothing     clothing
	Accessories  accessores
	Elecctronics electronics
	OwnerId      int
	CategoryID   int
}

type clothing struct {
	Upperwear      string
	UpperwearPrice uint32
	Lowerwear      string
	LowerwearPrice uint32
	Shoes          string
	ShoesPrice     uint32
	CategoryId     int
}

type accessores struct {
	Watch      string
	Belt       string
	Tie        string
	CategoryId int
}

type electronics struct {
	Mobile     string
	TV         string
	Headphone  string
	CategoryId int
}
