package main

import (
	"fmt"
	"strings"
)

/*
你有50枚金币，需要分配给以下几个人：Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
分配规则如下：
a. 名字中每包含1个'e'或'E'分1枚金币
b. 名字中每包含1个'i'或'I'分2枚金币
c. 名字中每包含1个'o'或'O'分3枚金币
d: 名字中每包含1个'u'或'U'分4枚金币
写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
程序结构如下，请实现 ‘dispatchCoin’ 函数
*/
var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func dispatchCoin() (int, map[string]int) {
	m1 := make(map[string][]string, 10)
	m2 := make(map[string]int, 10)

	for _, v := range users {
		m1[v] = strings.Split(v, "")
	}
	// fmt.Println(m1)

	for k, v := range m1 {
		for j := 0; j < len(v); j++ {
			switch v[j] {
			case "e":
				coins--
				m2[k]++
			case "E":
				coins--
				m2[k]++
			case "i":
				coins = coins - 2
				m2[k] = m2[k] + 2
			case "I":
				coins = coins - 2
				m2[k] = m2[k] + 2
				m2[k]++
			case "o":
				coins = coins - 3
				m2[k] = m2[k] + 3
			case "O":
				coins = coins - 3
				m2[k] = m2[k] + 3
				m2[k]++
			case "u":
				coins = coins - 4
				m2[k] = m2[k] + 4
			case "U":
				coins = coins - 4
				m2[k] = m2[k] + 4
			default:

			}
		}
	}
	fmt.Println(m2)
	fmt.Println(coins)
	return coins, m2

}
func main() {
	left, map1 := dispatchCoin()
	// dispatchCoin()
	fmt.Println("剩下：", left)
	fmt.Println("各自金币个数：", map1)
}
