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

type Report struct {
	gorm.Model
	Title   string
	Content string
}

func main() {
	// Replace "user:password@tcp(host:port)/dbname?charset=utf8mb4&parseTime=True&loc=Local" with your MySQL connection string
	dsn := "root:root@tcp(127.0.0.1:3306)/policy?charset=utf8mb4&parseTime=True&loc=Local"

	// Connect to the MySQL database
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}

	// Migrate the schema
	db.AutoMigrate(&Report{})

	// Set up the Gin router
	router := gin.Default()

	// Define API routes
	router.GET("/reports", getReportsHandler)
	router.POST("/reports", createReport)
	router.GET("/generate-pdf-report", generatePDFReportHandler) // New endpoint for generating a PDF report

	// Run the server
	port := 8080
	router.Run(fmt.Sprintf(":%d", port))
}

// Handler to get all reports
func getReportsHandler(c *gin.Context) {
	reports, err := getReports(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, reports)
}

// Handler to create a new report
func createReport(c *gin.Context) {
	var input Report
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	db.Create(&input)
	c.JSON(http.StatusOK, input)
}

// Handler to generate a PDF report
// Handler to generate a PDF report
func generatePDFReportHandler(c *gin.Context) {
	reports, err := getReports(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPage()
	pdf.SetFont("Arial", "B", 16)
	pdf.Cell(0, 10, "Generated Report")
	pdf.Ln(10)
	pdf.SetFont("Arial", "", 12)

	for _, report := range reports {
		pdf.Cell(0, 10, fmt.Sprintf("Title: %s", report.Title))
		pdf.Ln(8)
		pdf.MultiCell(0, 10, fmt.Sprintf("Content: %s", report.Content), "TLR", "", false)
		pdf.Ln(8)
	}

	c.Header("Content-Type", "application/pdf")
	err = pdf.Output(c.Writer)
	if err != nil {
		log.Fatal(err)
	}
}


// Function to get all reports
func getReports(c *gin.Context) ([]Report, error) {
	var reports []Report
	result := db.Find(&reports)
	if result.Error != nil {
		return nil, result.Error
	}
	return reports, nil
}
