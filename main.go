package selfserveruploader

import (
	"fmt"
	"strings"

	"github.com/Taiwrash/randomstring"
	"github.com/gofiber/fiber/v2"
)

func Upload(c *fiber.Ctx) error {
	file, err := c.FormFile("image")
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Image not selected",
		})
	}

	uniqid := randomstring.RandomString(16)

	filename := strings.Replace(uniqid, "-", "", -1)

	ext := strings.Split(file.Filename, ".")[1]
	image := fmt.Sprintf("%s.%s", filename, ext)
	err = c.SaveFile(file, fmt.Sprintf("./uploads/%s", image))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "unable to save!",
		})
	}

	imgUrl := fmt.Sprintf("http://localhost:4000/images/%s", image)

	data := map[string]interface{}{
		"name":     image,
		"imageURL": imgUrl,
		"header":   file.Header,
		"size":     file.Size,
	}

	return c.Status(fiber.StatusCreated).JSON(data)
}
