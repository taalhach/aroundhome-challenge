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
	Partner *dbutils.PartnerListItem `json:"partner"`
}

//PartnerDetails handler for getting partner details
// @Summary Get partner's details
// @Description This API returns partner details.
// @Success 200 {object} partnerDetailsResponse
// @Failure 404 {object} forms.BasicResponse
// @Param id path int 0 "Partner Id(example 272)"
// @Router /partners/{id} [get]
func PartnerDetails(c echo.Context) error {
	form := new(partnerDetailsForm)
	if err := c.Bind(form); err != nil {
		return err
	}
	//form validation
	if err := c.Validate(form); err != nil {
		return err
	}

	// check if partner exists and fetch details
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
