package models

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"reflect"
	"strconv"
	"time"
)

type Model interface {
	InitGSP(t string)
	GetID() int64
	GetUID() string
	SetID(id int64)
	Bind(v interface{})
}

type Entity interface {
	Bind(v interface{})
	New() Model
}

type CommonModel struct {
	ID        int64      `gorm:"primaryKey" mapstructure:"id" json:"id"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type GSPModel struct {
	CommonModel
	Entity  string `mapstructure:"entity" json:"entity"`
	UID     string `json:"uid" mapstructure:"uid"`
	IsValid bool   `json:"is_valid" gorm:"default:1"`
}

func (gsp *GSPModel) Clear() {
	p := reflect.ValueOf(gsp).Elem()
	p.Set(reflect.Zero(p.Type()))
}

func (gsp GSPModel) Bind(v interface{}) {
	v = GSPModel{}
}

func (gsp GSPModel) New() *GSPModel {
	return &GSPModel{}
}

// InitGSP initialize an GSPStructure
func (gsp *GSPModel) InitGSP(t string) {
	rand := strconv.Itoa(rand.Intn(9999-1000) + 1000)
	gsp.Entity = t
	gsp.UID = fmt.Sprintf("%s-%s", t, rand)
	gsp.IsValid = true
}

func (gsp *GSPModel) GetID() int64 {
	return gsp.ID
}

func (gsp *GSPModel) SetID(id int64) {
	gsp.ID = id
}

func (gsp *GSPModel) GetUID() string {
	return gsp.UID
}

func (gsp *GSPModel) ToString() string {
	s, _ := json.MarshalIndent(gsp, "", "  ")
	return string(s)
}

// SetInvalid invalids object
func (gsp *GSPModel) SetInvalid() {
	gsp.IsValid = false
}
