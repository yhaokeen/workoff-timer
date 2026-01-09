# Festival - 节日查询库

从 tyme4go 提取的精简节日查询库，专注于**获取最近节日**功能。

## 功能特点

- ✅ 查找最近的节日（公历节日、农历节日、节气）
- ✅ 精确的天文算法（基于寿星天文历）
- ✅ 支持公历农历互转
- ✅ 轻量级设计，核心代码仅 1135 行

## 文件结构

```
festival/
├── ShouXingUtil.go  (555行) - 天文计算核心（CalcQi节气, CalcShuo朔日）
├── lunar.go         (264行) - 农历系统（年月日 + 农历节日）
└── solar.go         (316行) - 公历系统（日期 + 公历节日 + 节气 + 统一接口）
```

## 使用示例

### 基本用法

```go
package main

import (
    "fmt"
    "github.com/6tail/tyme4go/festival"
)

func main() {
    // 获取今天之后 30 天内最近的节日
    f := festival.Today().GetNearestFestival(30)
    if f != nil {
        fmt.Println("节日名称:", f.Name)
        fmt.Println("节日类型:", f.Type)
        fmt.Println("公历日期:", f.SolarDay)
    }
}
```

### 查找指定日期的节日

```go
// 创建日期
day, _ := festival.NewSolarDay(2025, 1, 1)

// 查找 10 天内的节日
f := day.GetNearestFestival(10)
if f != nil {
    fmt.Printf("%s 最近的节日是: %s\n", day, f.Name)
}
```

### 节日类型

```go
switch f.Type {
case festival.FestivalTypeSolar:
    fmt.Println("公历节日:", f.Name)  // 元旦、国庆等
case festival.FestivalTypeLunar:
    fmt.Println("农历节日:", f.Name)  // 春节、中秋等
case festival.FestivalTypeSolarTerm:
    fmt.Println("节气:", f.Name)      // 冬至、立春等
}
```

## 支持的节日

### 公历节日（10个）
元旦、三八妇女节、植树节、五一劳动节、五四青年节、六一儿童节、建党节、八一建军节、教师节、国庆节

### 农历节日（13个）
春节、元宵节、龙头节、上巳节、清明节、端午节、七夕节、中元节、中秋节、重阳节、冬至节、腊八节、除夕

### 二十四节气
冬至、小寒、大寒、立春、雨水、惊蛰、春分、清明、谷雨、立夏、小满、芒种、夏至、小暑、大暑、立秋、处暑、白露、秋分、寒露、霜降、立冬、小雪、大雪

## 核心API

### Festival (节日)
```go
type Festival struct {
    Type     FestivalTypeEnum  // 节日类型
    Name     string            // 节日名称
    SolarDay SolarDay          // 公历日期
}
```

### SolarDay (公历日)
```go
// 获取最近节日（向后查找）
func (o SolarDay) GetNearestFestival(maxDays int) *Festival

// 获取农历日期
func (o SolarDay) GetLunarDay() LunarDay

// 日期推移
func (o SolarDay) Next(n int) SolarDay
```