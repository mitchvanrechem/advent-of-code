package main

func getMonkeys() []Monkey {
	return []Monkey{
		{
			items:              []int{76, 88, 96, 97, 58, 61, 67},
			operation:          func(oldWorryLevel int) int { return oldWorryLevel * 19 },
			divisionTestNumber: 3,
			positiveTestMonkey: 2,
			negativeTestMonkey: 3,
		},
		{
			items:              []int{93, 71, 79, 83, 69, 70, 94, 98},
			operation:          func(oldWorryLevel int) int { return oldWorryLevel + 8 },
			divisionTestNumber: 11,
			positiveTestMonkey: 5,
			negativeTestMonkey: 6,
		},
		{
			items:              []int{50, 74, 67, 92, 61, 76},
			operation:          func(oldWorryLevel int) int { return oldWorryLevel * 13 },
			divisionTestNumber: 19,
			positiveTestMonkey: 3,
			negativeTestMonkey: 1,
		},
		{
			items:              []int{76, 92},
			operation:          func(oldWorryLevel int) int { return oldWorryLevel + 6 },
			divisionTestNumber: 5,
			positiveTestMonkey: 1,
			negativeTestMonkey: 6,
		},
		{
			items:              []int{74, 94, 55, 87, 62},
			operation:          func(oldWorryLevel int) int { return oldWorryLevel + 5 },
			divisionTestNumber: 2,
			positiveTestMonkey: 2,
			negativeTestMonkey: 0,
		},
		{
			items:              []int{59, 62, 53, 62},
			operation:          func(oldWorryLevel int) int { return oldWorryLevel * oldWorryLevel },
			divisionTestNumber: 7,
			positiveTestMonkey: 4,
			negativeTestMonkey: 7,
		},
		{
			items:              []int{62},
			operation:          func(oldWorryLevel int) int { return oldWorryLevel + 2 },
			divisionTestNumber: 17,
			positiveTestMonkey: 5,
			negativeTestMonkey: 7,
		},
		{
			items:              []int{85, 54, 53},
			operation:          func(oldWorryLevel int) int { return oldWorryLevel + 3 },
			divisionTestNumber: 13,
			positiveTestMonkey: 4,
			negativeTestMonkey: 0,
		},
	}
}

// Data from the given example
func getExampleMonkeys() []Monkey {
	return []Monkey{
		{[]int{79, 98}, func(oldWorryLevel int) int { return oldWorryLevel * 19 }, 23, 2, 3},
		{[]int{54, 65, 75, 74}, func(oldWorryLevel int) int { return oldWorryLevel + 6 }, 19, 2, 0},
		{[]int{79, 60, 97}, func(oldWorryLevel int) int { return oldWorryLevel * oldWorryLevel }, 13, 1, 3},
		{[]int{74}, func(oldWorryLevel int) int { return oldWorryLevel + 3 }, 17, 0, 1},
	}
}
