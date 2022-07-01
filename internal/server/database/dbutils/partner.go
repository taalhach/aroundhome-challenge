package dbutils

import (
	"github.com/lib/pq"
	"github.com/taalhach/aroundhome-challennge/internal/server/common"
	"github.com/taalhach/aroundhome-challennge/internal/server/database"
	"github.com/taalhach/aroundhome-challennge/internal/server/models"
	"gorm.io/gorm"
)

type PartnerListItem struct {
	Id        int64          `gorm:"primaryKey" json:"id"`
	Name      string         `gorm:"column:name" json:"name"`
	Latitude  float64        `gorm:"column:latitude" json:"latitude"`
	Longitude float64        `gorm:"column:longitude" json:"longitude"`
	Distance  float64        `gorm:"column:distance" json:"distance"`
	Rating    float32        `gorm:"column:rating" json:"rating"`
	Radius    int            `gorm:"column:radius" json:"radius_in_meters"`
	Materials pq.StringArray `gorm:"column:materials;type:text[]" json:"materials"`
}

//FindMatchedPartners finds partners matched with customer, supports pagination
func FindMatchedPartners(listParams *common.BasicList, longitude, latitude float64) ([]*PartnerListItem, int64, error) {
	var (
		err   error
		total int64
		items []*PartnerListItem
	)
	// prepare query
	query := database.Db.Model(&models.Partner{}).
		Joins("INNER JOIN partner_materials AS pm ON pm.partner_id = partners.id").
		Joins("INNER JOIN materials ON materials.id = pm.material_id").
		Where("ST_DWithin(ST_SetSRID(geom, 4326)::geography, ST_MakePoint(?,?, 4326)::geography, radius) IS TRUE", longitude, latitude).
		Group("partners.id")
	columns := map[string]string{
		"rating":   "partners.rating",
		"distance": "distance",
		"material": "materials.name",
	}

	listParams.Columns = columns

	// first skip so that we can get count of all first
	listParams.SkipPagination = true
	query, err = listParams.PrepareSql(query)
	if err != nil {
		return nil, 0, err
	}

	// count total items
	if err = query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// apply pagination by adding limit and offset
	query = listParams.Paginate(query)

	const selectColumns = "partners.id, partners.name, latitude, longitude, rating, radius, ST_DistanceSphere(ST_MakePoint(longitude, latitude),ST_MakePoint(?, ?)) AS distance, string_to_array(STRING_AGG(materials.name,','), ',') AS materials"
	if err := query.Select(selectColumns, longitude, latitude).Order("partners.rating DESC, distance").Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, total, nil
}

//PartnerDetails get partner details
func PartnerDetails(target *models.Partner) (bool, *PartnerListItem, error) {
	var partner PartnerListItem
	if err := database.Db.Model(&models.Partner{}).
		Joins("INNER JOIN partner_materials AS pm ON pm.partner_id = partners.id").
		Joins("INNER JOIN materials ON materials.id = pm.material_id").
		Where(target).Select("partners.id,partners.name,partners.latitude,partners.longitude,radius, partners.rating, string_to_array(STRING_AGG(materials.name,','), ',') AS materials").
		Group("partners.id").Take(&partner).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil, nil
		}

		return false, nil, err
	}

	return true, &partner, nil
}
