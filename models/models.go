package models

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Model interface {
	InitGSP(t string)
}

type GSPModel struct {
	Entity  string `mapstructure:"entity" json:"entity"`
	UID     string `json:"uid" mapstructure:"uid"`
	IsValid bool   `json:"is_valid" gorm:"default:1"`
}

func (gsp *GSPModel) InitGSP(t string) {
	rand := strconv.Itoa(rand.Intn(9999-1000) + 1000)
	gsp.Entity = t
	gsp.UID = fmt.Sprintf("%s-%s", t, rand)
	gsp.IsValid = true
}

func (gsp *GSPModel) SetInvalid() {
	gsp.IsValid = false
}
