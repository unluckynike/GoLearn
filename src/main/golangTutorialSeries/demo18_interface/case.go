package main

import "fmt"

// 根据公司员工的个人薪资，计算公司的总支出
type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct { //结构体 长期工
	empId    int
	basicpay int
	pf       int
}

type Contract struct { //结构体 合同工
	empId    int
	basicpay int
}

// salary of permanent employee is sum of basic pay and pf
func (p Permanent) CalculateSalary() int { //长期工计算工资的方法
	return p.basicpay + p.pf
}

// salary of contract employee is the basic pay alone
func (c Contract) CalculateSalary() int { //合同工计算工资的方法
	return c.basicpay
}

func totalExpense(s []SalaryCalculator) { //方法接收一个 SalaryCalculator 接口的切片（[]SalaryCalculator）作为参数。
	expense := 0
	for _, v := range s {
		expense = expense + v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {
	pemp1 := Permanent{1, 5000, 20}
	pemp2 := Permanent{2, 6000, 30}
	cemp1 := Contract{3, 3000}
	employees := []SalaryCalculator{pemp1, pemp2, cemp1} //向 totalExpense 方法传递了一个包含 Permanent 和 Contact 类型的切片。通过调用不同类型对应的 CalculateSalary 方法，totalExpense 可以计算得到支出。
	totalExpense(employees)

}
