package game_of_three

func Solution(num int) []int {
	var res = []int{}
	isNumNegative := num < 0
	if isNumNegative {
		num = 0 - num
	}
	for {
		res = append(res, num)
		if num == 1 {
			break
		}

		switch num % 3 {
		case 1:
			num--
		case 2:
			num++
		case 0:
			num /= 3
		}
	}

	if isNumNegative {
		for i := range res {
			res[i] *= -1
		}
	}
	return res
}
