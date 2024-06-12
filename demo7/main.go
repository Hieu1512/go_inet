package main

import (
	"log"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	ID       uint
	Name     string
	Age      int
	Birthday time.Time
}

type Datab struct {
	ID   uint
	Name string `gorm:"primaryKey"`
	Age  string
}

// gorm.Model definition
type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
type Alo struct {
	Hoten string `gorm:"primaryKey"`
	ID    uint
	Tuoi  int
}

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "yuno:yuno@tcp(127.0.0.1:3306)/database1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	// Tự đồng bộ hóa cấu trúc User,.... với cơ sở dữ liệu
	db.AutoMigrate(&User{}, &Datab{}, &Model{}, &Alo{})

	// migrator := db.Migrator()
	// migrator.AlterColumn(&Alo{}, "ID", "varchar(100)")
	// db.Datab("ID","")
	// user := []*User{
	// 	{Name: "Jinzhu", Age: 18, Birthday: time.Now()},
	// 	{Name: " Hieu", Age: 19, Birthday: time.Now()},
	// 	{Name: " Hang", Age: 17, Birthday: time.Now()},
	// }
	datat := []*Datab{
		{Name: "Bon", Age: "9"},
		{Name: "Nam", Age: "8"},
	}
	db.Create(&datat) // pass pointer of data to Create
	// for _, datab := range data {
	// 	db.Create(datab)
	// }

	// log.Println("Error:", result.Error)                // returns error
	// log.Println("Rows affected:", result.RowsAffected) // returns inserted records count
}
