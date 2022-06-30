package dbutils

import (
	"github.com/taalhach/aroundhome-challennge/internal/server/common"
	"github.com/taalhach/aroundhome-challennge/internal/server/database"
	"github.com/taalhach/aroundhome-challennge/internal/server/models"
	"github.com/taalhach/aroundhome-challennge/pkg/forms"
)

type PartnerListItem struct {
	Id        int64   `gorm:"primaryKey" json:"id"`
	Name      string  `gorm:"column:name" json:"name"`
	Latitude  float64 `gorm:"column:latitude" json:"latitude"`
	Longitude float64 `gorm:"column:longitude" json:"longitude"`
	Distance  float64 `gorm:"column:distance" json:"distance"`
	Rating    float32 `gorm:"column:rating" json:"rating"`
}

//FindMatchedPartners finds partners matched with customer, supports pagination
func FindMatchedPartners(form *forms.BasicList, longitude, latitude float64) ([]*PartnerListItem, int64, error) {
	var (
		err   error
		total int64
		items []*PartnerListItem
	)
	// prepare query
	query := database.Db.Model(&models.Partner{}).
		Joins("INNER JOIN partner_materials AS pm ON pm.partner_id = partners.id").
		Joins("INNER JOIN materials ON materials.id = pm.material_id").
		Where("ST_DWithin(ST_SetSRID(geom, 4326)::geography, ST_MakePoint(?,?, 4326)::geography, radius) IS TRUE", longitude, latitude)
	columns := map[string]string{
		"rating":   "partners.rating",
		"distance": "distance",
		"material": "materials.name",
	}

	// prepare list params
	listParams := common.BasicList{
		Limit:   form.Limit,
		Page:    form.Page,
		Filters: form.Filters,
		Query:   query,
		Columns: columns,
	}

	// first skip so that we can get count of all first
	listParams.SkipPagination = true
	listParams.Query, err = listParams.PrepareSql()
	if err != nil {
		return nil, 0, err
	}

	// count total items
	if err = listParams.Query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// apply pagination by adding limit and offset
	listParams.Query = listParams.Paginate()

	const selectColumns = "partners.id, partners.name, latitude, longitude, rating, ST_DistanceSphere(ST_MakePoint(longitude, latitude),ST_MakePoint(?, ?)) AS distance"
	if err := listParams.Query.Select(selectColumns, longitude, latitude).Order("partners.rating DESC, distance DESC").Find(&items).Error; err != nil {
		return nil, 0, err
	}

	return items, total, nil
}
