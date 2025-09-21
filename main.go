package main

import (
	"fmt"
)

func main() {

	//3/8

	// var numbers = [5]int{1, 2, 4, 4, 5}
	// pickOneNumber(numbers[:])

	//fmt.Println(palindromeNumber(123))

	// str := "[]{}("
	// fmt.Println(isValid2(str))

	// var strs = []string{"flower", "flow", "flight"}.
	// longestCommonPrefix(strs)

}

// 1
func pickOneNumber(target []int) {

	//定义一个map，存对比的结果，每个元素出现了几次
	result := make(map[int]int)

	//对比
	for i := 0; i < len(target); i++ {
		// fmt.Printf("%d star\n", i)
		count := 1
		for j := 0; j < len(target); j++ {
			//自己和自己不比
			if i == j {
				continue
			}
			//已经比过的不比
			if _, ok := result[target[i]]; ok {
				// fmt.Println("key named in the map")
				continue
			} else {
				// fmt.Println("No key named in the map")
			}
			//
			if target[i] == target[j] {
				count++
			}
		}

		if _, ok := result[target[i]]; !ok {
			result[target[i]] = count
		}

	}
	for k, v := range result {
		if v == 1 {
			fmt.Println(k)
		}
	}

}

// 2
func palindromeNumber(num int) bool {
	var target = num
	var y = 0
	for target > 0 {
		y = y*10 + target%10
		target = target / 10
		fmt.Println(y)
	}

	if y == num {
		return true
	}
	return false

}

// 3
func isValid(targetStr string) bool {

	//构建一个map
	leftRightMap := map[string]string{")": "(", "}": "{", "]": "["}

	var left_slice = make([]string, 0) //构建左结构，为入栈操作

	var arr_left [3]string = [3]string{"(", "[", "{"}
	var arr_right [3]string = [3]string{")", "]", "}"}

	var pushCount, popCount int = 0, 0

	for index, char := range targetStr {

		strCurr := string(char) //转成字符串
		fmt.Printf(">>>>>>>>current position:%d: %s\n", index, strCurr)

		for _index, string := range arr_left {
			if string == strCurr {
				fmt.Printf("yes,%d,%s\n", _index, strCurr)
				//压栈操作
				left_slice = append(left_slice, strCurr)
				fmt.Printf("push operation:%v\n", left_slice)
				pushCount++
				break
			} else {
				fmt.Println("no")
			}
		}

		//判断下一个是不是右，如果是右，就要匹配是否成对，如果成对，就出栈，否则判定fasle
		for _index, rightString := range arr_right {
			if index+1 >= len(targetStr) {
				break
			}

			if string(targetStr[index+1]) == rightString {
				fmt.Println("right")
				fmt.Println(_index)
				//去left_slice里匹配
				for v1, v2 := range left_slice {
					fmt.Println(v1, v2)
					if v2 == leftRightMap[rightString] {
						fmt.Println("match2")
						//出栈
						if len(left_slice) > 0 {
							left_slice = left_slice[:len(left_slice)-1]
							fmt.Printf("pop operation:%v\n", left_slice)
							popCount++
						}

					}

				}

				//要判断是否成对，用rightString作为key取value是否和strCurr相等，相等则说明是成对的，否则不成对
				if leftRightMap[rightString] == strCurr {
					fmt.Println("match1")
					//出栈
					if len(left_slice) > 0 {
						left_slice = left_slice[:len(left_slice)-1]
						fmt.Printf("pop operation:%v\n", left_slice)
						popCount++
					}
				}

			} else {
			}
		}

		//判断最后一位的情况
		if index+1 == len(targetStr) && len(left_slice) > 0 {
			//1.最后一位是左，直接判断false
			for _index, string := range arr_left {
				if string == strCurr {
					fmt.Printf("last is left,%d,%s\n", _index, strCurr)
					break
				} else {
				}

			}

			//2，最后一位是右，要在left_slice取匹配，如果匹配到就出栈
			for _index, rightString := range arr_right {
				if rightString == strCurr {
					fmt.Printf("last is right,%d,%s\n", _index, strCurr)
					fmt.Println(left_slice)

					for _index, str := range left_slice {

						if leftRightMap[rightString] == str {
							fmt.Println("last match %d", _index)
							//出栈
							if len(left_slice) > 0 {
								left_slice = left_slice[:len(left_slice)-1]
								fmt.Printf("pop operation:%v\n", left_slice)
							}
						}

					}

				}

			}

		}

		fmt.Printf("<<<<<<<<%d end\n", index)
		fmt.Printf("%d:%v\n", len(left_slice), left_slice)
		fmt.Printf("<<<<<<<<\n")

	}
	fmt.Printf("~~~~~all end\n")
	fmt.Printf("%d:%v\n", len(left_slice), left_slice)
	fmt.Printf("~~~~~push:%d,pop:%d\n", pushCount, popCount)

	if len(left_slice) == 0 {
		return true
	}
	return false
}

func isValid2(target string) bool {

	leftRightMap := map[string]string{")": "(", "}": "{", "]": "["}

	var left_slice = make([]string, 0)
	var pushCount, popCount int = 0, 0

	for index, char := range target {

		strCurr := string(char) //转成字符串
		fmt.Printf(">>>>>>>>current position:%d: %s\n", index, strCurr)

		//1,先判断左右
		//字典的所有key是否包含strCurr，是为右。否为左
		//走完三次循环为左，否为右
		var count int = 0
		for v1, _ := range leftRightMap {
			fmt.Println(v1)
			if v1 == strCurr {
				// fmt.Println("yes match")
				break

			} else {

			}
			count++
		}
		if count == 3 {
			fmt.Printf("this is left\n")
			left_slice = append(left_slice, strCurr)
			pushCount++

		} else {
			fmt.Printf("this is right\n")

			//left_slice最后一位和当前匹配就pop
			//1.先取left_slice最后一位
			lastElement := left_slice[len(left_slice)-1]
			fmt.Printf("最后一个元素: %s\n", lastElement)
			// fmt.Printf(leftRightMap[strCurr])

			if lastElement == leftRightMap[strCurr] {
				//出栈
				if len(left_slice) > 0 {
					left_slice = left_slice[:len(left_slice)-1]
					popCount++
				}
			}
		}
	}

	fmt.Printf("~~~~~all end\n")
	fmt.Printf("%d:%v\n", len(left_slice), left_slice)
	fmt.Printf("~~~~~push:%d,pop:%d\n", pushCount, popCount)
	if len(left_slice) == 0 {
		return true
	}

	return false
}

// 4
func longestCommonPrefix(strs []string) string {
	return "ok"
}
