package festival

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"time"
)

// ============ 儒略日 ============

// J2000 2000年儒略日数(2000-1-1 12:00:00 UTC)
const J2000 = 2451545

// JulianDay 儒略日
type JulianDay struct {
	day float64
}

// NewJulianDay 从儒略日创建
func NewJulianDay(day float64) JulianDay {
	return JulianDay{day: day}
}

// JulianDayFromYmdHms 从年月日时分秒创建儒略日
func JulianDayFromYmdHms(year, month, day, hour, minute, second int) JulianDay {
	d := float64(day) + ((float64(second)/60+float64(minute))/60+float64(hour))/24
	n := 0
	g := year*372+month*31+int(d) >= 588829
	if month <= 2 {
		month += 12
		year--
	}
	if g {
		n = int(float64(year) * 0.01)
		n = 2 - n + int(float64(n)*0.25)
	}
	return NewJulianDay(float64(int(365.25*float64(year+4716))) + float64(int(30.6001*float64(month+1))) + d + float64(n) - 1524.5)
}

func (o JulianDay) GetDay() float64                   { return o.day }
func (o JulianDay) Next(n int) JulianDay              { return NewJulianDay(o.day + float64(n)) }
func (o JulianDay) Subtract(target JulianDay) float64 { return o.day - target.GetDay() }

// GetSolarDay 转换为公历日
func (o JulianDay) GetSolarDay() SolarDay {
	d := int(o.day + 0.5)
	f := o.day + 0.5 - float64(d)
	if d >= 2299161 {
		c := int((float64(d) - 1867216.25) / 36524.25)
		d += 1 + c - int(float64(c)*0.25)
	}
	d += 1524
	y := int((float64(d) - 122.1) / 365.25)
	d -= int(365.25 * float64(y))
	m := int(float64(d) / 30.601)
	d -= int(30.601 * float64(m))
	if m > 13 {
		m -= 12
	} else {
		y -= 1
	}
	m -= 1
	y -= 4715
	f *= 24
	hour := int(f)
	f -= float64(hour)
	f *= 60
	minute := int(f)
	f -= float64(minute)
	f *= 60
	second := int(math.Round(f))
	if second >= 60 {
		minute++
		second -= 60
	}
	if minute >= 60 {
		hour++
		minute -= 60
	}
	return SolarDay{year: y, month: m, day: d, hour: hour, minute: minute, second: second}
}

// ============ 公历日 ============

// SolarDay 公历日
type SolarDay struct {
	year   int
	month  int
	day    int
	hour   int
	minute int
	second int
}

// NewSolarDay 创建公历日
func NewSolarDay(year, month, day int) (SolarDay, error) {
	if month < 1 || month > 12 {
		return SolarDay{}, fmt.Errorf("非法月份: %d", month)
	}
	maxDay := GetSolarMonthDays(year, month)
	if day < 1 || day > maxDay {
		return SolarDay{}, fmt.Errorf("非法日期: %d-%d-%d", year, month, day)
	}
	return SolarDay{year: year, month: month, day: day}, nil
}

// NewSolarDayFromTime 从time.Time创建公历日
func NewSolarDayFromTime(t time.Time) SolarDay {
	return SolarDay{
		year:   t.Year(),
		month:  int(t.Month()),
		day:    t.Day(),
		hour:   t.Hour(),
		minute: t.Minute(),
		second: t.Second(),
	}
}

// Today 获取今天
func Today() SolarDay {
	return NewSolarDayFromTime(time.Now())
}

func (o SolarDay) GetYear() int  { return o.year }
func (o SolarDay) GetMonth() int { return o.month }
func (o SolarDay) GetDay() int   { return o.day }

// String 字符串表示
func (o SolarDay) String() string {
	return fmt.Sprintf("%d年%d月%d日", o.year, o.month, o.day)
}

// IsLeapYear 是否闰年
func IsLeapYear(year int) bool {
	if year < 1600 {
		return year%4 == 0
	}
	return (year%4 == 0 && year%100 != 0) || year%400 == 0
}

// GetSolarMonthDays 获取某月天数
func GetSolarMonthDays(year, month int) int {
	days := []int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if month == 2 && IsLeapYear(year) {
		return 29
	}
	return days[month-1]
}

func (o SolarDay) GetJulianDay() JulianDay {
	return JulianDayFromYmdHms(o.year, o.month, o.day, o.hour, o.minute, o.second)
}

func (o SolarDay) Next(n int) SolarDay { return o.GetJulianDay().Next(n).GetSolarDay() }

// Equals 是否相同日期
func (o SolarDay) Equals(target SolarDay) bool {
	return o.year == target.year && o.month == target.month && o.day == target.day
}

// Subtract 日期相减，获取天数差
func (o SolarDay) Subtract(target SolarDay) int {
	return int(o.GetJulianDay().Subtract(target.GetJulianDay()))
}

// GetLunarDay 获取农历日
func (o SolarDay) GetLunarDay() LunarDay {
	m, _ := NewLunarMonth(o.year, o.month)
	days := o.Subtract(m.GetFirstJulianDay().GetSolarDay())
	for days < 0 {
		m = m.Next(-1)
		days += m.GetDayCount()
	}
	d, _ := NewLunarDay(m.GetYear(), m.GetMonthWithLeap(), days+1)
	return d
}

// GetSolarFestival 获取公历节日
func (o SolarDay) GetSolarFestival() *SolarFestival {
	f, _ := GetSolarFestivalByYmd(o.year, o.month, o.day)
	return f
}

// GetLunarFestival 获取农历节日
func (o SolarDay) GetLunarFestival() *LunarFestival {
	lunarDay := o.GetLunarDay()
	f, _ := GetLunarFestivalByYmd(lunarDay.GetYear(), lunarDay.GetMonth(), lunarDay.GetDay())
	return f
}

// GetSolarTerm 获取当天节气
func (o SolarDay) GetSolarTerm() *SolarTerm {
	y := o.year
	for i := 0; i < 24; i++ {
		term := NewSolarTermFromIndex(y, i)
		if term.GetSolarDay().Equals(o) {
			return &term
		}
	}
	// 检查下一年的前几个节气
	for i := 0; i < 2; i++ {
		term := NewSolarTermFromIndex(y+1, i)
		if term.GetSolarDay().Equals(o) {
			return &term
		}
	}
	return nil
}

// ============ 公历节日 ============

// SolarFestivalNames 公历节日名称
var SolarFestivalNames = []string{"元旦", "三八妇女节", "植树节", "五一劳动节", "五四青年节", "六一儿童节", "建党节", "八一建军节", "教师节", "国庆节"}

// SolarFestivalData 公历节日数据 格式: @索引类型月日起始年
var SolarFestivalData = "@00001011950@01003081950@02003121979@03005011950@04005041950@05006011950@06007011941@07008011933@08009101985@09010011950"

// SolarFestival 公历节日
type SolarFestival struct {
	name string
}

// GetSolarFestivalByYmd 根据年月日获取公历节日
func GetSolarFestivalByYmd(year int, month int, day int) (*SolarFestival, error) {
	re, err := regexp.Compile(fmt.Sprintf("@\\d{2}0%02d%02d\\d+", month, day))
	if err != nil {
		return nil, err
	}
	data := re.FindString(SolarFestivalData)
	if data == "" {
		return nil, nil
	}
	startYear, _ := strconv.Atoi(data[8:])
	if year < startYear {
		return nil, nil
	}
	index, _ := strconv.Atoi(data[1:3])
	return &SolarFestival{name: SolarFestivalNames[index]}, nil
}

// GetName 获取名称
func (o SolarFestival) GetName() string {
	return o.name
}

// ============ 节气 ============

// SolarTermNames 节气名称（从冬至开始）
var SolarTermNames = []string{"冬至", "小寒", "大寒", "立春", "雨水", "惊蛰", "春分", "清明", "谷雨", "立夏", "小满", "芒种", "夏至", "小暑", "大暑", "立秋", "处暑", "白露", "秋分", "寒露", "霜降", "立冬", "小雪", "大雪"}

// SolarTerm 节气
type SolarTerm struct {
	name             string
	cursoryJulianDay float64
}

// NewSolarTermFromIndex 根据年份和索引创建节气
func NewSolarTermFromIndex(year int, index int) SolarTerm {
	size := 24
	idx := index % size
	if idx < 0 {
		idx += size
	}
	y := year + (index / size)
	if index < 0 && idx != 0 {
		y--
	}
	return SolarTerm{
		name:             SolarTermNames[idx],
		cursoryJulianDay: initTermByYear(y, idx),
	}
}

// GetName 获取名称
func (o SolarTerm) GetName() string {
	return o.name
}

// GetSolarDay 获取公历日
func (o SolarTerm) GetSolarDay() SolarDay {
	return NewJulianDay(o.cursoryJulianDay + J2000).GetSolarDay()
}

// initTermByYear 根据年份和节气索引初始化节气儒略日
func initTermByYear(year int, offset int) float64 {
	jd := math.Floor(float64(year-2000)*365.2422 + 180)
	w := math.Floor((jd-355+183)/365.2422)*365.2422 + 355
	if CalcQi(w) > jd {
		w -= 365.2422
	}
	return CalcQi(w + 15.2184*float64(offset))
}

// ============ 节日统一结构 ============

// FestivalTypeEnum 节日类型枚举
type FestivalTypeEnum int

const (
	// FestivalTypeSolar 公历节日
	FestivalTypeSolar FestivalTypeEnum = iota
	// FestivalTypeLunar 农历节日
	FestivalTypeLunar
	// FestivalTypeSolarTerm 节气
	FestivalTypeSolarTerm
)

// String 获取节日类型名称
func (t FestivalTypeEnum) String() string {
	switch t {
	case FestivalTypeSolar:
		return "公历节日"
	case FestivalTypeLunar:
		return "农历节日"
	case FestivalTypeSolarTerm:
		return "节气"
	default:
		return "未知"
	}
}

// Festival 节日
type Festival struct {
	Type     FestivalTypeEnum
	Name     string
	SolarDay SolarDay
}

// String 字符串表示
func (f Festival) String() string {
	return fmt.Sprintf("%s %s (%s)", f.SolarDay.String(), f.Name, f.Type.String())
}

// GetNearestFestival 获取最近的节日（向后查找）
// maxDays: 最大查找天数
func (o SolarDay) GetNearestFestival(maxDays int) *Festival {
	for i := 0; i <= maxDays; i++ {
		if f := o.Next(i).checkFestival(); f != nil {
			return f
		}
	}
	return nil
}

// checkFestival 检查当天是否有节日
func (o SolarDay) checkFestival() *Festival {

	// 检查公历节日
	if sf := o.GetSolarFestival(); sf != nil {
		return &Festival{Type: FestivalTypeSolar, Name: sf.GetName(), SolarDay: o}
	}
	// 检查农历节日
	if lf := o.GetLunarFestival(); lf != nil {
		return &Festival{Type: FestivalTypeLunar, Name: lf.GetName(), SolarDay: o}
	}
	// 检查节气
	if term := o.GetSolarTerm(); term != nil {
		return &Festival{Type: FestivalTypeSolarTerm, Name: term.GetName(), SolarDay: o}
	}
	return nil
}
