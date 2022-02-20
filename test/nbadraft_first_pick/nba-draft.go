package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

/**
https://www.zhihu.com/question/52895544
*/
func main() {
	n := 14
	m := 4
	nn := 1
	mm := 1
	nm := 1
	for i := 1; i <= n; i++ {
		nn = nn * i
	}
	for i := 1; i <= m; i++ {
		mm = mm * i
	}
	for i := 1; i <= (n - m); i++ {
		nm = nm * i
	}

	fmt.Println(nn)
	fmt.Println(mm)
	fmt.Println(nm)
	fmt.Println(nn / mm / nm)

	all_result := make([]string, 1000)
	all_result_map := make(map[string]int)

	count := 0
	for i := 1; i <= 11; i++ {
		for j := i + 1; j <= 12; j++ {
			for k := j + 1; k <= 13; k++ {
				for l := k + 1; l <= 14; l++ {
					//fmt.Printf("%d -> %d -> %d -> %d\n", i, j, k, l)
					if i == 11 && j == 12 && k == 13 && l == 14 {
						continue
					}
					all_result[count] = fmt.Sprintf("%d-%d-%d-%d", i, j, k, l)
					count++
					all_result_map[fmt.Sprintf("%d-%d-%d-%d", i, j, k, l)] = count
				}
			}
		}
	}
	fmt.Println(count)
	fmt.Println(all_result[999])
	fmt.Println(all_result_map[all_result[999]])

	weight_arr := [14]int{250, 199, 156, 119, 88, 63, 43, 28, 17, 11, 8, 7, 6, 5}
	all_weight := 0
	source_map := make(map[int]int)

	for num, weight := range weight_arr {

		for i := 0; i < weight; i++ {
			source_map[i+1+all_weight] = num + 1
		}
		all_weight = all_weight + weight

	}
	fmt.Println(all_weight)
	fmt.Println(source_map[1000])
	fmt.Println(source_map[1])
	fmt.Println(source_map[250])
	fmt.Println(source_map[251])
	fmt.Println(source_map[995])
	fmt.Println(source_map[996])

	fmt.Println("模拟抽签---")

	all_ball := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

	chosen_ball := make([]int, 4)
	for i := 0; i < 4; i++ {
		r := rand.New(rand.NewSource(time.Now().Unix()))
		hit := r.Intn(14 - i)
		//fmt.Println("hit:", hit)
		chosen_ball[i] = all_ball[hit]

		// 指定删除位置
		index := hit
		// 输出删除位置之前和之后的元素
		//fmt.Println(all_ball[:index], all_ball[index+1:])
		// seq[index+1:]... 表示将后段的整个添加到前段中
		// 将删除前后的元素连接起来
		all_ball = append(all_ball[:index], all_ball[index+1:]...)

	}

	for _, result_num := range chosen_ball {
		//fmt.Printf("%d-> %d ", i, result_num)
		fmt.Printf("%d ", result_num)
	}
	fmt.Println()
	sort.Ints(chosen_ball)
	result := fmt.Sprintf("%d-%d-%d-%d", chosen_ball[0], chosen_ball[1], chosen_ball[2], chosen_ball[3])

	//fmt.Println(result)
	//result = "9-12-13-14"
	//fmt.Println(all_result_map[result])

	//fmt.Println(all_result_map["4-9-13-14"])
	fmt.Println(source_map[all_result_map[result]])

	//模拟100次，统计每个排名获得状元签的概率
	all_hit_result := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for xx := 0; xx < 100000; xx++ {
		all_ball := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14}

		chosen_ball := make([]int, 4)
		for i := 0; i < 4; i++ {

			hit := r.Intn(14 - i)
			//fmt.Println("hit:", hit)
			chosen_ball[i] = all_ball[hit]

			// 指定删除位置
			index := hit
			// 输出删除位置之前和之后的元素
			//fmt.Println(all_ball[:index], all_ball[index+1:])
			// seq[index+1:]... 表示将后段的整个添加到前段中
			// 将删除前后的元素连接起来
			all_ball = append(all_ball[:index], all_ball[index+1:]...)

		}

		sort.Ints(chosen_ball)
		result := fmt.Sprintf("%d-%d-%d-%d", chosen_ball[0], chosen_ball[1], chosen_ball[2], chosen_ball[3])

		//fmt.Println(result)
		//result = "9-12-13-14"
		//fmt.Println(all_result_map[result])

		//fmt.Println(source_map[all_result_map[result]])
		//fmt.Println(result)
		if result == "11-12-13-14" {
			xx--
			continue
		}

		//fmt.Println("---", all_result_map[result])
		//fmt.Println("====", source_map[all_result_map[result]])
		all_hit_result[source_map[all_result_map[result]]-1]++
	}
	for i, resut_hit := range all_hit_result {
		fmt.Printf("%d -> %d\n", i+1, resut_hit)
	}

	/**
		模拟10w次！
	1 -> 25088
	2 -> 19949
	3 -> 15772
	4 -> 11926
	5 -> 8730
	6 -> 6201
	7 -> 4276
	8 -> 2830
	9 -> 1641
	10 -> 1032
	11 -> 802
	12 -> 673
	13 -> 580
	14 -> 500
	*/

	/**
	模拟前三顺位
	*/

}
