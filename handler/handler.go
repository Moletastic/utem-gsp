package handler

import (
	"github.com/Moletastic/utem-gsp/decoder"
	"github.com/Moletastic/utem-gsp/store"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	AccStore store.AccessStore
	EduStore store.EducationStore
	ProStore store.ProjectStore
}

// CRUDReq for CreateUpdateReq
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

func NewHandler(as store.AccessStore, es store.EducationStore, ps store.ProjectStore) *Handler {
	return &Handler{
		AccStore: as,
		EduStore: es,
		ProStore: ps,
	}
}

func (h *Handler) DecodeCRUDReq(c echo.Context, req *CRUDReq, data interface{}) error {
	if err := c.Bind(req); err != nil {
		return err
	}
	d, err := decoder.NewDecoder(data)
	if err != nil {
		return err
	}
	err = d.Decode(req.Data)
	return nil
}
