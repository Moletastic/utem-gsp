package services

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/Moletastic/utem-gsp/models"
	"github.com/Moletastic/utem-gsp/utils"
	"github.com/labstack/echo/v4"
	"github.com/mitchellh/mapstructure"
)

// CUReq for CreateUpdateReq
type CRUDReq struct {
	Data map[string]interface{} `json:"data"`
}

type ListReq struct {
	Params ListParams `json:"params"`
}

type ICRUDHandler interface {
	Create(c echo.Context) error
	Update(c echo.Context) error
	Delete(c echo.Context) error
	GetByID(c echo.Context) error
	List(c echo.Context) error
}

type CRUDHandler struct {
	Service *CRUDService
	Name    string
}

func NewCRUDHandler(n string, s *CRUDService) *CRUDHandler {
	return &CRUDHandler{
		Service: s,
		Name:    n,
	}
}

func (crud *CRUDHandler) decodeReq(c echo.Context, req *CRUDReq, data interface{}) error {
	if err := c.Bind(req); err != nil {
		return err
	}
	err := mapstructure.Decode(req.Data, data)
	if err != nil {
		return err
	}
	return nil
}

func (crud *CRUDHandler) decodeListReq(c echo.Context, req *ListReq) error {
	if err := c.Bind(req); err != nil {
		return err
	}
	return nil
}

func (crud *CRUDHandler) Create(c echo.Context) error {
	req := new(CRUDReq)
	obj := crud.Service.Model
	err := crud.decodeReq(c, req, &obj)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	obj.InitGSP(crud.Service.Entity)
	err = crud.Service.Create(obj)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	err = crud.Service.GetByID(obj.GetID(), obj)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, obj)
}

func (crud *CRUDHandler) Update(c echo.Context) error {
	req := new(CRUDReq)
	id, _ := strconv.Atoi(c.Param("id"))
	obj := crud.Service.Model
	err := crud.decodeReq(c, req, &obj)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	obj.InitGSP(crud.Service.Entity)
	err = crud.Service.Update(obj, int64(id))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, obj)
}

func (crud *CRUDHandler) Delete(c echo.Context) error {
	req := new(CRUDReq)
	obj := crud.Service.Model
	id, _ := strconv.Atoi(c.Param("id"))
	err := crud.decodeReq(c, req, &obj)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	err = crud.Service.Delete(obj, int64(id))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, id)
}

func (crud *CRUDHandler) GetByID(c echo.Context) error {
	var obj models.Model
	id, _ := strconv.Atoi(c.Param("id"))
	obj = crud.Service.GetModel()
	obj.SetID(0)
	c.Logger().Debug(obj.GetID())
	err := crud.Service.GetByID(int64(id), obj)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	obj.SetID(int64(id))
	return c.JSON(http.StatusOK, obj)
}

func (crud *CRUDHandler) List(c echo.Context) error {
	req := new(ListReq)
	crud.decodeListReq(c, req)
	fmt.Println(req)
	list, _, err := crud.Service.List(&req.Params)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, list)
}
