package api

import (
	"clips/internal/config"
	"clips/internal/db"
	"clips/pkg/models"
	"fmt"
	"github.com/go-pg/pg/v10"
	"github.com/gofiber/fiber/v2"
	"log"
	"strconv"
	"strings"
	"time"
)

// NewClip is called whenever a user hits POST /api/clip/new. It reads the options given in the form in Upload.svelte,
// creates a Clip in the database, stores the uploaded video file in the configured clip storage path, and responds with
// the JSON representation of the newly created Clip.
func NewClip(c *fiber.Ctx) error {
	if c.Locals("active") != true {
		return c.Status(401).JSON(fiber.Map{
			"error":"Guests are not allowed to upload",
		})
	}

	if !config.AllowUpload {
		return c.Status(500).JSON(fiber.Map{
			"error":"Uploading is disabled",
		})
	}

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
		Type: clipFile.Header.Get("Content-Type"),
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

// GetClip is called whenever a user hits GET /api/clip/:clip_id. It responds with information about the specified Clip
// in JSON format.
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

// GetClips is called whenever a user hits GET /api/clip/get. It returns a JSON array of Clips, retrieved by pagination
// from the db, specified by the query parameters (page,amount).
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

// ViewClip is called whenever a user hits GET /api/clip/view/:clip_id. It responds with the specified Clip's raw video
// data stored in the configured clip storage path.
func ViewClip(c *fiber.Ctx) error {
	clipID,_ := strconv.ParseInt(c.Params("clip_id","0"),10,64)

	var clipSource string

	err := db.Database.Model(&models.Clip{}).
		Where("clip_id = ?",clipID).
		Column("source").
		Select(&clipSource)

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

	return c.SendFile(clipSource,true)
}

// IncrementViews is called whenever a user hits POST /api/clip/views/:clip_id. It increments the amount of views of
// the specified Clip by the :clip_id parameter, and responds with an empty body & 200 status code.
func IncrementViews(c *fiber.Ctx) error {
	clipID,_ := strconv.ParseInt(c.Params("clip_id","0"),10,64)

	_,err := db.Database.Model(&models.Clip{}).
		Set("views = views + 1").
		Where("clip_id = ?",clipID).
		Update()

	if err != nil {
		log.Println(err)

		return c.Status(500).JSON(fiber.Map{
			"error":"Unknown Error",
		})
	}

	return c.SendStatus(200)
}