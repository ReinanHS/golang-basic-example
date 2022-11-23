// Package util é responsável com conter vários métodos, usados como helpers
package util

import (
	"fmt"
	"time"
)

//FormatDataPtBr é um método responsável por converter uma time.Time em uma string com uma data no formato brasileiro
func FormatDataPtBr(t time.Time) string {
	return fmt.Sprintf("%02d/%02d/%d %02d:%02d", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute())
}
