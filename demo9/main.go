package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

type Page struct {
	ID      uint
	Name    string
	Link    string
	Data    string
	Display bool
}

func main() {
	app := fiber.New()
	type InputData struct {
		Name string `json:"name"`
	}
	app.Post("/page", func(c *fiber.Ctx) error {
		// Phân tích dữ liệu đầu vào từ thân yêu cầu
		input := new(InputData)
		if err := c.BodyParser(input); err != nil {
			return err
		}
		// Xây dựng thông báo đầu ra
		message := "Hello, " + input.Name

		//Connect db
		// dsn := "root:Root@2022@tcp(127.0.0.1:8080)/page?charset=utf8mb4&parseTime=True&loc=Local"
		// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		// db.AutoMigrate(&InputData{})
		// var input1 InputData
		// err1 := db.Modal(&InputData{}).where("name = ?", input.Name).First(&input1)
		// if err1 != nil {
		// 	return c.JSON(fiber.Map{"success": false, "data": err1})
		// }
		// Tạo dữ liệu đầu ra
		// output := fiber.Map{
		// 	"message": message,
		// }

		// Trả về kết quả dưới dạng JSON
		return c.JSON(fiber.Map{"success": true, "data": message})
	})
	log.Fatal(app.Listen(":3000"))

}
