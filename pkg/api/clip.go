package api

import (
	"clips/pkg/config"
	"clips/pkg/db"
	"clips/pkg/models"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"strings"
	"time"
)

func NewClip(c *fiber.Ctx) error {
	user := c.Locals("user").(models.User)

	clipFile,err := c.FormFile("clip")

	if err != nil {
		log.Println(err)

		return c.Status(500).JSON(fiber.Map{
			"error":"Unexpected Error",
		})
	}

	splitFilename := strings.Split(clipFile.Filename,".")

	if len(splitFilename) <= 1 {
		log.Println(err)

		return c.Status(400).JSON(fiber.Map{
			"error":"Invalid Filename",
		})
	}

	ext := splitFilename[len(splitFilename) - 1]
	filename := fmt.Sprintf("%s-%d.%s",user.Username,time.Now().UnixNano(),ext)
	clipFilePath := fmt.Sprintf("%s/%s",config.ClipSavePath,filename)

	if err = c.SaveFile(clipFile,clipFilePath); err != nil {
		log.Println(err)
	}

	clipTitle := c.FormValue("title","No Title")
	clipDescription := c.FormValue("description","No Description")

	clip := models.Clip{
		UserID: user.UserID,
		Creator: user.Username,
		Title: clipTitle,
		Description: clipDescription,
		Source: clipFilePath,
	}

	_,err = db.Database.Model(&clip).
		Insert()

	if err != nil {
		log.Println(err)

		return c.Status(500).JSON(fiber.Map{
			"error":"Unexpected Error",
		})
	}

	return c.JSON(&clip)
}

func GetClip(c *fiber.Ctx) error {
	clipID,_ := strconv.ParseInt(c.Params("clip_id","0"),10,64)

	clip := models.Clip{}

	err := db.Database.Model(&clip).
		Where("clip_id = ?",clipID).
		Select()

	if err != nil {

		if err == pg.ErrNoRows {
			return c.Status(404).JSON(fiber.Map{
				"error":"Clip not found",
			})
		}

		log.Println(err)

		return c.Status(500).JSON(fiber.Map{
			"error":"Unknown Error",
		})
	}

	return c.JSON(&clip)
}

func GetClips(c *fiber.Ctx) error {
	page,_ := strconv.Atoi(c.Query("page","0"))
	amount,_ := strconv.Atoi(c.Query("amount","10"))

	var clips []models.Clip

	err := db.Database.Model(&clips).
		Order("clip_id ASC").
		Offset(page * amount).
		Limit(amount).
		Select()

	if err != nil {

		if err == pg.ErrNoRows {
			return c.Status(404).JSON(fiber.Map{
				"error":"No clips found",
			})
		}

		log.Println(err)

		return c.Status(500).JSON(fiber.Map{
			"error":"Unknown Error",
		})
	}

	return c.JSON(&clips)
}

func ViewClip(c *fiber.Ctx) error {
	clipID,_ := strconv.ParseInt(c.Params("clip_id","0"),10,64)

	clip := models.Clip{}

	err := db.Database.Model(&clip).
		Where("clip_id = ?",clipID).
		Select()

	if err != nil {

		if err == pg.ErrNoRows {
			if err == pg.ErrNoRows {
				return c.Status(404).JSON(fiber.Map{
					"error":"Clip not found",
				})
			}

			log.Println(err)

			return c.Status(500).JSON(fiber.Map{
				"error":"Unknown Error",
			})
		}

	}

	return c.SendFile(clip.Source,true)
}

func IncrementViews(c *fiber.Ctx) error {
	clipID,_ := strconv.ParseInt(c.Params("clip_id","0"),10,64)

	var views uint64

	err := db.Database.Model(&models.Clip{}).
		Column("views").
		Where("clip_id = ?",clipID).
		Select(&views)

	if err != nil {

		if err == pg.ErrNoRows {
			if err == pg.ErrNoRows {
				return c.Status(404).JSON(fiber.Map{
					"error":"Clip not found",
				})
			}

			log.Println(err)

			return c.Status(500).JSON(fiber.Map{
				"error":"Unknown Error",
			})
		}

	}

	_,err = db.Database.Model(&models.Clip{}).
		Set("views = views + 1").
		Where("clip_id = ?",clipID).
		Update()

	if err != nil {
		log.Println(err)

		return c.Status(500).JSON(fiber.Map{
			"error":"Unknown Error",
		})
	}

	return c.JSON(views + 1)
}