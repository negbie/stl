package stl

import (
	"encoding/csv"
	"fmt"
	"math"
	"os"
	"strconv"
	"testing"
)

var data []float64
var dataSmall = []float64{1, 2, 3, 4, 5, 6, 7, math.NaN(), 9, 10, 11, 12, 13, 14, 15}

func init() {
	f, _ := (os.Open("testdata/co2.csv"))
	r := csv.NewReader(f)
	r.Read()
	for rec, err := r.Read(); err == nil; rec, err = r.Read() {
		if co2, err := strconv.ParseFloat(rec[0], 64); err == nil {
			data = append(data, co2)
		}
	}
}

func TestDecompose(t *testing.T) {
	trend, seasonal, remainder, err := Decompose(
		dataSmall,
		5,
		SDegree(0),
		TDegree(0),
		LDegree(0),
		SWindow(200),
		TWindow(200),
		LWindow(200),
		SJump(0),
		TJump(0),
		LJump(0),
		OuterLoop(10),
		InnerLoop(10),
		CritFreq(0.01),
	)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(trend, seasonal, remainder)
}

func BenchmarkDecompose(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _, err := Decompose(data, 12)
		if err != nil {
			b.Fatal(err)
		}
	}
}
