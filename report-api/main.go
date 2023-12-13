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
	
	dsn := "root:root@tcp(127.0.0.1:3306)/policy?charset=utf8mb4&parseTime=True&loc=Local"

	
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	
	// db.AutoMigrate(&Client{})

	
	router := gin.Default()

	
	router.GET("/clients", getClientsHandler)
	router.POST("/clients", createClient)
	router.GET("/generate-pdf-client", generatePDFClientHandler) // New endpoint for generating a PDF client

	
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

	for _, client := range clients {
		pdf.Cell(0, 10, fmt.Sprintf("Name: %s %s", client.ClientShortName,client.ClientLongName))
		pdf.Ln(8)
		pdf.MultiCell(0, 10, fmt.Sprintf("ClientEmail: %s", client.ClientEmail), "TLR", "", false)
		pdf.MultiCell(0, 10, fmt.Sprintf("Salutation: %s", client.Salutation), "TLR", "", false)
		pdf.MultiCell(0, 10, fmt.Sprintf("Language: %s", client.Language), "TLR", "", false)
		pdf.MultiCell(0, 10, fmt.Sprintf("Gender: %s", client.Gender), "TLR", "", false)
		pdf.MultiCell(0, 10, fmt.Sprintf("ClientDob: %s", client.ClientDob), "TLR", "", false)
		pdf.Ln(8)
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
