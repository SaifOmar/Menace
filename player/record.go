package player

import (
	"strconv"
	"strings"
)

// import "fmt"

type Eval int

// const (
//
//	terrible Eval = iota
//	bad
//	good
//	excelent
//	notEvenClose
//
// )
var evalName = map[int]string{
	0: "Terrible",
	1: "Bad",
	2: "Good",
	3: "Excelent",
	4: "Not Even Close",
}

// func (e Eval) String() string {
// 	return evalName[e]
// }

type Record struct {
	RecordArr     []bool
	Evaluations   []string
	WinPercentage float64
}

func NewRecord(recordArr []bool) *Record {
	return &Record{
		RecordArr: recordArr,
		// Evaluations:   evals,
		WinPercentage: 0.00,
	}
}

func (r *Record) EvaluatePlayerPerformance(n int) string {
	eval := evalName[n]
	r.Evaluations = append(r.Evaluations, eval)
	return eval
}

func (r *Record) GetNormalizedWp() string {
	wp := r.WinPercentage
	s := strconv.FormatFloat(wp, 'f', 1, 64)
	b := strings.Split(s, ".")
	s = "%" + b[0]
	return s
}

func (r *Record) CalculateWinPercentage() int {
	winCount := 0
	nMatches := len(r.RecordArr)
	for _, record := range r.RecordArr {
		if record {
			winCount += 1
		}
	}
	r.WinPercentage = float64((winCount/nMatches)*100 + int(r.WinPercentage))
	return int(r.WinPercentage)
}
