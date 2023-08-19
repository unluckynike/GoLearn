package main

import "fmt"

/*
map
*/
func main() {
	/*
		创建map
				map 是在 Go 中将值（value）与键（key）关联的内置类型。通过相应的键可以获取到值。
				通过向 make 函数传入键和值的类型，可以创建 map。make(map[type of key]type of value) 是创建 map 的语法。
				personSalary := make(map[string]int)
			// 创建一个空的 Map
			m := make(map[string]int)

			// 创建一个初始容量为 10 的 Map
			m := make(map[string]int, 10)
	*/

	//personSalary := make(map[string]int) //创建了一个名为 personSalary 的 map，其中键是 string 类型，而值是 int 类型。	map 的零值是 nil。
	var personSalary map[string]int
	if personSalary == nil {
		fmt.Println("map is nil. Going to make one.")
		personSalary = make(map[string]int) //添加元素到 nil map 中，会触发运行时 panic。因此 map 必须使用 make 函数初始化。
	}

	/*
		添加元素
	*/
	personSalary["park"] = 100
	personSalary["alice"] = 200
	personSalary["bob"] = 300
	fmt.Println(personSalary)

	/*
		声明时初始化
	*/
	m := map[string]int{ //在声明的同时添加两个元素
		"jack": 100,
		"eric": 900,
	}
	m["juice"] = 899
	fmt.Println(m)

	/*
				获取元素
			获取 map 元素的语法是 map[key]
		// 遍历 Map
		for k, v := range m {
		    fmt.Printf("key=%s, value=%d\n", k, v)
		}
	*/

	fmt.Println(personSalary["park"])
	fmt.Println(personSalary["kk"]) //如果获取一个不存在的元素，map 会返回该元素类型的零值

	for k, v := range personSalary {
		fmt.Println(k, v)
	}

	/*
			删除元素
		删除 map 中 key 的语法是 delete(map, key)。这个函数没有返回值。
	*/
	fmt.Println("未删除")
	for k, v := range m {
		fmt.Println(k, v)
	}
	delete(m, "eric")
	fmt.Println("删除后")
	for k, v := range m {
		fmt.Println(k, v)
	}

	/*
			获取长度
		获取 map 的长度使用 len 函数。
	*/
	fmt.Println(len(m), len(personSalary))

	/*
		引用类型
		 slices 类似，map 也是引用类型。当 map 被赋值为一个新变量的时候，它们指向同一个内部数据结构。因此，改变其中一个变量，就会影响到另一变量。
	*/
	fmt.Println(personSalary)
	newPersonSalary := personSalary
	newPersonSalary["alice"] = 999
	fmt.Println(personSalary) //被改掉了

	/*
			相等性
		map 之间不能使用 == 操作符判断，== 只能用来检查 map 是否为 nil。
	*/
	map1 := map[string]int{
		"one": 1,
		"two": 2,
	}

	map2 := map1

	fmt.Println(map2)
	//if map1 == map2 { 抛出编译错误 invalid operation: map1 == map2 (map can only be compared to nil)。
	//}

	var siteMap map[string]string /*创建集合 */
	siteMap = make(map[string]string)

	/* map 插入 key - value 对,各个国家对应的首都 */
	siteMap["Google"] = "谷歌"
	siteMap["Runoob"] = "菜鸟教程"
	siteMap["Baidu"] = "百度"
	siteMap["Wiki"] = "维基百科"

	/*使用键输出地图值 */
	for site := range siteMap {
		fmt.Println(site, "首都是", siteMap[site])
	}

	/*查看元素在集合中是否存在 */
	name, ok := siteMap["Facebook"] /*如果确定是真实的,则存在,否则不存在 */
	/*fmt.Println(capital) */
	/*fmt.Println(ok) */
	if ok {
		fmt.Println("Facebook 的 站点是", name)
	} else {
		fmt.Println("Facebook 站点不存在")
	}
}
