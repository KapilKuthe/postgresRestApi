// database/database.go
package database

import (
	"fmt"
	"net/url"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

type User struct {
	ID   uint   `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func InitializeDB() {
	username := "postgres"
	password := "Pass#1230"
	dbName := "test"
	host := "localhost"
	port := 5432

	escapedPassword := url.QueryEscape(password)
	dns := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable", username, escapedPassword, host, port, dbName)
	db, err = gorm.Open(postgres.Open(dns), &gorm.Config{})
	if err != nil {
		fmt.Println(err.Error())
		panic("Failure! DB connection not established...")
	}
	fmt.Println("Connected...")
	db.AutoMigrate(&User{}) // Assuming you have a User model defined in the same package
}

func CloseDB() {
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	sqlDB.Close()
}

func CreateUser(user User) (User, error) {
	result := db.Create(&user)
	if result.Error != nil {
		fmt.Println(result.Error)
		return User{}, result.Error
	}
	return user, nil
}

// func UpdateUser(user User) error {
// 	result := db.Save(&user)
// 	if result.Error != nil {
// 		fmt.Println(result.Error)
// 		return result.Error
// 	}
// 	return nil
// }

func UpdateUser(user User) (User, error) {
	var existingUser User
	result := db.First(&existingUser, user.ID)
	if result.Error != nil {
		fmt.Println(result.Error)
		return User{}, result.Error
	}

	// Update only the non-zero fields from the provided user
	if user.Name != "" {
		existingUser.Name = user.Name
	}
	if user.Age != 0 {
		existingUser.Age = user.Age
	}

	// // result = db.Save(&existingUser)
	// if result.Error != nil {
	// 	fmt.Println(result.Error)
	// 	return result.Error
	// }
	// return nil

	result = db.Model(&existingUser).Updates(existingUser)
	if result.Error != nil {
		fmt.Println(result.Error)
		return User{}, result.Error
	}
	return existingUser, nil
}

func DeleteUser(userID uint) error {
	result := db.Delete(&User{}, userID)
	if result.Error != nil {
		fmt.Println(result.Error)
		return result.Error
	}
	return nil
}

func GetUsers() ([]User, error) {
	var users []User
	result := db.Find(&users)
	if result.Error != nil {
		fmt.Println(result.Error)
		return nil, result.Error
	}
	return users, nil
}

func GetUser(userID uint) (User, error) {
	var user User
	result := db.First(&user, userID)
	if result.Error != nil {
		fmt.Println(result.Error)
		return User{}, result.Error
	}
	return user, nil
}
