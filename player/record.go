package player

type Eval int

const (
	terrible Eval = iota
	bad
	good
	excelent
	notEvenClose
)

var evalName = map[Eval]string{
	terrible:     "Terrible",
	bad:          "Bad",
	good:         "Good",
	excelent:     "Excelent",
	notEvenClose: "Not Even Close",
}

func (e Eval) String() string {
	return evalName[e]
}

type Record struct {
	RecordArr     []bool
	Evaluations   []Eval
	winPercentage float64
}

func NewRecord(recordArr []bool, evals []Eval) *Record {
	return &Record{
		RecordArr:     recordArr,
		Evaluations:   evals,
		winPercentage: 0.00,
	}
}

func (r *Record) CalculateWinPercentage() (float64, error) {
	winCount := 0
	nMatches := len(r.RecordArr)
	if nMatches == 0 {
		r.winPercentage = 0.00
		return r.winPercentage, nil
	}
	for _, record := range r.RecordArr {
		if record {
			winCount += 1
		}
	}
	r.winPercentage = float64((winCount/nMatches)*100 + int(r.winPercentage))
	return r.winPercentage, nil
}
