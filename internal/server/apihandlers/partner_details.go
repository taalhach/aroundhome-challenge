package apihandlers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/taalhach/aroundhome-challennge/internal/server/database/dbutils"
	"github.com/taalhach/aroundhome-challennge/internal/server/models"
	"github.com/taalhach/aroundhome-challennge/pkg/forms"
)

type partnerDetailsForm struct {
	Id int64 `param:"id" validate:"required"`
}

type partnerDetailsResponse struct {
	forms.BasicResponse
	Partner *dbutils.PartnerListItem
}

//PartnerDetails handler for to get page details
// @Summary Get page data's details
// @Description This API returns page data details.
// @Success 200 {object} pageDataResponse
// @Failure 404 {object} forms.BasicResponse
// @Param url path string false "URL"
// @Router /page_data/{url} [get]
func PartnerDetails(c echo.Context) error {
	form := new(partnerDetailsForm)
	if err := c.Bind(form); err != nil {
		return err
	}
	//form validation
	if err := c.Validate(form); err != nil {
		return err
	}

	// check if partner exists
	has, partner, err := dbutils.PartnerDetails(&models.Partner{Id: form.Id})
	if err != nil {
		return err
	}

	var (
		ret  partnerDetailsResponse
		code int
	)

	if has {
		ret.Success = true
		ret.Message = "get partner details successfully"
		ret.Partner = partner
		code = http.StatusOK
	} else {
		ret.Success = false
		ret.Message = fmt.Sprintf("partner %v not found", form.Id)
		code = http.StatusNotFound
	}

	return c.JSON(code, ret)
}
