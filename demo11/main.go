package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Header struct {
	gorm.Model
	Label  string `json:"label"`
	Format string `json:"format"`
	Key    string `json:"key"`
	Url    string `json:"url"`
}

type SessionHeader struct {
	gorm.Model
	Key string `json:"key"`
	PageId uint `json:"page_id"`
	Metadata json `json:"metadata"`
}

func main() {
	dsn := "yuno:yuno@tcp(127.0.0.1:3306)/backend?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	db.AutoMigrate(&Header{})

	app := fiber.New()

	app.Post("/header", func(c *fiber.Ctx) error {
		input := new(Header)
		if err := c.BodyParser(input); err != nil {
			return err
		}
		db.Create(input)
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Đã thêm thành công",
			"header":  input,
			 // Trả về thông tin của người dùng đã thêm thành công
		})
		// return c.Status(fiber.StatusCreated).JSON(fiber.Map{

		// 	"Label":    input.Label,
		// 	"fulldata": fiber.Map{"key": input.Key, "Data": fiber.Map{"url": input.Url, "Format": input.Format}},
		// })
	})
	app.Get("/header", func(c *fiber.Ctx) error {
		var header []Header
		db.Find(&header)
		return c.JSON(header)
	})

	// Cập nhật thông tin  dựa trên ID
	app.Patch("/header/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var header Header
		// Tìm dựa trên ID
		if err := db.First(&header, id).Error; err != nil {
			// Trả về lỗi nếu không tìm thấy
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Không tìm thấy ",
			})
		}
		// Parse dữ liệu từ yêu cầu vào struct input
		input := new(Header)
		if err := c.BodyParser(input); err != nil {
			return err
		}
		// Cập nhật thông tin với dữ liệu từ input
		if err := db.Model(&header).Updates(input).Error; err != nil {
			// Trả về lỗi nếu cập nhật thất bại
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Cập nhật thông tin không thành công",
			})
		}
		// Trả về thông tin  sau khi cập nhật thành công
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Cập nhật thông tin thành công",
			"header":  header,
		})
	})

	log.Fatal(app.Listen(":3000"))
}
