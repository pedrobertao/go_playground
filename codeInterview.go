package main

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var depth = 0.0
var depth2 = 0.0

type Queue []int

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}

func (q *Queue) Enqueue(i int) {
	*q = append(*q, i)
}

func (q *Queue) Dequeue() (int, bool) {
	if q.IsEmpty() {
		return 0, false
	} else {
		element := (*q)[0]
		*q = (*q)[1:]
		return element, true
	}
}
func jumpingOnClouds(c []int32) int32 {
	// Write your code here
	jumps := int32(0)
	max := len(c)

	i := 0
	for i < max {
		fmt.Println(i)
		if i+2 < max {
			if c[i+2] == 0 {
				jumps++
				i = i + 2
				continue
			}
		}
		if i+1 < max {
			if c[i+1] == 0 {
				jumps++
				i = i + 1
				continue
			}
		}
		break
	}
	fmt.Println(jumps)
	return jumps
}

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, num := range nums {
		diff := target - num
		if idx, found := m[diff]; found {
			return []int{i, idx}
		}
		m[num] = i
	}

	return nil
}

func isValid(s string) bool {

	if len(s) == 1 {
		return false
	}
	count := 0
	for i, v := range s {
		if v == '(' {
			count += 2
		}
		if v == ')' {
			if i-1 > 0 && (s[i-1] == '[' || s[i-1] == '{') {
				return false
			}
			count -= 2

		}

		if v == '[' {
			count += 3
		}
		if v == ']' {
			if i-1 > 0 && (s[i-1] == '{' || s[i-1] == '(') {
				return false
			}
			count -= 3
		}

		if v == '{' {
			count += 4
		}
		if v == '}' {
			if i-1 > 0 && (s[i-1] == '(' || s[i-1] == '[') {
				return false
			}
			count -= 4
		}
		if count < 0 {
			return false
		}
	}
	return count == 0
}

func sherlockAndAnagrams(s string) int32 {
	subArr := []string{}
	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			subArr = append(subArr, s[i:j+1])
		}
	}

	count := int32(0)
	for i := 0; i < len(subArr); i++ {
		for j := i + 1; j < len(subArr); j++ {
			if isAnagram2(subArr[i], subArr[j]) {
				count++
			}
		}

	}
	return count
}

func isAnagram2(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	countLetterS := make(map[rune]int)
	countLetterT := make(map[rune]int)
	for _, v := range s {
		if countLetterS[v] == 0 {
			countLetterS[v] = 1
		} else if countLetterS[v] > 0 {
			countLetterS[v]++
		}
	}
	for _, v := range t {
		if countLetterT[v] == 0 {
			countLetterT[v] = 1
		} else if countLetterT[v] > 0 {
			countLetterT[v]++
		}
	}

	for k, vS := range countLetterS {
		vT := countLetterT[k]
		if vT != vS {
			return false
		}
	}
	return true
}

func isAnagram(s string, t string) bool {
	dif := make(map[rune]int)

	for _, v := range s {
		if dif[v] == 0 {
			dif[v] = 1
		}
		if dif[v] > 0 {
			dif[v]++
		}
	}
	for _, v := range t {
		if dif[v] == 0 {
			dif[v] = 1
		}
		if dif[v] > 0 {
			dif[v]++
		}
	}
	for _, v := range dif {
		if v%2 != 0 {
			return false
		}
	}
	return true
}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + int(math.Max(float64(maxDepth(root.Right)), float64(maxDepth(root.Left))))
}

func isPalindrome(s string) bool {
	var nonAlphanumericRegex = regexp.MustCompile(`[^a-zA-Z0-9]+`)
	newWord := ""
	for i := len(s) - 1; i >= 0; i-- {
		r := s[i]
		newWord = newWord + string(r)
	}
	newWord = strings.ToLower(strings.ReplaceAll(newWord, " ", ""))
	originWord := strings.ToLower(strings.ReplaceAll(s, " ", ""))
	return nonAlphanumericRegex.ReplaceAllString(newWord, "") == nonAlphanumericRegex.ReplaceAllString(originWord, "")
}

func removeDuplicates2(nums []int) int {
	dif := make(map[int]int, len(nums))
	k := 0
	// temp := make([]int, 0)
	for i, v := range nums {
		if dif[v] == 2 {
			nums[i] = math.MaxInt
		}
		if dif[v] == 1 {
			// temp = append(temp, v)
			nums[i] = v
			dif[v] = 2
			k++
		}
		if dif[v] == 0 {
			// temp = append(temp, v)
			nums[i] = v
			dif[v] = 1
			k++
		}
	}
	sort.Ints(nums)

	return k
}

func rotLeft(a []int32, d int32) []int32 {

	lenArr := len(a)

	ref := make([]int32, len(a))
	copy(ref, a)
	for i := int32(0); i < d; i++ {
		newArr := []int32{}
		newArr = append(newArr, ref[1:lenArr]...)
		newArr = append(newArr, ref[0])
		fmt.Println("Res:", newArr)
		copy(ref, newArr)
	}

	// 1
	// 1,2,3,4,5
	// 2,3,4,5,1
	// fmt.Println(".1", newArr)
	// fmt.Println(".2", newArr)

	return ref

}

func minimumBribes(q []int32) {
	// Write your code here

	// bribes := make(map[int]int)
	bribes := 0
	for i := len(q) - 1; i >= 0; i-- {
		if q[i] == int32(i+1) {
			continue
		} else {
			if i-1 >= 0 {
				fmt.Println("Testing waters...", q[i], int32(i-1), q[i]-int32(i-1))
				if int32(i-1)-q[i] < 0 {
					bribes++
				}
			}

			// fmt.Println("How many this guy bribe", q[i], "[", i, "]", " was : ", len(q)-i)
		}
	}
	fmt.Println("Too chaotic", bribes)
}

func minimumDistances(a []int32) int32 {
	type Dist struct {
		FirstIndex int32
		LastIndex  int32
		Dist       int32
	}
	if len(a) == 0 || len(a) == 1 {
		return -1
	}

	distMap := make(map[int32]Dist)
	min := int32(-1)
	for cI, val := range a {
		trueIndex := int32(cI) + 1
		if distMap[val].FirstIndex == 0 {
			distMap[val] = Dist{
				FirstIndex: trueIndex,
				LastIndex:  trueIndex,
				Dist:       0,
			}
		} else {
			if (trueIndex - 1) >= distMap[val].LastIndex {
				cp := distMap[val]
				cp.LastIndex = trueIndex
				cp.Dist = cp.LastIndex - cp.FirstIndex
				distMap[val] = cp

				if min == -1 || cp.Dist <= min {
					min = cp.Dist
				}
			}
		}
	}
	return min
}

func isSubsequence(s string, t string) bool {
	i := 0
	j := 0
	for i < len(s) && j < len(t) {
		if string(t[j]) == string(s[i]) {
			i++
		}

		j++
	}
	return i == len(s)
}

func isHappy(n int) bool {
	if n <= 0 {
		return false
	}

	sn := strconv.Itoa(n)

	isRepeated := make(map[int]bool)

	sum := 0
	nextString := sn
	for {
		sum = 0
		for _, v := range nextString {
			nv, _ := strconv.Atoi(string(v))
			sum += nv * nv
		}
		nextSum := fmt.Sprintf("%d", sum)
		fmt.Println("nextString", nextString, "sum:", nextSum)
		if isRepeated[sum] || (n > 1 && sum == n) {
			return false
		}

		isRepeated[sum] = true
		if sum == 1 {
			return true
		}

		if nextSum == nextString {
			return false
		}
		nextString = nextSum
	}

}

func twoStrings(s1 string, s2 string) string {
	// Write your code here
	dic1 := make(map[string]int)
	for _, v := range s1 {
		dic1[string(v)]++
	}

	for _, v := range s2 {
		if dic1[string(v)] > 0 {
			return "YES"
		}
	}
	return "NO"
}

func checkMagazine(magazine []string, note []string) {
	dic := make(map[string]int)
	for _, n := range magazine {
		dic[n]++
	}
	for _, n := range note {
		if val, ok := dic[n]; !ok || val == 0 {
			fmt.Println("No")
			return
		} else {
			dic[n]--
		}
	}
	fmt.Println("Yes")
}
func removeDuplicates(nums []int) int {
	dif := make(map[int]bool)
	k := 0
	for i, v := range nums {
		if dif[v] {
			nums[i] = -1
			continue
		} else {
			dif[v] = true
			nums[k] = v
			k++
		}
	}
	return k
}

func removeElement(nums []int, val int) int {
	k := 0
	for i, v := range nums {
		if v != val {
			k++
			continue
		} else {
			nums[i] = math.MaxInt
		}
	}
	sort.Ints(nums)
	return k
}

func fibRecursive(position uint) uint {
	if position <= 2 {
		return 1
	}
	return fibRecursive(position-1) + fibRecursive(position-2)
}

func fibIterative(position uint) uint {
	slc := make([]uint, position)
	slc[0] = 1
	slc[1] = 1

	if position <= 2 {
		return 1
	}

	var result, i uint
	for i = 2; i < position; i++ {
		result = slc[i-1] + slc[i-2]
		slc[i] = result
	}

	return result
}
func maximumToys(prices []int32, k int32) int32 {
	sort.Slice(prices, func(i, j int) bool {
		return prices[j] > prices[i]
	})
	count := k
	countMaxNum := int32(0)
	for i := 0; i < len(prices); i++ {
		count -= prices[i]
		if count >= 0 {
			countMaxNum++
		} else if count < 0 {
			break
		}
	}
	return countMaxNum
}

func bubbleSort(arr []int32) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				prev := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = prev
			}
		}
	}
}

func countSwaps(a []int32) {
	count := 0
	// fmt.Println(a)
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a)-1; j++ {
			if a[j] > a[j+1] {
				prev := a[j]
				a[j] = a[j+1]
				a[j+1] = prev
				// fmt.Println(a)
			}
		}

	}
	fmt.Println(count)

}
