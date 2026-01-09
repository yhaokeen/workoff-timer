package festival_test

import (
	"fmt"
	"testing"

	"workoff-timer/internal/festival"
)

// ExampleGetNearestFestival 获取最近节日示例
func ExampleGetNearestFestival() {
	// 从今天开始查找30天内最近的节日
	f := festival.Today().GetNearestFestival(30)
	if f != nil {
		fmt.Printf("最近的节日: %s (%s)\n", f.Name, f.Type)
		fmt.Printf("日期: %s\n", f.SolarDay)
	}
}

// TestFestival 节日查找测试
func TestFestival(t *testing.T) {
	// 测试公历节日
	day, _ := festival.NewSolarDay(2025, 1, 1)
	f := day.GetNearestFestival(30)
	if f == nil || f.Name != "元旦" {
		t.Errorf("期望找到元旦节日")
	}

	// 测试节气
	day2, _ := festival.NewSolarDay(2024, 12, 20)
	f2 := day2.GetNearestFestival(5)
	if f2 == nil {
		t.Errorf("期望找到冬至节气")
	}
	fmt.Printf("找到节气: %s, 日期: %s\n", f2.Name, f2.SolarDay)
}
