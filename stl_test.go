package stl

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"testing"
)

func TestDecompose(t *testing.T) {
	trend, seasonal, remainder, err := Decompose([]float64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}, 5)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(trend, seasonal, remainder)
}

func BenchmarkDecompose(b *testing.B) {
	var (
		data []float64
		//trend     []float64
		//seasonal  []float64
		//remainder []float64
		err error
	)
	f, _ := (os.Open("testdata/co2.csv"))
	r := csv.NewReader(f)
	r.Read()
	for rec, err := r.Read(); err == nil; rec, err = r.Read() {
		if co2, err := strconv.ParseFloat(rec[0], 64); err == nil {
			data = append(data, co2)
		}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _, _, err = Decompose(data, 12)
		if err != nil {
			b.Fatal(err)
		}
	}
	//fmt.Println("trend", trend)
	//fmt.Println("seasonal", seasonal)
	//fmt.Println("remainder", remainder)
}
