package main

import (
	"fmt"
	"time"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	nums = []int{-4,-2,1,-5,-4,-4,4,-2,0,4,0,-2,3,1,-5,0}

	t1 := time.Now().Unix()
	ret := threeSum(nums)
	t2 := time.Now().Unix()
	fmt.Printf("%+v\n",ret)
	fmt.Printf("%+v",t2 - t1)

}
func threeSum(nums []int) [][]int {
	var tmpInt  = make(map[int][]int)
	var ret [][]int
	var retMap = make(map[int]map[int]bool)
	for idx , it := range nums {
		tmpInt[it] = append(tmpInt[it],idx)
	}
	var index = 0
	for i := 0 ; i < len(nums) ; i++ {
		for j := i ; j < len(nums) ; j++ {
			if i == j {
				continue
			}
			tmp := nums[i] + nums[j]
			tmp = -tmp
			if idxList , ok := tmpInt[tmp] ; ok {
				for _ , idx := range idxList {
					if idx != i && idx != j {
						tmpOne := []int{ nums[i] , nums[j]}
						tmpOne = append(tmpOne , nums[idx])
						isFind := false
						isFind = enquel(retMap , tmpOne)
						if !isFind {
							tmpMap := make(map[int]bool)
							ret = append(ret , tmpOne)
							tmpMap[nums[i]] = true
							tmpMap[nums[j]] = true
							tmpMap[nums[idx]] = true
							retMap[index] = tmpMap
							index++
						}
					}
				}
			}
		}
	}
	return ret
}
func enquel(a map[int]map[int]bool , b []int) bool {
	for _ , ta := range a {
		_ , ok1 := ta[b[0]]
		_ , ok2 := ta[b[1]]
		_ , ok3 := ta[b[2]]


		if ok1 && ok2 && ok3 {
			if b[0] == 0 && b[1] == 0 && b[2] == 0 && len(ta) != 1 {
				continue
			}
			return  true
		}
	}
	return false
}
