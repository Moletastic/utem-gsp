package models

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

type Model interface {
	InitGSP(t string)
	GetID() uint
	GetUID() string
}

type CommonModel struct {
	ID        uint       `gorm:"primary_key" mapstructure:"id" json:"id"`
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

// InitGSP initialize an GSPStructure
func (gsp *GSPModel) InitGSP(t string) {
	rand := strconv.Itoa(rand.Intn(9999-1000) + 1000)
	gsp.Entity = t
	gsp.UID = fmt.Sprintf("%s-%s", t, rand)
	gsp.IsValid = true
}

func (gsp *GSPModel) GetID() uint {
	return gsp.ID
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
