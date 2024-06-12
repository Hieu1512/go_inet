package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	type SomeStruct struct {
		Name    string
		Age     uint8
		Address string
	}
	type Kame struct {
		Ten    string
		Tuoi   uint8
		Diachi string
	}

	app := fiber.New()
	type InputData struct {
		Name string `json:"name"`
	}
	// vd nếu 1 là nam 2 là nữ
	type Gioitinh int

	const (
		Nam Gioitinh = 1
		Nu  Gioitinh = 2
	)

	type Gioitinhdata struct {
		Gioitinh Gioitinh `json:"Sex"`
	}

	app.Post("/hello", func(c *fiber.Ctx) error {
		input := new(Gioitinhdata)
		if err := c.BodyParser(input); err != nil {
			return err
		}
		var gioitinh string
		switch input.Gioitinh {
		case Nam:
			gioitinh = "Gioi tinh la: Nam"
		case Nu:
			gioitinh = "Gioi tinh la: Nu"
		default:
			gioitinh = "k hop le"
		}

		output := fiber.Map{
			"sex": gioitinh,
		}
		return c.JSON(output)
	})

	type Gioitinhd struct {
		Gt int `json:"Sex"`
	}

	app.Put("/demo", func(c *fiber.Ctx) error {
		input := new(Gioitinhd)
		if err := c.BodyParser(input); err != nil {
			return err
		}
		var Sex string
		if input.Gt == 1 {
			Sex = "Nam"
		} else {
			Sex = "Nu"
		}
		ouput := fiber.Map{
			"Sex": Sex,
		}
		return c.JSON(ouput)
	})

	// vd Tính tổng các số trong một dãy số
	// type tongso struct {
	// 	Values int `json:"Number"`
	// }
	// app.Get("/demo3", func(c *fiber.Ctx) error {
	// 	input := new(tongso)
	// 	if err := c.BodyParser(input); err != nil {
	// 		return err
	// 	}
	// 		total := sum(input.Values)

	// 		output := map[string]int{"total": total}
	// 		return c.JSON(output)
	// 	})

	// 1 ví dụ lấy ra tuổi
	type tuoi struct {
		Tuoib int `json:"Age"`
	}

	app.Get("/demo3", func(c *fiber.Ctx) error {
		input := new(tuoi)
		if err := c.BodyParser(input); err != nil {
			return err
		}
		add := input.Tuoib

		ouput := fiber.Map{
			"Age": add,
		}
		return c.JSON(ouput)
	})

	//
	type animal struct {
		Dongv string `json:"Animal"`
		Tuoi  int    `json:"Age"`
	}
	app.Put("/demo2", func(c *fiber.Ctx) error {
		input := new(animal)
		if err := c.BodyParser(input); err != nil {
			return err
		}
		dongvat := input.Dongv
		age := input.Tuoi
		output := fiber.Map{
			"Animal": dongvat,
			"Age":    age,
		}
		return c.JSON(output)
	})

	type tgtn struct {
		Gio float64 `json:"Hour"`
	}
	app.Patch("/thoigian", func(c *fiber.Ctx) error {
		input := new(tgtn)
		if err := c.BodyParser(input); err != nil {
			return err
		}
		gio := input.Gio
		var period string
		if gio >= 0.00 && gio < 12.00 {
			period = "sáng"
		} else if gio >= 12.00 && gio < 18.00 {
			period = "chiều"
		} else {
			period = "tối"
		}
		output := fiber.Map{
			"Thoi gian": period,
		}
		return c.JSON(output)
	})

	//  1 ví dụ về thời gian
	type time struct {
		Tgt string `json:"Times"`
	}
	app.Get("/demo4", func(c *fiber.Ctx) error {
		input := new(time)
		if err := c.BodyParser(input); err != nil {
			return err
		}
		timet := "Tgian sang toi :" + input.Tgt
		ouput := fiber.Map{
			"Time": timet,
		}
		return c.JSON(ouput)
	})
	app.Post("/home", func(c *fiber.Ctx) error {
		// Phân tích dữ liệu đầu vào từ thân yêu cầu
		input := new(InputData)
		if err := c.BodyParser(input); err != nil {
			return err
		}
		// Xây dựng thông báo đầu ra
		message := "Hello, " + input.Name

		// Tạo dữ liệu đầu ra
		output := fiber.Map{
			"message": message,
		}

		// Trả về kết quả dưới dạng JSON
		return c.JSON(output)
	})
	app.Get("/home", func(c *fiber.Ctx) error {
		// Phân tích dữ liệu đầu vào từ thân yêu cầu
		input := new(InputData)
		if err := c.BodyParser(input); err != nil {
			return err
		}

		// Xây dựng thông báo đầu ra
		message := "Hello, " + input.Name

		// Tạo dữ liệu đầu ra
		output := fiber.Map{
			"message": message,
		}

		// Trả về kết quả dưới dạng JSON
		return c.JSON(output)
	})

	app.Get("/", func(c *fiber.Ctx) error {
		data := Kame{
			Ten: "Hieu",
		}
		return c.JSON(data)
	})
	app.Post("/", func(c *fiber.Ctx) error {
		data := SomeStruct{
			Name:    "Grame",
			Age:     20,
			Address: "Ha Noi",
		}
		return c.JSON(data)
	})
	app.Post("/hello", func(c *fiber.Ctx) error {
		p := new(InputData)

		if err := c.BodyParser(p); err != nil {
			return err
		}
		return c.JSON(p)
	})
	app.Delete("/", func(c *fiber.Ctx) error {
		return c.SendString("KJHHJ")

	})

	// vd về nhận nhập và thông báo tk,mk vào data
	type Person struct {
		Name string `json:"name"`
		Pass string `json:"pass"`
	}
	app.Post("/person", func(c *fiber.Ctx) error {
		p := new(Person)
		if err := c.BodyParser(p); err != nil {
			return err
		}

		// Ghi nhật ký dữ liệu đã nhận hiển thị lên terminal
		log.Println("Name:", p.Name)
		log.Println("Password:", p.Pass)

		ouput := map[string]interface{}{
			"tk": p.Name,
			"mk": p.Pass,
		}

		// Trả về phản hồi thành công
		return c.JSON(ouput)
	})

	log.Fatal(app.Listen(":3000"))

}
