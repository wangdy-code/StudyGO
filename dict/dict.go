package dict

import "fmt"

/*对一个nil map 做一些操作 遍历打印  不会panic*/
func NilDictDoSomething() {
	var m map[int]int
	for key, value := range m {
		fmt.Printf("map's key = %d,value = %d", key, value)
	}
}
