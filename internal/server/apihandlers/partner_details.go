package apihandlers

import (
	"github.com/labstack/echo/v4"
)

type pageDataDetailsForm struct {
	Url string `param:"url" validate:"required"`
}

//PageDataDetails handler for to get page details
// @Summary Get page data's details
// @Description This API returns page data details.
// @Success 200 {object} pageDataResponse
// @Failure 404 {object} forms.BasicResponse
// @Param url path string false "URL"
// @Router /page_data/{url} [get]
func PageDataDetails(c echo.Context) error {
	form := new(pageDataDetailsForm)
	if err := c.Bind(form); err != nil {
		return err
	}
	//form validation
	if err := c.Validate(form); err != nil {
		return err
	}

	// check if page data exists
	//has, pageData, err := db.PageData(&models.PageData{Url: form.Url})
	//if err != nil {
	//	return err
	//}

	var (
		ret  interface{}
		code int
	)

	//if has {
	//	ret.Success = true
	//	ret.Message = "get page data successfully"
	//	ret.Page = pageData
	//	code = http.StatusOK
	//} else {
	//	ret.Success = false
	//	ret.Message = fmt.Sprintf("page %v not found", form.Url)
	//	code = http.StatusNotFound
	//}

	return c.JSON(code, ret)
}
