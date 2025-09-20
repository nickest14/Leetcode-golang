// https://leetcode.com/problems/design-spreadsheet/
// Level: Medium

package leetcode

import (
	"strconv"
	"strings"
)

type Spreadsheet struct {
	cellValues map[string]int
}

func Constructor(rows int) Spreadsheet {
	return Spreadsheet{cellValues: make(map[string]int)}
}

func (this *Spreadsheet) SetCell(cell string, value int) {
	this.cellValues[cell] = value
}

func (this *Spreadsheet) ResetCell(cell string) {
	this.cellValues[cell] = 0
}

func (this *Spreadsheet) GetValue(formula string) int {
	formula = formula[1:]
	terms := strings.Split(formula, "+")
	sum := 0
	for _, term := range terms {
		if numericValue, err := strconv.Atoi(term); err == nil {
			sum += numericValue
		} else {
			sum += this.cellValues[term]
		}
	}
	return sum
}
