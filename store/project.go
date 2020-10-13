package store

import (
	"github.com/Moletastic/utem-gsp/models"
	"github.com/Moletastic/utem-gsp/services"
	"github.com/jinzhu/gorm"
)

type ProjectStore struct {
	Project *services.CRUDService
	Related []*services.CRUDHandler
	db      *gorm.DB
}

func NewProjectStore(db *gorm.DB) *ProjectStore {
	project := services.NewCrudService(
		&models.Project{},
		"project:project",
		[]string{
			"Authors",
			"Milestones",
			"Meets",
			"Meets.Channel",
			"Guides",
			"Guides.User",
			"Subjects",
			"Progress",
			"Commits",
			"ProjectType",
			"ProjectState",
			"Links",
			"Links.LinkType",
			"Reviews",
		},
		db,
	)
	commit := services.NewCrudService(
		&models.Commit{},
		"project:commit",
		[]string{},
		db,
	)
	meet := services.NewCrudService(
		&models.Meet{},
		"project:meet",
		[]string{"Channel"},
		db,
	)
	milestone := services.NewCrudService(
		&models.Milestone{},
		"project:milestone",
		[]string{},
		db,
	)
	subject := services.NewCrudService(
		&models.Subject{},
		"project:subject",
		[]string{},
		db,
	)
	progress := services.NewCrudService(
		&models.Progress{},
		"project:progress",
		[]string{},
		db,
	)
	channel := services.NewCrudService(
		&models.Channel{},
		"project:channel",
		[]string{},
		db,
	)
	link := services.NewCrudService(
		&models.Link{},
		"project:link",
		[]string{"LinkType"},
		db,
	)
	linktype := services.NewCrudService(
		&models.LinkType{},
		"project:linktype",
		[]string{},
		db,
	)
	rubric := services.NewCrudService(
		&models.Rubric{},
		"project:rubric",
		[]string{"Reviews"},
		db,
	)
	review := services.NewCrudService(
		&models.Review{},
		"project:review",
		[]string{"Rubric", "Reviewer", "Reviewer.User"},
		db,
	)
	ptype := services.NewCrudService(
		&models.ProjectType{},
		"project:type",
		[]string{},
		db,
	)
	pstate := services.NewCrudService(
		&models.ProjectState{},
		"project:state",
		[]string{},
		db,
	)
	related := []*services.CRUDHandler{
		services.NewCRUDHandler("project", project),     //0
		services.NewCRUDHandler("commit", commit),       // 1
		services.NewCRUDHandler("meet", meet),           // 2
		services.NewCRUDHandler("milestone", milestone), // 3
		services.NewCRUDHandler("subject", subject),     // 4
		services.NewCRUDHandler("progress", progress),   // 5
		services.NewCRUDHandler("channel", channel),     // 6
		services.NewCRUDHandler("link", link),           // 7
		services.NewCRUDHandler("linktype", linktype),   // 8
		services.NewCRUDHandler("rubric", rubric),       // 9
		services.NewCRUDHandler("review", review),       // 10
		services.NewCRUDHandler("ptype", ptype),
		services.NewCRUDHandler("pstate", pstate),
	}
	return &ProjectStore{
		Project: project,
		Related: related,
		db:      db,
	}
}
