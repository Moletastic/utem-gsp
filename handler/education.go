package handler

import (
	"fmt"
	"net/http"

	"github.com/Moletastic/utem-gsp/models"
	"github.com/Moletastic/utem-gsp/utils"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

// CUReq for CreateUpdateReq
type CRUDReq struct {
	Data map[string]interface{} `json:"data"`
}

type DeleteReq struct {
	Data struct {
		Entity string `json:"entity"`
		ID     int    `json:"ID"`
		UID    string `json:"uid"`
	} `json:"data"`
}

func (h *Handler) CreateCareer(c echo.Context) error {
	req := new(CRUDReq)
	cr := new(models.Career)
	err := h.DecodeCRUDReq(c, req, cr)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	fmt.Println(utils.Pretty(req))
	fmt.Println(utils.Pretty(cr))
	entity := req.Data["entity"].(string)
	err = h.CreateEduEntity(cr, entity)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	car := h.eduStore.GetCareerByCode(cr.Code)
	if car == nil {
		return c.JSON(http.StatusForbidden, utils.AccessForbidden())
	}
	return c.JSON(http.StatusOK, car)
}

func (h *Handler) CreateEduEntity(data models.Model, t string) error {
	data.InitGSP(t)
	return h.eduStore.CreateEntity(data)
}

func (h *Handler) DecodeCRUDReq(c echo.Context, req *CRUDReq, data interface{}) error {
	if err := c.Bind(req); err != nil {
		return err
	}
	err := mapstructure.Decode(req.Data, data)
	if err != nil {
		return err
	}
	fmt.Println(data)
	return nil
}

func (h *Handler) ListCareers(c echo.Context) error {
	projects, _, err := h.eduStore.ListCareers()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, projects)
}
