package lib

import (
	"errors"
	"fmt"
	"log"
	"math"
)

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}

func Mu(a, b float32, x []float32) ([]float32, error) {
	mu := make([]float32, len(x))
	for i := 0; i < len(x); i++ {
		if x[i] >= b-a && x[i] < b+a {
			mu[i] = float32(0.5 * (1 + math.Cos(math.Pi*float64(x[i]-b)/float64(a))))
		} else {
			mu[i] = 0
		}
	}
	return mu, nil
}

func CreateX(a, b float32) ([]float32, error) {
	l := b - a
	r := b + a
	x := make([]float32, int(r-l))
	for i := 0; i < int(r-l); i++ {
		x[i] = l + float32(i)
	}
	return x, nil
}

func Find(m map[string][]float32) (float32, error) {
	var sum1, sum2 float32
	for k := 1; k < len(m["mu"]); k++ {
		sum1, sum2 = 0, 0
		for i := 0; i < k; i++ {
			sum1 += m["mu"][i]
		}
		for j := k + 1; j < len(m["mu"]); j++ {
			sum2 += m["mu"][j]
		}
		if float32(roundFloat(float64(sum1), 2)) == float32(roundFloat(float64(sum2), 2)) {
			return m["x"][k], nil
		}
	}
	return 0, errors.New("no solution found")
}

func Count(a, b float32) (float32, error) {
	x, err := CreateX(a, b)
	if err != nil {
		log.Printf("error creating x: %v", err)
		return 0, err
	}
	mu, err := Mu(a, b, x)
	if err != nil {
		log.Printf("error creating mu: %v", err)
		return 0, err
	}
	m := map[string][]float32{
		"mu": mu,
		"x":  x,
	}
	fmt.Printf("+----------+----------+\n")
	fmt.Printf("| %-10s | %-10s |\n", "μ(x)", "x")
	fmt.Printf("+----------+----------+\n")

	for i := 0; i < len(m["mu"]); i++ {
		fmt.Printf("| %-10.2f | %-10.2f |\n", m["mu"][i], m["x"][i])
	}

	fmt.Printf("+----------+----------+\n")
	result, err := Find(m)
	fmt.Printf("Result: %v\n", result)
	if err != nil {
		log.Printf("error finding result: %v", err)
		return 0, err
	}
	return result, nil
}
