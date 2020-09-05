package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type LinkType struct {
	gorm.Model
	Name  string `json:"name"`
	Links []Link
}

type Link struct {
	gorm.Model
	URL        string `json:"url"`
	LinkTypeID uint
	LinkType   *LinkType `json:"type"`
}

type Subject struct {
	gorm.Model
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	Projects []Project
}

type Channel struct {
	gorm.Model
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	URL      string `json:"url"`
	IsOnline bool   `json:"is_online"`
	Meets    []Meet
}

type Meet struct {
	gorm.Model
	Name      string    `json:"name"`
	Date      time.Time `json:"date"`
	ChannelID uint
	Channel   *Channel `json:"channel"`
	Done      bool     `json:"done"`
	Project   Project
	ProjectID uint
}

type Commit struct {
	gorm.Model
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	Solved    bool      `json:"solved"`
	SolvedAt  time.Time `json:"solved_at,omitempty"`
	LimitDate time.Time `json:"limit_date"`
	Project   Project
	ProjectID uint
}

// Milestone ...
type Milestone struct {
	gorm.Model
	Title     string    `json:"title"`
	Desc      string    `json:"desc"`
	FileURL   string    `json:"file_url"`
	Solved    bool      `json:"solved"`
	Date      time.Time `json:"date"`
	Project   Project
	ProjectID uint
}

type Progress struct {
	gorm.Model
	Name      string `json:"name"`
	Project   Project
	ProjectID uint
}

type ProjectState struct {
	gorm.Model
	Name string `json:"name"`
}

type ProjectType struct {
	gorm.Model
	Name string `json:"name"`
}

type Project struct {
	gorm.Model
	ProjectState *ProjectState `json:"state"`
	Title        string        `json:"title"`
	ProjectType  *ProjectType  `json:"project_type"`
	Desc         string        `json:"desc"`
	Authors      []*Student    `gorm:"many2many:project_authors" json:"authors"`
	Guides       []*Teacher    `gorm:"many2many:project_guides" json:"guides"`
	Links        []Link        `json:"links"`
	Subjects     []Subject     `gorm:"many2many:project_subjects" json:"subjects"`
	Meets        []Meet        `json:"meets"`
	Milestones   []Milestone   `json:"milestones"`
	Progress     []Progress    `json:"progress"`
	Tags         string        `json:"tags"`
	Commits      []Commit      `json:"commits"`
}
