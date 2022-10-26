package main

import (
	"encoding/json"
	"flag"
	"io/ioutil"
	"log"
	"net/http"

	"golang.org/x/exp/slices"

	"github.com/gin-gonic/gin"
)

type TenantLicense struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Services    []string `json:"services"`
	PeriodStart string   `json:"period_start"`
	PeriodEnd   string   `json:"period_end"`
}

func loadMockData(filePath string) ([]TenantLicense, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal("Error when opening file: ", err)
		return nil, err
	}

	var tenants []TenantLicense
	err = json.Unmarshal(content, &tenants)
	if err != nil {
		log.Fatal("Error during Unmarshal(): ", err)
	}
	return tenants, nil
}

var (
	sourceFile = flag.String("source_file", "tenants.json", "Source file to load tenants from")
)

func main() {
	flag.Parse()

	r := gin.Default()
	tenants, err := loadMockData(*sourceFile)

	if err != nil {
		log.Fatalf("Exiting service since failed to load data")
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/licenses/:tenantID", func(c *gin.Context) {
		tenantID := c.Param("tenantID")
		index := slices.IndexFunc(tenants, func(t TenantLicense) bool {
			return t.ID == tenantID
		})

		if index == -1 {
			c.String(http.StatusNotFound, "Tenant not found")
		} else {
			c.JSON(http.StatusOK, gin.H{
				"tenant": &tenants[index],
			})
		}

	})
	r.Run()
}
