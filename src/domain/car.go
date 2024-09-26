package domain

import (
	"fmt"
	"strconv"
	"strings"
)

type Car struct {
	ID      uint
	Year    uint
	Model   CarModel
	Edition string
	Engine  CarEngine
}

type CarModel struct {
	ID    uint
	Name  string
	Brand CarBrand
}

type CarBrand struct {
	ID   uint
	Name string
}

type CarEngine struct {
	ID            uint
	Code          string
	LiterCapacity float64
}

func (c Car) GetName() (finalString string) {
	stringComposition := []string{
		c.Model.Brand.Name, c.Model.Name, c.Edition, strconv.FormatFloat(c.Engine.LiterCapacity, 'f', -1, 64),
	}

	for _, e := range stringComposition {
		if e != "" {
			finalString = fmt.Sprintf("%s %s", finalString, e)
		}
	}

	return strings.TrimSpace(finalString)
}
