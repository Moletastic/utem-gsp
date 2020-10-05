package models

import (
	"time"
)

type LinkType struct {
	GSPModel
	Name  string `mapstructure:"name" json:"name"`
	Icon  string `mapstructure:"icon" json:"icon"`
	Links []Link `mapstructure:"link" json:"link"`
}

func NewLinkType(n string, i string) LinkType {
	lt := LinkType{
		Name: n,
		Icon: i,
	}
	lt.InitGSP("project:link_type")
	return lt
}

type Link struct {
	GSPModel
	URL        string   `mapstructure:"url" json:"url"`
	LinkTypeID uint     `mapstructure:"link_type_id" json:"link_type_id"`
	LinkType   LinkType `mapstructure:"link_type" json:"link_type"`
	Project    Project  `mapstructure:"project" json:"project" gorm:"foreignKey:ProjectID"`
	ProjectID  uint     `mapstructure:"project_id" json:"project_id,omitempty" gorm:"column:project_id"`
}

func NewLink(url string, tid uint, pid uint) Link {
	l := Link{
		URL:        url,
		LinkTypeID: tid,
		ProjectID:  pid,
	}
	l.InitGSP("project:link")
	return l
}

type Subject struct {
	GSPModel
	Name     string    `json:"name" mapstructure:"name"`
	Icon     string    `json:"icon" mapstructure:"icon"`
	Projects []Project `gorm:"many2many:project_subjects" json:"projects" mapstructure:"projects"`
}

func NewSubject(n string, i string) Subject {
	s := Subject{
		Name: n,
		Icon: i,
	}
	s.InitGSP("project:subject")
	return s
}

type Channel struct {
	GSPModel
	Name     string `json:"name" mapstructure:"name"`
	Icon     string `mapstructure:"icon" json:"icon"`
	URL      string `mapstructure:"url" json:"url"`
	IsOnline bool   `mapstructure:"is_online" json:"is_online"`
	Meets    []Meet `mapstructure:"meets" json:"meets"`
}

func NewChannel(n string, i string, online bool) Channel {
	c := Channel{
		Name:     n,
		Icon:     i,
		IsOnline: online,
	}
	c.InitGSP("project:channel")
	return c
}

type Meet struct {
	GSPModel
	Name      string    `mapstructure:"name" json:"name"`
	Date      time.Time `mapstructure:"date" json:"date"`
	ChannelID uint      `mapstructure:"channel_id" json:"channel_id"`
	Channel   *Channel  `mapstructure:"channel" json:"channel"`
	Done      bool      `mapstructure:"done" json:"done"`
	Project   Project   `mapstructure:"project" json:"project"`
	ProjectID uint      `mapstructure:"project_id" json:"project_id"`
}

func NewMeet(n string, d time.Time, chid uint, pid uint) Meet {
	m := Meet{
		Name:      n,
		Date:      d,
		ChannelID: chid,
		ProjectID: pid,
	}
	m.InitGSP("project:meet")
	return m
}

type Commit struct {
	GSPModel
	Title     string     `json:"title" mapstructure:"title"`
	Desc      string     `json:"desc" mapstructure:"desc"`
	Solved    bool       `json:"solved" mapstructure:"solved"`
	SolvedAt  *time.Time `json:"solved_at,omitempty" mapstructure:"solved_at"`
	LimitDate *time.Time `json:"limit_date" mapstructure:"limit_date"`
	Project   Project    `json:"project" mapstructure:"project"`
	ProjectID uint       `json:"project_id" mapstructure:"project_id"`
}

func NewCommit(t string, limit time.Time, pid uint) Commit {
	c := Commit{
		Title:     t,
		LimitDate: &limit,
		ProjectID: pid,
	}
	c.InitGSP("project:commit")
	return c
}

// Milestone ...
type Milestone struct {
	GSPModel
	Title     string    `json:"title" mapstructure:"title"`
	Desc      string    `json:"desc" mapstructure:"desc"`
	FileURL   string    `json:"file_url" mapstructure:"file_url"`
	Solved    bool      `json:"solved" mapstructure:"solved"`
	Date      time.Time `json:"date" mapstructure:"date"`
	Project   Project   `json:"project" mapstructure:"project"`
	ProjectID uint      `json:"project_id" mapstructure:"project_id"`
}

func NewMilestone(t string, d time.Time, pid uint) Milestone {
	m := Milestone{
		Title:     t,
		Date:      d,
		ProjectID: pid,
	}
	m.InitGSP("project:milestone")
	return m
}

type Progress struct {
	GSPModel
	Name      string  `json:"name" mapstructure:"name"`
	Project   Project `json:"project" mapstructure:"project"`
	ProjectID uint    `json:"project_id" mapstructure:"project_id"`
}

func NewProgress(n string, pid uint) Progress {
	p := Progress{
		Name:      n,
		ProjectID: pid,
	}
	p.InitGSP("project:progress")
	return p
}

type ProjectState struct {
	GSPModel
	Name     string    `json:"name" mapstructure:"name"`
	Projects []Project `json:"projects" mapstructure:"projects"`
}

func NewProjectState(n string) ProjectState {
	s := ProjectState{
		Name: n,
	}
	s.InitGSP("project:state")
	return s
}

type ProjectType struct {
	GSPModel
	Name     string    `json:"name" mapstructure:"name"`
	Projects []Project `json:"projects" mapstructure:"projects"`
}

func NewProjectType(n string) ProjectType {
	t := ProjectType{
		Name: n,
	}
	t.InitGSP("project:type")
	return t
}

type Rubric struct {
	GSPModel
	Name    string   `json:"name" mapstructure:"name"`
	FileURL string   `json:"file_url" mapstructure:"file_url"`
	Reviews []Review `gorm:"->" json:"reviews" mapstructure:"reviews"`
}

func NewRubric(n string, url string) Rubric {
	r := Rubric{
		Name:    n,
		FileURL: url,
	}
	r.InitGSP("project:rubric")
	return r
}

type Review struct {
	GSPModel
	Name      string  `mapstructure:"name" json:"name"`
	RubricID  uint    `json:"rubric_id" mapstructure:"rubric_id"`
	Rubric    Rubric  `mapstructure:"rubric" json:"rubric"`
	Project   Project `json:"project" mapstructure:"project"`
	ProjectID uint    `json:"project_id" mapstructure:"project_id"`
	FileURL   string  `mapstructure:"file_url" json:"file_url"`
	// JSON String
	Score      string  `mapstructure:"score" json:"score"`
	ReviewerID uint    `json:"reviewer_id" mapstructure:"reviewer_id"`
	Reviewer   Teacher `mapstructure:"reviewer" json:"reviewer"`
}

func NewReview(n string, rid uint, pid uint, url string, rvid uint, score string) Review {
	r := Review{
		Name:       n,
		RubricID:   rid,
		ProjectID:  pid,
		FileURL:    url,
		ReviewerID: rvid,
		Score:      score,
	}
	r.InitGSP("project:review")
	return r
}

type Project struct {
	GSPModel
	ProjectState   ProjectState `gorm:"foreignKey:ProjectStateID" json:"project_state" mapstructure:"project_state"`
	ProjectStateID uint         `gorm:"column:project_state_id" json:"project_state_id,omitempty" mapstructure:"project_state_id"`
	Title          string       `json:"title" mapstructure:"title"`
	ProjectTypeID  uint         `json:"project_type_id" mapstructure:"project_type_id"`
	ProjectType    ProjectType  `json:"project_type" mapstructure:"project_type"`
	Desc           string       `json:"desc" mapstructure:"desc"`
	Authors        []Student    `gorm:"many2many:project_authors" json:"authors" mapstructure:"authors"`
	Guides         []Teacher    `gorm:"many2many:project_guides" json:"guides" mapstructure:"guides"`
	Links          []Link       `gorm:"->" json:"links" mapstructure:"links"`
	Subjects       []Subject    `gorm:"many2many:project_subjects" json:"subjects" mapstructure:"subjects"`
	Meets          []Meet       `json:"meets" mapstructure:"meets"`
	Milestones     []Milestone  `json:"milestones" mapstructure:"milestones"`
	Progress       []Progress   `json:"progress" mapstructure:"progress"`
	Tags           string       `json:"tags" mapstructure:"tags"`
	Commits        []Commit     `json:"commits" mapstructure:"commits"`
	Reviews        []Review     `json:"reviews" mapstructure:"reviews"`
}

func NewProject(title string, authors []Student, guides []Teacher, subjects []Subject, ptype ProjectType) Project {
	p := Project{
		Title:         title,
		Authors:       authors,
		Guides:        guides,
		Subjects:      subjects,
		ProjectTypeID: ptype.ID,
	}
	p.InitGSP("project:project")
	return p
}