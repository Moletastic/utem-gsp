package models

import (
	"time"
)

type LinkType struct {
	GSPModel
	Name string `mapstructure:"name" json:"name"`
	Icon string `mapstructure:"icon" json:"icon"`
}

func (l LinkType) Bind(v interface{}) {
	v = LinkType{}
}

func (l LinkType) New() Model {
	return &LinkType{}
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
	LinkTypeID int64    `mapstructure:"link_type_id" json:"link_type_id"`
	LinkType   LinkType `mapstructure:"link_type" json:"link_type"`
	Project    Project  `mapstructure:"project" json:"project" gorm:"foreignKey:ProjectID"`
	ProjectID  int64    `mapstructure:"project_id" json:"project_id,omitempty" gorm:"column:project_id"`
}

func (l Link) Bind(v interface{}) {
	v = Link{}
}

func (l Link) New() Model {
	return &Link{}
}

func NewLink(url string, tid int64, pid int64) Link {
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
	Name string `json:"name" mapstructure:"name"`
	Icon string `json:"icon" mapstructure:"icon"`
}

func (s Subject) Bind(v interface{}) {
	v = Subject{}
}

func (s Subject) New() Model {
	return &Subject{}
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

func (c Channel) Bind(v interface{}) {
	v = Channel{}
}

func (c Channel) New() Model {
	return &Channel{}
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
	ChannelID int64     `mapstructure:"channel_id" json:"channel_id"`
	Channel   *Channel  `mapstructure:"channel" json:"channel"`
	Done      bool      `mapstructure:"done" json:"done"`
	Project   Project   `mapstructure:"project" json:"project"`
	ProjectID int64     `mapstructure:"project_id" json:"project_id"`
}

func (m Meet) Bind(v interface{}) {
	v = Meet{}
}

func (m Meet) New() Model {
	return &Meet{}
}

func NewMeet(n string, d time.Time, chid int64, pid int64) Meet {
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
	ProjectID int64      `json:"project_id" mapstructure:"project_id"`
}

func (c Commit) Bind(v interface{}) {
	v = Commit{}
}

func (c Commit) New() Model {
	return &Commit{}
}

func NewCommit(t string, limit time.Time, pid int64) Commit {
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
	ProjectID int64     `json:"project_id" mapstructure:"project_id"`
}

func (m Milestone) Bind(v interface{}) {
	v = Milestone{}
}

func (m Milestone) New() Model {
	return &Milestone{}
}

func NewMilestone(t string, d time.Time, pid int64) Milestone {
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
	ProjectID int64   `json:"project_id" mapstructure:"project_id"`
}

func (p Progress) Bind(v interface{}) {
	v = Progress{}
}

func (Progress) New() Model {
	return &Progress{}
}

func NewProgress(n string, pid int64) Progress {
	p := Progress{
		Name:      n,
		ProjectID: pid,
	}
	p.InitGSP("project:progress")
	return p
}

type ProjectState struct {
	GSPModel
	Name string `json:"name" mapstructure:"name"`
}

func (s ProjectState) Bind(v interface{}) {
	v = ProjectState{}
}

func (ProjectState) New() Model {
	return &ProjectState{}
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

func (t ProjectType) Bind(v interface{}) {
	v = ProjectType{}
}

func (ProjectType) New() Model {
	return &ProjectType{}
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

func (r Rubric) Bind(v interface{}) {
	v = Rubric{}
}

func (r Rubric) New() Model {
	return &Rubric{}
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
	RubricID  int64   `json:"rubric_id" mapstructure:"rubric_id"`
	Rubric    Rubric  `mapstructure:"rubric" json:"rubric"`
	Project   Project `json:"project" mapstructure:"project"`
	ProjectID int64   `json:"project_id" mapstructure:"project_id"`
	FileURL   string  `mapstructure:"file_url" json:"file_url"`
	// JSON String
	Score      string  `mapstructure:"score" json:"score"`
	ReviewerID int64   `json:"reviewer_id" mapstructure:"reviewer_id"`
	Reviewer   Teacher `mapstructure:"reviewer" json:"reviewer"`
}

func (r Review) New() Model {
	return &Review{}
}

func (r Review) Bind(v interface{}) {
	v = Review{}
}

func NewReview(n string, rid int64, pid int64, url string, rvid int64, score string) Review {
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
	ProjectStateID int64        `gorm:"column:project_state_id" json:"project_state_id,omitempty" mapstructure:"project_state_id"`
	Title          string       `json:"title" mapstructure:"title"`
	ProjectTypeID  int64        `json:"project_type_id" mapstructure:"project_type_id"`
	ProjectType    ProjectType  `json:"project_type" mapstructure:"project_type"`
	Desc           string       `json:"desc" mapstructure:"desc"`
	Authors        []Student    `gorm:"many2many:project_authors;" json:"authors" mapstructure:"authors"`
	Guides         []Teacher    `gorm:"many2many:project_guides" json:"guides" mapstructure:"guides"`
	Links          []Link       `json:"links" mapstructure:"links"`
	Subjects       []Subject    `gorm:"many2many:project_subjects" json:"subjects" mapstructure:"subjects"`
	Meets          []Meet       `json:"meets" mapstructure:"meets"`
	Milestones     []Milestone  `json:"milestones" mapstructure:"milestones"`
	Progress       []Progress   `json:"progress" mapstructure:"progress"`
	Tags           string       `json:"tags" mapstructure:"tags"`
	Commits        []Commit     `json:"commits" mapstructure:"commits"`
	Reviews        []Review     `json:"reviews" mapstructure:"reviews"`
}

func (p Project) Bind(v interface{}) {
	v = Project{}
}

func (p Project) New() Model {
	return &Project{}
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
