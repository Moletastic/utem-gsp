package services

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/Moletastic/utem-gsp/models"
	"github.com/Moletastic/utem-gsp/utils"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
)

type ProjectHandler struct {
	CRUDHandler
	Teacher      CRUDService
	Student      CRUDService
	Subject      CRUDService
	ProjectState CRUDService
	Preloads     []string
}

type ProjectPatchForm struct {
	Title          string `json:"title"`
	Desc           string `json:"desc"`
	ProjectStateID int64  `json:"project_state_id"`
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
	student := NewCrudService(
		&models.Student{},
		models.Student{},
		"access:student",
		[]string{},
		d,
	)
	subject := NewCrudService(
		&models.Student{},
		models.Student{},
		"project:subject",
		[]string{},
		d,
	)
	ph := &ProjectHandler{
		CRUDHandler: *crud,
		Teacher:     *teacher,
		Student:     *student,
		Subject:     *subject,
	}
	state := NewCrudService(
		&models.ProjectState{},
		models.ProjectState{},
		"project:state",
		[]string{},
		d,
	)
	ph.ProjectState = *state
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

func (ph *ProjectHandler) Patch(c echo.Context) error {
	req := new(CRUDReq)
	p, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	id := int64(p)
	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	jsonbody, err := json.Marshal(req.Data)
	form := new(ProjectPatchForm)
	if err = json.Unmarshal(jsonbody, &form); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	obj := new(models.Project)
	if err = ph.Service.GetByID(obj, id); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	obj.Title = form.Title
	obj.Desc = form.Desc
	state := new(models.ProjectState)
	if err = ph.ProjectState.GetByID(state, form.ProjectStateID); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	obj.ProjectState = *state
	if err = ph.Service.Update(obj, id); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	obj = new(models.Project)
	if err = ph.Service.GetByID(obj, id); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, obj)
}

func (ph *ProjectHandler) associateTeachers(p *models.Project) error {
	model := ph.Service.db.Model(p)
	teachers := make([]models.Teacher, 0)
	for _, guide := range p.Guides {
		t := models.Teacher{}
		if err := ph.Teacher.GetByID(&t, guide.ID); err != nil {
			return err
		}
		teachers = append(teachers, t)
	}
	if err := model.Association("Guides").Replace(teachers).Error; err != nil {
		return err
	}
	return nil
}

func (ph *ProjectHandler) associateStudents(p *models.Project) error {
	model := ph.Service.db.Model(p)
	authors := make([]models.Student, 0)
	for _, author := range p.Authors {
		a := models.Student{}
		if err := ph.Student.GetByID(&a, author.ID); err != nil {
			return err
		}
		authors = append(authors, a)
	}
	if err := model.Association("Authors").Replace(authors).Error; err != nil {
		return err
	}
	return nil
}

func (ph *ProjectHandler) associateSubjects(p *models.Project) error {
	model := ph.Service.db.Model(p)
	subjects := make([]models.Subject, 0)
	for _, subject := range p.Subjects {
		s := models.Subject{}
		if err := ph.Subject.GetByID(&s, subject.ID); err != nil {
			return err
		}
		subjects = append(subjects, s)
	}
	if err := model.Association("Subjects").Replace(subjects).Error; err != nil {
		return err
	}
	return nil
}

func (ph *ProjectHandler) Create(c echo.Context) error {
	req := new(CRUDReq)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	jsonbody, err := json.Marshal(req.Data)
	obj := new(models.Project)
	if err = json.Unmarshal(jsonbody, &obj); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	obj.InitGSP(ph.Service.EntityName)
	guides := obj.Guides
	authors := obj.Authors
	subjects := obj.Subjects
	obj.Guides = make([]models.Teacher, 0)
	obj.Authors = make([]models.Student, 0)
	obj.Subjects = make([]models.Subject, 0)
	obj.InitGSP(ph.Service.EntityName)
	if err := ph.Service.Create(obj); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	id := obj.ID
	obj.Guides = guides
	obj.Authors = authors
	obj.Subjects = subjects
	if err = ph.associateTeachers(obj); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err = ph.associateStudents(obj); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err = ph.associateSubjects(obj); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err = ph.Service.Update(obj, id); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	obj = new(models.Project)
	if err = ph.Service.GetByID(obj, id); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, obj)
}

func (ph *ProjectHandler) Update(c echo.Context) error {
	req := new(CRUDReq)
	p, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	id := int64(p)
	if err = c.Bind(req); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	jsonbody, err := json.Marshal(req.Data)
	obj := new(models.Project)
	if err = json.Unmarshal(jsonbody, &obj); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	obj.SetID(id)
	obj.InitGSP(ph.Service.EntityName)
	if err = ph.associateTeachers(obj); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err = ph.associateStudents(obj); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	if err = ph.associateSubjects(obj); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	obj = new(models.Project)
	if err = ph.Service.GetByID(obj, id); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, obj)
}

func (ph *ProjectHandler) List(c echo.Context) error {
	req := new(ListReq)
	teacher := models.Teacher{}
	ph.decodeListReq(c, req)
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(*utils.GSPClaim)
	if claims.User.Account.ID == 0 {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(errors.New("Empty Account ID")))
	}
	actype := claims.User.Account.AccountType
	if actype != "Admin" && actype != "Teacher" {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(errors.New("Unknown account type")))
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
