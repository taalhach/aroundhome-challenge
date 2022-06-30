package apihandlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/taalhach/aroundhome-challennge/internal/server/common"
	"github.com/taalhach/aroundhome-challennge/internal/server/models"
	"github.com/taalhach/aroundhome-challennge/pkg/forms"
)

type PageDataListResponse struct {
	common.BasicListRet
	Items []*models.Partner `json:"items"`
}

//PageDataList finds slowest queries
// @Summary Get page data list
// @Description This API can be used to retrieve page data list,
// @Description which is filterable and provides capabilities on basis of all fields.
// @Success 200 {object} pageDataResponse
// @Failure 404 {object} forms.BasicResponse
// @Router /page_data [get]
func PageDataList(c echo.Context) error {
	form := new(forms.BasicList)
	if err := c.Bind(form); err != nil {
		return err
	}

	// load default params
	// find statements
	//items, total, err := db.FindPagesData(form)
	//if err != nil {
	//	return err
	//}

	ret := PageDataListResponse{
		BasicListRet: common.BasicListRet{
			Page: form.Page,
			//Pages: int(math.Ceil(float64(total) / float64(form.Limit))),
			//Total: total,
		},
		//Items: items,
	}
	return c.JSON(http.StatusOK, ret)
}
