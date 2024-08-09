package main

import( "fmt"
       "sort")
// func UpdeteMap(m map[string]int, key string, newvalue int){
// 	if _, ok :=m[key]; ok{
// 		m[key] = newvalue
// 	}else {
// 		fmt.Println("Ключь не найден")
// 	}
// }
// func del(m map[string]int, key string) {
//     delete(m, key)}

	func sortedKeys(m map[string]int) []string {
		keys := make([]string, 0, len(m))
		for key := range m {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		return keys
	}
func main() {
	//1
	Per:=map[string]int{"Anton":18,"Stepan":21}
	fmt.Printf("Type : %T Value : %#v\n",Per,Per)
	fmt.Printf("Len : %d\n\n",len(Per))
	//2
	vtr:=map[int]string{
		15:"1-floor",
		25:"2-floor",
		35:"3-floor",
		45:"4-Floor",
		55:"5-Floor",
	}
_, exists:=vtr[45]
fmt.Println(exists)
	//3
// fmt.Println(UpdeteMap())
//4
// myMap := map[string]int{
// 	"one": 1,
// 	"two": 2,
// 	"three": 3,
// }

// fmt.Println("До удаления:", myMap)

// del(myMap, "two")

// fmt.Println("После удаления:", myMap)

myMap := map[string]int{
	"banana": 3,
	"apple":  5,
	"cherry": 2,
}
sorted := sortedKeys(myMap)
fmt.Println("Отсортированные ключи:", sorted)
}
