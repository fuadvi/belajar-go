package belajar_go_gorm

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"strconv"
	"testing"
)

func OpenConnection() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/belajar_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		panic(err)
	}
	return db
}

var db = OpenConnection()

func TestOpenConnection(t *testing.T) {
	assert.NotNil(t, db)
}

func TestCreateUsers(t *testing.T) {
	user := User{
		ID:       "1",
		Password: "rahasia",
		Name:     "fuad",
	}

	response := db.Create(&user)
	assert.Nil(t, response.Error)
	assert.Equal(t, int64(1), response.RowsAffected)
}

func TestCreateBulkUsers(t *testing.T) {
	var user []User

	for i := 2; i < 10; i++ {
		user = append(user, User{
			ID:       strconv.Itoa(i),
			Password: "rahasia",
			Name:     "fuad " + strconv.Itoa(i),
		})
	}

	response := db.Create(&user)
	assert.Nil(t, response.Error)
	assert.Equal(t, int64(8), response.RowsAffected)
}

func TestTransaction(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {

		tx.Create(&User{
			ID:       "10",
			Password: "rahasia",
			Name:     "fuad 10",
		})

		tx.Create(&User{
			ID:       "11",
			Password: "rahasia",
			Name:     "fuad 11",
		})

		tx.Create(&User{
			ID:       "13",
			Password: "rahasia",
			Name:     "fuad 13",
		})

		return nil
	})

	assert.Nil(t, err)
}

func TestErrorTransaction(t *testing.T) {
	err := db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(&User{
			ID:       "14",
			Password: "rahasia",
			Name:     "fuad 14",
		}).Error

		if err != nil {
			return err
		}

		err = tx.Create(&User{
			ID:       "13",
			Password: "rahasia",
			Name:     "fuad 13",
		}).Error

		if err != nil {
			return err
		}

		return nil
	})

	assert.NotNil(t, err)
}

func TestCreateWallet(t *testing.T) {
	wallet := Wallet{
		ID:      "1",
		UserId:  "1",
		Balance: 10000,
	}
	err := db.Create(&wallet).Error
	assert.Nil(t, err)
}

func TestPreload(t *testing.T) {
	var user User
	err := db.Model(&User{}).Preload("Wallet").First(&user).Error

	assert.Nil(t, err)

	fmt.Println(user.Wallet)
}

func TestJoin(t *testing.T) {
	var user User
	err := db.Model(&User{}).Joins("Wallet").First(&user).Error

	assert.Nil(t, err)

	fmt.Println(user.Wallet)
}
