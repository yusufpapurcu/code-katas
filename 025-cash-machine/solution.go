package cashmachine

var cashList = []Cash{
	{Name: "£50", Value: 50},
	{Name: "£20", Value: 20},
	{Name: "£10", Value: 10},
	{Name: "£5", Value: 5},
	{Name: "£2", Value: 2},
	{Name: "£1", Value: 1},
	{Name: "50p", Value: 0.50},
	{Name: "20p", Value: 0.20},
	{Name: "10p", Value: 0.10},
	{Name: "5p", Value: 0.05},
	{Name: "2p", Value: 0.02},
	{Name: "1p", Value: 0.01},
}

func BreakIntoChange(amount float64) map[string]int {
	res := make(map[string]int)
	for _, c := range cashList {
		res[c.Name] = int(amount / c.Value)
		amount -= float64(res[c.Name]) * c.Value
	}
	return res
}

type Cash struct {
	Value float64
	Name  string
}
