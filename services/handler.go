package services

import (
	"net/http"
	"strconv"

	"github.com/Moletastic/utem-gsp/decoder"
	"github.com/Moletastic/utem-gsp/utils"
	"github.com/labstack/echo/v4"
)

// CRUDReq for CreateUpdateReq
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
	decoder *decoder.GSPDecoder
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
	d, err := decoder.NewDecoder(data)
	if err != nil {
		return nil
	}
	if err = d.Decode(req.Data); err != nil {
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
	obj := crud.Service.GetNew()
	if err := crud.decodeReq(c, req, &obj); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	obj.InitGSP(crud.Service.EntityName)
	if err := crud.Service.Create(obj); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err := crud.Service.GetByID(obj, obj.GetID()); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, obj)
}

func (crud *CRUDHandler) Update(c echo.Context) error {
	req := new(CRUDReq)
	p, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	id := int64(p)
	obj := crud.Service.GetNew()
	if err = crud.decodeReq(c, req, &obj); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	obj.InitGSP(crud.Service.EntityName)
	if err = crud.Service.Update(obj, id); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err = crud.Service.GetByID(obj, id); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, obj)
}

func (crud *CRUDHandler) Delete(c echo.Context) error {
	req := new(CRUDReq)
	p, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	id := int64(p)
	obj := crud.Service.GetNew()
	if err = crud.decodeReq(c, req, &obj); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err = crud.Service.Delete(obj, id); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, id)
}

func (crud *CRUDHandler) GetByID(c echo.Context) error {
	p, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	id := int64(p)
	obj := crud.Service.GetNew()
	if err = crud.Service.GetByID(obj, id); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, obj)
}

func (crud *CRUDHandler) List(c echo.Context) error {
	req := new(ListReq)
	crud.decodeListReq(c, req)
	list, _, err := crud.Service.List(&req.Params)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, list)
}
