package manager

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/taalhach/aroundhome-challennge/internal/server/database"
	"github.com/taalhach/aroundhome-challennge/internal/server/models"
	"github.com/taalhach/aroundhome-challennge/pkg/items"
	"github.com/xuri/excelize/v2"
)

//CreateMockData create partners data from file at dataFilePath
// counts partners and if already exist and force flag is not set otherwise create data
func CreateMockData(db *database.DbSession, dataFilePath string, force bool) error {
	var partnersCount int64
	if err := db.Model(&models.Partner{}).Count(&partnersCount).Error; err != nil {
		return err
	}

	// check if data already exists and force flag is not set
	if partnersCount != 0 && !force {
		fmt.Println("data already exists so aborting")
		return nil
	}

	//open mock data xls file
	f, err := excelize.OpenFile(dataFilePath)
	if err != nil {
		return err
	}

	defer f.Close()
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		return err
	}

	first := true
	materials := []string{items.MaterialWood, items.MaterialTiles, items.MaterialCarpet}
	tx := db.Begin()
	// iterate rows
	for _, row := range rows {
		// skip for first row
		if first {
			first = !first
			continue
		}

		// not enough values for partner
		if len(row) < 3 {
			continue
		}

		partner := models.Partner{
			Name: fmt.Sprintf("aroundhome/%s", row[0]),
		}

		partner.Latitude, err = strconv.ParseFloat(row[1], 64)
		if err != nil {
			continue
		}

		partner.Longitude, err = strconv.ParseFloat(row[2], 64)
		if err != nil {
			continue
		}

		// random radius = (latitude + longitude)/10
		partner.Radius = int((partner.Longitude+partner.Latitude)/10) * 1000

		// 1 to 5 random rating
		partner.Rating = (rand.Float32() * (5 - 1)) + 1

		if err := tx.Create(&partner).Error; err != nil {
			tx.Rollback()
			return err
		}

		// randomly select partner experienced materials
		upperIndex := time.Now().UnixNano() % int64(len(materials))
		if upperIndex == 0 {
			upperIndex = 1
		}

		// create partner_materials entries
		for _, name := range materials[:upperIndex] {
			var material models.Material
			if err := tx.Where("lower(name) = ?", strings.ToLower(name)).Take(&material).Error; err != nil {
				tx.Rollback()
				return err
			}

			partnerMaterial := models.PartnerMaterial{
				MaterialID: material.Id,
				PartnerID:  partner.Id,
			}
			if err := tx.Create(&partnerMaterial).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}

	return tx.Commit().Error
}
