package entities

type Menu map[string]CoffeeList

type CoffeeList map[string]Coffee

type Coffee struct {
	Volume      []float64
	Prices      []int
	Composition []string
}
