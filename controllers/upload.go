package controllers

import (
	"os"
	"time"

	beego "github.com/beego/beego/v2/server/web"
	"github.com/google/uuid"
)

type UploadsController struct {
	beego.Controller
}

func (c *UploadsController) Get() {
	c.TplName = "upload.html"
	return
}

func (c *UploadsController) Nmap() {

	file, _, err := c.GetFile("nmap")
	if err != nil {
		c.Ctx.WriteString("Error in file upload: " + err.Error())
		return
	}
	defer file.Close()

	// Generate a random filename
	uniqueFileName := GenerateRandomFilename()

	// Define the full path
	tempFilePath := "uploads/" + uniqueFileName + ".xml"

	// Save the file
	err = c.SaveToFile("nmap", tempFilePath)
	if err != nil {
		c.Ctx.WriteString("Error saving file: " + err.Error())
		return
	}

	// Process the file with ParseNmap function
	err = ParseNmapScan(tempFilePath)
	if err != nil {
		c.Ctx.WriteString("Error parsing the nmap scan: " + err.Error())
		return
	}

	// Delete the file after processing
	err = os.Remove(tempFilePath)
	if err != nil {
		c.Ctx.WriteString("Error deleting file: " + err.Error())
		return
	}

	c.Ctx.WriteString("File processed successfully")
}

//////////////////////////////////////////////////////////////////////////////////////////////////
//////////////////////////////////////// HELPER FUNCTIONS ////////////////////////////////////////
//////////////////////////////////////////////////////////////////////////////////////////////////

func GenerateRandomFilename() string {
	timestamp := time.Now().Format("20060102150405")
	uuidStr := uuid.New().String()
	return timestamp + "_" + uuidStr
}