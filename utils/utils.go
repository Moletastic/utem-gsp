package utils

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"strconv"
)

func GenerateID(prefix string) string {
	rand := strconv.Itoa(rand.Intn(9999-1000) + 1000)
	return fmt.Sprintf("%s-%s", prefix, rand)
}

func Pretty(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "  ")
	return string(s)
}

func GetLocalAddress(port uint) string {
	return fmt.Sprintf(":%d", port)
}
