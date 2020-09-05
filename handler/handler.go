package handler

import (
	"net/http"

	"github.com/Moletastic/utem-gsp/models"
	"github.com/Moletastic/utem-gsp/store"
	"github.com/Moletastic/utem-gsp/utils"
	"github.com/labstack/echo/v4"
)

type Handler struct {
	accessStore store.AccessStore
	eduStore    store.EducationStore
	proStore    store.ProjectStore
}

func NewHandler(as store.AccessStore, es store.EducationStore, ps store.ProjectStore) *Handler {
	return &Handler{
		accessStore: as,
		eduStore:    es,
		proStore:    ps,
	}
}

func (h *Handler) ListTeachers(c echo.Context) error {
	teachers, _, err := h.eduStore.ListTeachers()
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, teachers)
}

func (h *Handler) CreateProject(c echo.Context) error {

	project := models.Project{
		Title: "project-1",
		Authors: []*models.Student{
			{
				Career: &models.Career{
					Code: 21041,
					Name: "Ing civil",
				},
				EntryYear: 2010,
				FirstName: "Yeikeb",
				LastName:  "Romero",
				RUT:       "195239525",
			},
		},
		Tags: "vue;golang;rest",
	}
	err := h.proStore.CreateProject(&project)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, "")
}

func (h *Handler) ListProjects(c echo.Context) error {
	projects, _, err := h.proStore.ListProjects(&store.ListConf{
		Limit: 2,
	})
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}
	return c.JSON(http.StatusOK, projects)
}
