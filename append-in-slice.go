package main

import "fmt"

type cost struct {
	day   int
	value float64
}

// this function takes slice of cost and calculate the total price in each day of week is utilizes in slice

func getCostByDay(costs []cost) []float64 {
	costByDay := []float64{} // empty slice points to array of length 0 [...]float is for array not for slice
	for i := 0; i < len(costs); i++ {
		cost := costs[i]
		for cost.day >= len(costByDay) {
			costByDay = append(costByDay, 0.0) // here we are appending the 0.0 in slice until the day is reached for example if there only day 1 and day2 in slice the appending will be done till 2 values which means that the slice only two value will be returned
		}
		costByDay[cost.day] = cost.value
	}
	return costByDay
}

func main() {
	fmt.Println("appending in slice")
	costs := []cost{
		cost{day: 0,
			value: 0.0},
		cost{day: 2,
			value: 0.017},
		cost{day: 3,
			value: 0.12},
		cost{day: 4,
			value: 0.01},
		cost{day: 0,
			value: 0.01},
	}
	day_cost := getCostByDay(costs)
	days := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	for i := 0; i < len(day_cost); i++ {
		fmt.Println(days[i], " Cost :", day_cost[i])
	}
}

//slices
// var s1 []int        // nil slice (pointer = nil, len = 0, cap = 0)
// s2 := []int{}       // empty slice (points to array of length 0)
// fmt.Println(s1 == nil) // true
// fmt.Println(s2 == nil) // false

// s1 = append(s1, 1,2,3,4)
// s1 = append(s1, 4)
// s2:= []int{1,2,3,4}
// s1 = append(s1, s2...)
