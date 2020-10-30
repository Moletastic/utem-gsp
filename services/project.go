package services

import (
	"errors"
	"net/http"

	"github.com/Moletastic/utem-gsp/models"
	"github.com/Moletastic/utem-gsp/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type ProjectHandler struct {
	CRUDHandler
	Teacher  CRUDService
	Preloads []string
}

func NewProjectHandler(d *gorm.DB, s *CRUDService) *ProjectHandler {
	crud := NewCRUDHandler("project", s)
	teacher := NewCrudService(
		&models.Teacher{},
		models.Teacher{},
		"access:teacher",
		[]string{"Projects"},
		d,
	)
	ph := &ProjectHandler{
		CRUDHandler: *crud,
		Teacher:     *teacher,
	}
	ph.Preloads = []string{
		"Projects",
		"Projects.Authors",
		"Projects.Milestones",
		"Projects.Meets",
		"Projects.Meets.Channel",
		"Projects.Guides",
		"Projects.Guides.Account",
		"Projects.Subjects",
		"Projects.Progress",
		"Projects.Commits",
		"Projects.ProjectType",
		"Projects.ProjectState",
		"Projects.Links",
		"Projects.Links.LinkType",
		"Projects.Reviews",
	}
	return ph
}

func (ph *ProjectHandler) GetByID(c echo.Context) error {
	return ph.GetByID(c)
}

func (ph *ProjectHandler) List(c echo.Context) error {
	req := new(ListReq)
	teacher := models.Teacher{}
	ph.decodeListReq(c, req)
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*utils.GSPClaim)
	if claims.User.Account.ID == 0 {
		return c.JSON(http.StatusUnprocessableEntity, errors.New("Empty Account ID"))
	}
	actype := claims.User.Account.AccountType
	if actype != "Admin" && actype != "Teacher" {
		return c.JSON(http.StatusUnprocessableEntity, errors.New("Unknown account type"))
	}
	if actype == "Admin" {
		list, _, err := ph.Service.List(&req.Params)
		if err != nil {
			return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
		}
		return c.JSON(http.StatusOK, list)
	}
	d := ph.Teacher.db
	for _, preload := range ph.Preloads {
		d = d.Preload(preload)
	}
	err := d.Where("account_id = ?", claims.User.Account.ID).First(&teacher).Error
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, teacher.Projects)
}
