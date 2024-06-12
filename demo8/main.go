package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email"`
}

type Information struct {
	gorm.Model
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Sex     string `json:"sex"`
	Phone   int    `json:"phone"`
	Address string `json:"address"`
}
type Test struct {
	gorm.Model
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Sex     string `json:"sex"`
	Phone   int    `json:"phone"`
	Address string `json:"address"`
}

func main() {
	// Kết nối cơ sở dữ liệu SQL
	dsn := "yuno:yuno@tcp(127.0.0.1:3306)/database1?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}
	// Tự đồng bộ hóa cấu trúc User,.... với cơ sở dữ liệu
	db.AutoMigrate(&User{}, &Information{}, &Test{})
	// Khởi tạo Fiber app
	app := fiber.New()

	// Tạo một user mới
	app.Post("/users", func(c *fiber.Ctx) error {
		input := new(User)
		if err := c.BodyParser(input); err != nil {
			return err
		}
		db.Create(input)
		// Thêm người dùng thành công
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"message": "Đã thêm user thành công",
			"user":    input, // Trả về thông tin của người dùng đã thêm thành công
		})
	})

	// Lấy danh sách tất cả các users
	app.Get("/users", func(c *fiber.Ctx) error {
		var users []User
		db.Find(&users)
		return c.JSON(users)
	})

	// Lấy thông tin của một user dựa trên ID
	app.Get("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var user User
		if err := db.First(&user, id).Error; err != nil {
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Không tìm thấy người dùng",
			})
		}
		return c.JSON(user)
	})
	// Lấy thông tin của một user dựa trên Name
	app.Get("/users/name/:name", func(c *fiber.Ctx) error {
		name := c.Params("name")
		var user User
		db.Where("name = ?", name).First(&user)
		return c.JSON(user)
	})
	//
	app.Get("/users/search", func(c *fiber.Ctx) error {
		// Tạo một struct để lưu trữ thông tin tìm kiếm từ body của yêu cầu
		type SearchRequest struct {
			Name  string `json:"name"`
			Email string `json:"email"`
		}

		// Đọc thông tin tìm kiếm từ body của yêu cầu
		searchReq := new(SearchRequest)
		if err := c.BodyParser(searchReq); err != nil {
			return err
		}

		// Lấy thông tin của người dùng dựa trên tên và email
		var user User
		db.Where(&User{Name: searchReq.Name, Email: searchReq.Email}).First(&user)

		// Kiểm tra xem người dùng có tồn tại không
		if user.ID == 0 {
			// Trả về lỗi nếu không tìm thấy người dùng
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Không tìm thấy người dùng",
			})
		}

		// Trả về thông tin của người dùng nếu tìm thấy
		return c.JSON(user)
	})

	// Cập nhật thông tin của một user dựa trên ID
	app.Patch("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var user User
		// Tìm user dựa trên ID
		if err := db.First(&user, id).Error; err != nil {
			// Trả về lỗi nếu không tìm thấy user
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Không tìm thấy người dùng",
			})
		}
		// Parse dữ liệu từ yêu cầu vào struct input
		input := new(User)
		if err := c.BodyParser(input); err != nil {
			return err
		}
		// Cập nhật thông tin user với dữ liệu từ input
		if err := db.Model(&user).Updates(input).Error; err != nil {
			// Trả về lỗi nếu cập nhật thất bại
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Cập nhật thông tin không thành công",
			})
		}
		// Trả về thông tin của user sau khi cập nhật thành công
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Cập nhật thông tin thành công",
			"user":    user,
		})
	})

	// Xóa một user dựa trên ID
	app.Delete("/users/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var user User
		db.First(&user, id)
		db.Delete(&user)
		return c.JSON(fiber.Map{"message": "Xóa người dùng thành công"})
	})

	app.Post("/information", func(c *fiber.Ctx) error {
		// Tạo một struct mới để lưu trữ dữ liệu từ yêu cầu
		newInfo := new(Information)

		// Parse dữ liệu từ yêu cầu vào struct newInfo
		if err := c.BodyParser(newInfo); err != nil {
			return err
		}

		// Kiểm tra xem dữ liệu đã tồn tại trong cơ sở dữ liệu chưa
		existingInfo := new(Information)
		result := db.Where(" phone = ?", newInfo.Phone).First(existingInfo)
		if result.RowsAffected > 0 {
			// Trả về lỗi nếu dữ liệu đã tồn tại
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "Dữ liệu đã tồn tại",
			})
		}

		// Thêm dữ liệu vào cơ sở dữ liệu
		db.Create(newInfo)

		// Trả về thông báo thành công
		return c.Status(fiber.StatusCreated).JSON(fiber.Map{
			"information": newInfo,
			"thongbao":    "Đã thêm thành công",
		})
	})
	app.Get("/information", func(c *fiber.Ctx) error {
		type SearchInfor struct {
			Name  string `json:"name"`
			Phone int    `json:"phone"`
		}
		readInfor := new(SearchInfor)
		if err := c.BodyParser(readInfor); err != nil {
			return err
		}
		var thongtin Information
		db.Where(&Information{Name: readInfor.Name, Phone: readInfor.Phone}).First(&thongtin)
		// Kiểm tra xem người dùng có tồn tại không
		if thongtin.ID == 0 {
			// Trả về lỗi nếu không tìm thấy người dùng
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Không tìm thấy thong tin",
			})
		}
		// Trả về thông tin của người dùng nếu tìm thấy
		return c.JSON(thongtin)
	})
	app.Patch("/information", func(c *fiber.Ctx) error {
		// Tạo một struct để lưu trữ yêu cầu cập nhật
		type UpdateRequest struct {
			ID      uint   `json:"id"`
			Name    string `json:"name"`
			Age     int    `json:"age"`
			Sex     string `json:"sex"`
			Phone   int    `json:"phone"`
			Address string `json:"address"`
		}

		// Parse dữ liệu từ yêu cầu vào struct UpdateRequest
		updateRequest := new(UpdateRequest)
		if err := c.BodyParser(updateRequest); err != nil {
			return err
		}

		// Tạo một biến để lưu thông tin cần cập nhật
		var thongtin Information

		// Tìm thông tin dựa trên ID
		if err := db.First(&thongtin, updateRequest.ID).Error; err != nil {
			// Trả về lỗi nếu không tìm thấy thông tin
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Không tìm thấy thông tin",
			})
		}

		// Cập nhật thông tin chỉ khi có dữ liệu mới được cung cấp trong yêu cầu cập nhật
		if updateRequest.Name != "" {
			thongtin.Name = updateRequest.Name
		}
		if updateRequest.Age != 0 {
			thongtin.Age = updateRequest.Age
		}
		if updateRequest.Sex != "" {
			thongtin.Sex = updateRequest.Sex
		}
		if updateRequest.Phone != 0 {
			thongtin.Phone = updateRequest.Phone
		}
		if updateRequest.Address != "" {
			thongtin.Address = updateRequest.Address
		}

		// Lưu thay đổi vào cơ sở dữ liệu
		if err := db.Save(&thongtin).Error; err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Cập nhật thông tin không thành công",
			})
		}

		// Trả về thông tin sau khi cập nhật thành công
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message":     "Cập nhật thông tin thành công",
			"information": thongtin,
		})

	})

	app.Delete("/information", func(c *fiber.Ctx) error {
		// Tạo một struct để lưu trữ yêu cầu xóa
		type DeleteRequest struct {
			ID uint `json:"id"`
		}

		// Parse dữ liệu từ yêu cầu vào struct DeleteRequest
		deleteRequest := new(DeleteRequest)
		if err := c.BodyParser(deleteRequest); err != nil {
			return err
		}

		// Tạo một biến để lưu thông tin cần xóa
		var thongtin Information

		// Tìm thông tin dựa trên ID
		if err := db.First(&thongtin, deleteRequest.ID).Error; err != nil {
			// Trả về lỗi nếu không tìm thấy thông tin
			return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
				"error": "Không tìm thấy thông tin",
			})
		}

		// Xóa thông tin từ cơ sở dữ liệu
		if err := db.Delete(&thongtin).Error; err != nil {
			// Trả về lỗi nếu không thể xóa thông tin
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": "Xóa thông tin không thành công",
			})
		}

		// Trả về thông báo thành công nếu xóa thành công
		return c.Status(fiber.StatusOK).JSON(fiber.Map{
			"message": "Xóa thông tin thành công",
		})
	})

	// Chạy ứng dụng trên cổng 3000
	log.Fatal(app.Listen(":3000"))

}
