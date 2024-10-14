package main

import "fmt"

func main() {
	myMap1 := make(map[string]string)
	myMap1["name"] = "zhangsan"
	myMap1["age"] = "18"
	fmt.Println(myMap1)
	myMap1["sex"] = "男"

	delete(myMap1, "age")

	myMap1["sex"] = "女"
	fmt.Println(myMap1)

	for key, value := range myMap1 {
		fmt.Println(key, value)
	}

	fmt.Println(toSum([]int{1, 2, 3, 4, 5}, 6))
}
func toSum(nums []int, target int) (r1 int, r2 int) {
	if nums == nil {
		return
	}
	myMap := make(map[int]int)
	for index, value := range nums {
		if _, ok := myMap[target-value]; ok {
			r1 = index
			r2 = myMap[target-value]
			return
		} else {
			myMap[value] = index
		}
	}
	return
}
