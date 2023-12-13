package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jung-kurt/gofpdf"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

type Client struct {
	gorm.Model
	ClientShortName string `gorm:"type:varchar(50)"`
	ClientLongName  string `gorm:"type:varchar(50)"`
	ClientSurName   string `gorm:"type:varchar(50)"`
	Gender          string `gorm:"type:varchar(05)"`
	Salutation      string `gorm:"type:varchar(05)"`
	Language        string `gorm:"type:varchar(05)"`
	ClientDob       string `gorm:"type:varchar(8)"`
	ClientDod       string `gorm:"type:varchar(8)"`
	ClientEmail     string `gorm:"type:varchar(100)"`
	ClientMobCode   string `gorm:"type:varchar(05)"`
	ClientMobile    string `gorm:"type:varchar(20)"`
	ClientStatus    string `gorm:"type:varchar(05)"`
	ClientType      string `gorm:"type:varchar(01)"` // C CORPORATE I FOR INDIVIDUAL
	NationalId      string `gorm:"type:varchar(50);unique"`
	Nationality     string `gorm:"type:varchar(02)"`
}

func main() {
	// Database setup
	dsn := "root:root@tcp(127.0.0.1:3306)/policy?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Router setup
	router := gin.Default()
	router.GET("/clients", getClientsHandler)
	router.POST("/clients", createClient)
	router.GET("/generate-pdf-client", generatePDFClientHandler)
	port := 8080
	router.Run(fmt.Sprintf(":%d", port))
}

func getClientsHandler(c *gin.Context) {
	clients, err := getClients(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, clients)
}

func createClient(c *gin.Context) {
	var input Client
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&input)
	c.JSON(http.StatusOK, input)
}

func generatePDFClientHandler(c *gin.Context) {
	clients, err := getClients(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, "Generated Client")
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)

	// Initialize column widths
	colWidths := make([]float64, 6)

	// Add table headers
	pdf.SetFillColor(200, 220, 255)
	headers := []string{"Name", "Email", "Salutation", "Language", "Gender", "ClientDob"}
	for i, header := range headers {
		maxWidth := pdf.GetStringWidth(header) + 6 // Add padding
		for _, client := range clients {
			fieldValue := ""
			switch i {
			case 0:
				fieldValue = fmt.Sprintf("%s %s", client.ClientShortName, client.ClientLongName)
			case 1:
				fieldValue = client.ClientEmail
			case 2:
				fieldValue = client.Salutation
			case 3:
				fieldValue = client.Language
			case 4:
				fieldValue = client.Gender
			case 5:
				fieldValue = client.ClientDob
			}

			textWidth := pdf.GetStringWidth(fieldValue) + 6 // Add padding
			if textWidth > maxWidth {
				maxWidth = textWidth
			}
		}
		colWidths[i] = maxWidth
		pdf.CellFormat(maxWidth, 7, header, "1", 0, "C", true, 0, "")
	}
	pdf.Ln(7)

	// Add table rows
	pdf.SetFillColor(255, 255, 255)
	for _, client := range clients {
		for i := range headers {
			fieldValue := ""
			switch i {
			case 0:
				fieldValue = fmt.Sprintf("%s %s", client.ClientShortName, client.ClientLongName)
			case 1:
				fieldValue = client.ClientEmail
			case 2:
				fieldValue = client.Salutation
			case 3:
				fieldValue = client.Language
			case 4:
				fieldValue = client.Gender
			case 5:
				fieldValue = client.ClientDob
			}

			pdf.CellFormat(colWidths[i], 10, fieldValue, "1", 0, "L", false, 0, "")
		}
		pdf.Ln(10)
	}

	c.Header("Content-Type", "application/pdf")
	err = pdf.Output(c.Writer)
	if err != nil {
		log.Fatal(err)
	}
}




func getClients(c *gin.Context) ([]Client, error) {
	var client []Client
	result := db.Find(&client)
	if result.Error != nil {
		return nil, result.Error
	}
	return client, nil
}
