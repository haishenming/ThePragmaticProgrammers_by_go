package main

import (
	Tips "Chapter1/Tips"
	"fmt"
	"os"
)

func T1() {
	fmt.Println("计算小费")

	var total float32    // 总金额可能是小数
	var tip int

	fmt.Println("请输入总金额：")
	_, err := fmt.Scanf("%f", &total)
	if err != nil {
		fmt.Println("请输入正确的数值类型")
		os.Exit(0)
	}

	fmt.Println("请输入小费比例：")
	fmt.Scanf("%d", &tip)

	tips := Tips.Tips(total, tip)
	totalTips := float32(total) - tips
	fmt.Printf("小费金额是：%.2f\n", tips)
	fmt.Printf("应付金额是：%.2f\n", totalTips)
}

func main() {
	T1()

}
