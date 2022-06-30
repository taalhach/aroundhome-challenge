package apihandlers

import (
	"fmt"
	"math"
	"net/http"

	"github.com/taalhach/aroundhome-challennge/pkg/forms"

	"github.com/taalhach/aroundhome-challennge/internal/server/database/dbutils"

	"github.com/labstack/echo/v4"
	"github.com/taalhach/aroundhome-challennge/internal/server/common"
)

type matchedPartnersListForm struct {
	Limit     int     `query:"limit"  json:"limit"`
	Page      int     `query:"page"  json:"page"`
	Latitude  float64 `query:"latitude"  validate:"required" json:"latitude"`
	Longitude float64 `query:"longitude" validate:"required" json:"longitude"`
	FloorArea float32 `query:"floor_area"`
}

type matchedPartnersListParams struct {
	Material float32 `param:"material"`
}
type partnersListResponse struct {
	common.BasicListRet
	Items []*dbutils.PartnerListItem `json:"items"`
}

//PartnersList finds best possible matched partners
// @Summary Get best possible matched partners
// @Description This API can be used to retrieve page data list,
// @Description which is filterable and provides capabilities on basis of all fields.
// @Success 200 {object} partnersListResponse
// @Failure 404 {object} forms.BasicResponse
// @Router /page_data [get]
func PartnersList(c echo.Context) error {
	form := new(matchedPartnersListForm)
	if err := c.Bind(form); err != nil {
		return err
	}

	if err := c.Validate(form); err != nil {
		return err
	}

	material := c.Param("material")
	basicList := forms.BasicList{
		Limit:   form.Limit,
		Page:    form.Limit,
		Filters: []string{fmt.Sprintf("material:eq:%s", material)},
	}
	basicList.AttachDefaults()

	items, total, err := dbutils.FindMatchedPartners(&basicList, form.Longitude, form.Latitude)
	if err != nil {
		return err
	}

	ret := partnersListResponse{
		BasicListRet: common.BasicListRet{
			Page:  form.Page,
			Pages: int(math.Ceil(float64(total) / float64(basicList.Limit))),
			Total: total,
		},
		Items: items,
	}
	return c.JSON(http.StatusOK, ret)
}