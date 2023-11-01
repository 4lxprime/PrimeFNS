package handlers

import (
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"time"

	utils "github.com/4lxprime/Goldmp/internal/utils"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/text/encoding/charmap"
)

type cloudFile struct {
	UniqueFilename string            `json:"uniqueFilename"`
	Filename       string            `json:"filename"`
	Hash           string            `json:"hash"`
	Hash256        string            `json:"hash256"`
	Length         int               `json:"length"`
	ContentType    string            `json:"contentType"`
	Uploaded       time.Time         `json:"uploaded"`
	StorageType    string            `json:"storageType"`
	StorageIds     map[string]string `json:"storageIds"`
	AccountId      string            `json:"accountId"`
	DoNotCache     bool              `json:"doNotCache"`
}

// ["/fortnite/api/cloudstorage/system"]
func HandleCloudstorageSystem(c *fiber.Ctx) error {
	dir := "internal/api/cloudstorage/"

	var cloudFiles []cloudFile

	files, _ := os.ReadDir(dir)
	for _, file := range files {
		if filepath.Ext(file.Name()) == ".ini" {
			filePath := filepath.Join(dir, file.Name())
			content, _ := os.ReadFile(filePath)
			fileInfo, _ := os.Stat(filePath)

			cloudFiles = append(cloudFiles, cloudFile{
				UniqueFilename: file.Name(),
				Filename:       file.Name(),
				Hash:           utils.HashString(sha1.New(), string(content)),
				Hash256:        utils.HashString(sha256.New(), string(content)),
				Length:         len(content),
				ContentType:    "application/octet-stream",
				Uploaded:       fileInfo.ModTime(),
				StorageType:    "S3",
				StorageIds:     map[string]string{},
				DoNotCache:     true,
			})
		}
	}

	return c.JSON(cloudFiles)
}

// ["/fortnite/api/cloudstorage/user/:accountId"]
func HandleCloudstorageUser(c *fiber.Ctx) error {
	version, err := utils.GetVersion(&c.Request().Header)
	if err != nil {
		return err
	}

	file := fmt.Sprintf("internal/api/clientsettings/ClientSettings-%s.Sav", version.CL)

	fileInfo, err := os.Stat(file)
	if os.IsNotExist(err) {
		return c.JSON([]fiber.Map{})
	}

	fileContent, err := os.ReadFile(file)
	if err != nil {
		return err
	}

	decoder := charmap.ISO8859_1.NewDecoder()
	latin1FileContent, err := decoder.Bytes(fileContent)
	if err != nil {
		return err
	}

	cfile := cloudFile{
		UniqueFilename: "ClientSettings.Sav",
		Filename:       "ClientSettings.Sav",
		Hash:           utils.HashString(sha1.New(), string(latin1FileContent)),
		Hash256:        utils.HashString(sha256.New(), string(latin1FileContent)),
		Length:         len(latin1FileContent),
		ContentType:    "application/octet-stream",
		Uploaded:       fileInfo.ModTime(),
		StorageType:    "S3",
		StorageIds:     map[string]string{},
		AccountId:      "", // todo: handle this
		DoNotCache:     true,
	}

	return c.JSON(cfile)
}

// ["/fortnite/api/cloudstorage/user/:accountId/:file"]
func HandleCloudstorageUserFile(c *fiber.Ctx) error {
	if c.Params("file") != "clientsettings.sav" {
		return c.JSON(fiber.Map{"error": "file not found"})
	}

	version, err := utils.GetVersion(&c.Request().Header)
	if err != nil {
		return err
	}

	file := fmt.Sprintf("internal/api/clientsettings/ClientSettings-%s.Sav", version.CL)

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := charmap.ISO8859_1.NewEncoder()
	w := encoder.Writer(f)

	if _, err := w.Write(c.Request().Body()); err != nil {
		return err
	}

	return c.SendStatus(204)
}

// ["/fortnite/api/cloudstorage/system/:file"]
func HandleCloudstorageSystemFile(c *fiber.Ctx) error {
	fileName := c.Params("file")
	if fileName == "" {
		return c.SendStatus(200) // blank file
	}

	filePath := fmt.Sprintf("internal/api/cloudstorage/%s", fileName)
	if _, err := os.Stat(filePath); os.IsExist(err) {
		file, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}

		c.Type("application/octet-stream")
		return c.Send(file)
	}

	return c.SendStatus(200) // file does not exist
}
