package festival

import (
	"container/list"
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

// ============ 农历年 ============

var lunarYearLeap []*list.List
var onceLunarYear sync.Once

// LunarYear 农历年
type LunarYear struct {
	year int
}

// NewLunarYear 创建农历年
func NewLunarYear(year int) (LunarYear, error) {
	initLunarYearLeap()
	if year < -1 || year > 9999 {
		return LunarYear{}, fmt.Errorf("非法农历年: %d", year)
	}
	return LunarYear{year: year}, nil
}

func initLunarYearLeap() {
	onceLunarYear.Do(func() {
		chars := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ_@"
		months := []string{
			"080b0r0j0j0j0C0j0j0C0j0j0j0C0j0C0j0C0F0j0V0V0V0u0j0j0C0j0j0j0j0V0C0j1v0u0C0V1v0C0b080u110u0C0j0C1v9K1v2z0j1vmZbl1veN3s1v0V0C2S1v0V0C2S2o0C0j1Z1c2S1v0j1c0j2z1v0j1c0j392H0b2_2S0C0V0j1c0j2z0C0C0j0j1c0j0N250j0C0j0b081n080b0C0C0C1c0j0N",
			"0r1v1c1v0V0V0F0V0j0C0j0C0j0V0j0u1O0j0C0V0j0j0j0V0b080u0r0u080b0j0j0C0V0C0V0j0b080V0u080b0j0j0u0j1v0u080b1c0j080b0j0V0j0j0V0C0N1v0j1c0j0j1v2g1v420j1c0j2z1v0j1v5Q9z1v4l0j1vfn1v420j9z4l1v1v2S1c0j1v2S3s1v0V0C2S1v1v2S1c0j1v2S2_0b0j2_2z0j1c0j",
			"0z0j0j0j0C0j0j0C0j0j0j0C0j0C0j0j0j0j0m0j0C0j0j0C0j0j0j0j0b0V0j0j0C0j0j0j0j0V0j0j0j0V0b0V0V0C0V0C0j0j0b080u110u0V0C0j0N0j0b080b080b0j0r0b0r0b0j0j0j0j0C0j0b0r0C0j0b0j0C0C0j0j0j0j0j0j0j0j0j0b110j0b0j0j0j0C0j0C0j0j0j0j0b080b080b0V080b080b0j0j0j0j0j0j0V0j0j0u1v0j0j0j0C0j0j0j0V0C0N1c0j0C0C0j0j0j1n080b0j0V0C0j0C0C2g0j1c0j0j1v2g1v0j0j1v7N0j1c0j3L0j0j1v5Q1Z5Q1v4lfn1v420j1v5Q1Z5Q1v4l1v2z1v",
			"0H140r0N0r140r0u0r0V171c11140C0j0u110j0u0j1v0j0C0j0j0j0b080V0u080b0C1v0j0j0j0C0j0b080V0j0j0b080b0j0j0j0j0b080b0C080j0b080b0j0j0j0j0j0j0b080j0b080C0b080b080b080b0j0j0j0j080b0j0C0j0j0j0b0j0j080C0b0j0j0j0j0j0j0b08080b0j0C0j0j0j0b0j0j0K0b0j0C0j0j0j0b080b080j0C0b0j080b080b0j0j0j0j080b0j0b0r0j0j0j0b0j0C0r0b0j0j0j0j0j0j0j0b080j0b0r0C0j0b0j0j0j0r0b0j0C0j0j0j0u0r0b0C0j080b0j0j0j0j0j0j0j1c0j0b0j0j0j0C0j0j0j0j0j0j0j0b080j1c0u0j0j0j0C0j1c0j0u0j1c0j0j0j0j0j0j0j0j1c0j0u1v0j0j0V0j0j2g0j0j0j0C1v0C1G0j0j0V0C1Z1O0j0V0j0j2g1v0j0j0V0C2g5x1v4l1v421O7N0V0C4l1v2S1c0j1v2S2_",
			"050b080C0j0j0j0C0j0j0C0j0j0j0C0j0C0j0C030j0j0j0j0j0j0j0j0j0C0j0b080u0V080b0j0j0V0j0j0j0j0j0j0j0j0j0V0N0j0C0C0j0j0j0j0j0j0j0j1c0j0u0j1v0j0j0j0j0j0b080b080j0j0j0b080b080b080b080b0j0j0j080b0j0b080j0j0j0j0b080b0j0j0r0b080b0b080j0j0j0j0b080b080j0b080j0b080b080b080b080b0j0j0r0b0j0b080j0j0j0j0b080b0j0j0C080b0b080j0j0j0j0j0j0j0b080u080j0j0b0j0j0j0C0j0b080j0j0j0j0b080b080b080b0C080b080b080b0j0j0j0j0j0j0b0C080j0j0b0j0j0j0C0j0b080j0j0C0b080b080j0b0j0j0C080b0j0j0j0j0j0j0b0j0j080C0b0j080b0j0j0j0j0j0j0j0C0j0j0j0b0j0j0C080b0j0j0j0j0j0j0b080b080b0K0b080b080b0j0j0j0j0j0j0j0C0j0j0u0j0j0V0j080b0j0C0j0j0j0b0j0r0C0b0j0j0j0j0j0j0j0j0j0C0j0b080b080b0j0C0C0j0C0j0j0j0u110u0j0j0j0j0j0j0j0j0C0j0j0u0j1c0j0j0j0j0j0j0j0j0V0C0u0j0C0C0V0C1Z0j0j0j0C0j0j0j1v0u0j1c0j0j0j0C0j0j2g0j1c1v0C1Z0V0j4l0j0V0j0j2g0j1v0j1v2S1c7N1v",
			"0w0j1c0j0V0j0j0V0V0V0j0m0V0j0C1c140j0j0j0C0V0C0j1v0j0N0j0C0j0j0j0V0j0j1v0N0j0j0V0j0j0j0j0j0j080b0j0j0j0j0j0j0j080b0j0C0j0j0j0b0j0j080u080b0j0j0j0j0j0j0b080b080b080C0b0j080b080b0j0j0j0j080b0j0C0j0j0j0b0j0j080u080b0j0j0j0j0j0j0b080b080b080b0r0b0j080b080b0j0j0j0j080b0j0b0r0j0j0b080b0j0j080b0j080b0j080b080b0j0j0j0j0j0b080b0r0C0b080b0j0j0j0j080b0b080b080j0j0j0b080b080b080b0j0j0j0j080b0j0b080j0j0j0j0b080b0j0j0r0b080b0j0j0j0j0j0b080b080j0b0r0b080j0b080b0j0j0j0j080b0j0b080j0j0j0j0b080b0j080b0r0b0j080b080b0j0j0j0j0j0b080b0r0C0b080b0j0j0j0j0j0j0b080j0j0j0b080b080b080b0j0j0j0r0b0j0b080j0j0j0j0b080b0r0b0r0b0j080b080b0j0j0j0j0j0j0b0r0j0j0j0b0j0j0j0j080b0j0b080j0j0j0j0b080b080b0j0r0b0j080b0j0j0j0j0j0j0j0b0r0C0b0j0j0j0j0j0j0j080b0j0C0j0j0j0b0j0C0r0b0j0j0j0j0j0j0b080b080u0r0b0j080b0j0j0j0j0j0j0j0b0r0C0u0j0j0j0C0j080b0j0C0j0j0j0u110b0j0j0j0j0j0j0j0j0j0C0j0b080b0j0j0C0C0j0C0j0j0j0b0j1c0j080b0j0j0j0j0j0j0V0j0j0u0j1c0j0j0j0C0j0j2g0j0j0j0C0j0j0V0j0b080b1c0C0V0j0j2g0j0j0V0j0j1c0j1Z0j0j0C0C0j1v",
			"160j0j0V0j1c0j0C0j0C0j1f0j0V0C0j0j0C0j0j0j1G080b080u0V080b0j0j0V0j1v0j0u0j1c0j0j0j0C0j0j0j0C0C0j1D0b0j080b0j0j0j0j0C0j0b0r0C0j0b0j0C0C0j0j0j0j0j0j0j0j0j0b0r0b0r0j0b0j0j0j0C0j0b0r0j0j0j0b080b080j0b0C0j080b080b0j0j0j0j0j0j0b0C080j0j0b0j0j0j0C0j0b080j0j0j0j0b080b080j0b0C0r0j0b0j0j0j0j0j0j0b0C080j0j0b0j0j0j0C0j0j0j0j0C0j0j0b080b0j0j0C080b0j0j0j0j0j0j0b080b080b080C0b080b080b080b0j0j0j0j0j0b080C0j0j0b080b0j0j0C080b0j0j0j0j0j0j0b080j0b0C080j0j0b0j0j0j0j0j0j0b080j0b080C0b080b080b080b0j0j0j0j080b0j0C0j0j0b080b0j0j0C080b0j0j0j0j0j0j0b080j0b080u080j0j0b0j0j0j0j0j0j0b080C0j0j0b080b0j0j0C0j0j080b0j0j0j0j0j0b080b0C0r0b080b0j0j0j0j0j0j0b080j0b080u080b080b080b0j0j0j0C0j0b080j0j0j0j0b0j0j0j0C0j0j080b0j0j0j0j0j0b080b0C0r0b080b0j0j0j0j0j0j0b080j0b0r0b080b080b080b0j0j0j0r0b0j0b0r0j0j0j0b0j0j0j0r0b0j080b0j0j0j0j0j0j0j0b0r0C0b0j0j0j0j0j0j0j0b080j0C0u080b080b0j0j0j0r0b0j0C0C0j0b0j110b0j080b0j0j0j0j0j0j0u0r0C0b0j0j0j0j0j0j0j0j0j0C0j0j0j0b0j1c0j0C0j0j0j0b0j0814080b080b0j0j0j0j0j0j1c0j0u0j0j0V0j0j0j0j0j0j0j0u110u0j0j0j",
			"020b0r0C0j0j0j0C0j0j0V0j0j0j0j0j0C0j1f0j0C0j0V1G0j0j0j0j0V0C0j0C1v0u0j0j0j0V0j0j0C0j0j0j1v0N0C0V0j0j0j0K0C250b0C0V0j0j0V0j0j2g0C0V0j0j0C0j0j0b081v0N0j0j0V0V0j0j0u0j1c0j080b0j0j0j0j0j0j0V0j0j0u0j0j0V0j0j0j0C0j0b080b080V0b0j080b0j0j0j0j0j0j0j0b0r0C0j0b0j0j0j0C0j080b0j0j0j0j0j0j0u0r0C0u0j0j0j0j0j0j0b080j0C0j0b080b080b0j0C0j080b0j0j0j0j0j0j0b080b110b0j0j0j0j0j0j0j0j0j0b0r0j0j0j0b0j0j0j0r0b0j0b080j0j0j0j0b080b080b080b0r0b0j080b080b0j0j0j0j0j0j0b0r0C0b080b0j0j0j0j080b0j0b080j0j0j0j0b080b080b0j0j0j0r0b0j0j0j0j0j0j0b080b0j080C0b0j080b080b0j0j0j0j080b0j0b0r0C0b080b0j0j0j0j080b0j0j0j0j0j0b080b080b080b0j0j080b0r0b0j0j0j0j0j0j0b0j0j080C0b0j080b080b0j0j0j0j0j0b080C0j0j0b080b0j0j0C0j0b080j0j0j0j0b080b080b080b0C0C080b0j0j0j0j0j0j0b0C0C080b080b080b0j0j0j0j0j0j0b0C080j0j0b0j0j0j0C0j0b080j0b080j0j0b080b080b080b0C0r0b0j0j0j0j0j0j0b080b0r0b0r0b0j080b080b0j0j0j0j0j0j0b0r0C0j0b0j0j0j0j0j0j0b080j0C0j0b080j0b0j0j0K0b0j0C0j0j0j0b080b0j0K0b0j080b0j0j0j0j0j0j0V0j0j0b0j0j0j0C0j0j0j0j",
			"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_1Z0V2o141v0j0C0C0j2z1v0j1c0j7N1O420j1c0j1v2S1c0j1v2S2_0b0j0V0j0j1v0N1v0j0j1c0j1v140j0V0j0j0C0C0b080u1v0C0V0u110u0j0j0j0C0j0j0j0C0C0N0C0V0j0j0C0j0j0b080u110u0C0j0C0u0r0C0u080b0j0j0C0j0j0j",
		}
		for _, m := range months {
			n := 0
			size := len(m) / 2
			l := list.New()
			for y := 0; y < size; y++ {
				z := y * 2
				t := 0
				c := 1
				for x := 1; x > -1; x-- {
					i := z + x
					t += c * strings.Index(chars, m[i:i+1])
					c *= 64
				}
				n += t
				l.PushBack(n)
			}
			lunarYearLeap = append(lunarYearLeap, l)
		}
	})
}

func (o LunarYear) GetYear() int { return o.year }

func (o LunarYear) GetMonthCount() int {
	if o.GetLeapMonth() < 1 {
		return 12
	}
	return 13
}

func (o LunarYear) Next(n int) LunarYear {
	y, _ := NewLunarYear(o.year + n)
	return y
}

func (o LunarYear) GetLeapMonth() int {
	if o.year == -1 {
		return 11
	}
	for i, value := range lunarYearLeap {
		for e := value.Front(); e != nil; e = e.Next() {
			if e.Value == o.year {
				return i + 1
			}
		}
	}
	return 0
}

// ============ 农历月 ============

var lunarMonthCache sync.Map

// LunarMonth 农历月
type LunarMonth struct {
	year           LunarYear
	month          int
	leap           bool
	dayCount       int
	indexInYear    int
	firstJulianDay JulianDay
}

// NewLunarMonth 创建农历月
func NewLunarMonth(year int, month int) (LunarMonth, error) {
	key := fmt.Sprintf("%d_%d", year, month)
	if c, ok := lunarMonthCache.Load(key); ok {
		return c.(LunarMonth), nil
	}

	currentYear, err := NewLunarYear(year)
	if err != nil {
		return LunarMonth{}, err
	}
	currentLeapMonth := currentYear.GetLeapMonth()
	if month == 0 || month > 12 || month < -12 {
		return LunarMonth{}, fmt.Errorf("非法农历月: %d", month)
	}
	leap := month < 0
	m := month
	if m < 0 {
		m = -m
	}
	if leap && m != currentLeapMonth {
		return LunarMonth{}, fmt.Errorf("农历%d年没有闰%d月", year, m)
	}

	dongZhiJd := initTermByYear(year, 0)
	w := CalcShuo(dongZhiJd)
	if w > dongZhiJd {
		w -= 29.53
	}

	offset := 2
	if year > 8 && year < 24 {
		offset = 1
	} else if year != 239 && year != 240 {
		y, _ := NewLunarYear(year - 1)
		if y.GetLeapMonth() > 10 {
			offset = 3
		}
	}

	index := m - 1
	if leap || (currentLeapMonth > 0 && m > currentLeapMonth) {
		index += 1
	}

	w += 29.5306 * float64(offset+index)
	firstDay := CalcShuo(w)

	result := LunarMonth{
		year:           currentYear,
		month:          m,
		leap:           leap,
		dayCount:       int(CalcShuo(w+29.5306) - firstDay),
		indexInYear:    index,
		firstJulianDay: NewJulianDay(J2000 + firstDay),
	}
	lunarMonthCache.Store(key, result)
	return result, nil
}

func (o LunarMonth) GetYear() int                 { return o.year.GetYear() }
func (o LunarMonth) GetMonth() int                { return o.month }
func (o LunarMonth) GetDayCount() int             { return o.dayCount }
func (o LunarMonth) GetFirstJulianDay() JulianDay { return o.firstJulianDay }

func (o LunarMonth) GetMonthWithLeap() int {
	if o.leap {
		return -o.month
	}
	return o.month
}

func (o LunarMonth) Next(n int) LunarMonth {
	if n == 0 {
		month, _ := NewLunarMonth(o.GetYear(), o.GetMonthWithLeap())
		return month
	}
	m := o.indexInYear + 1 + n
	y := o.year
	if n > 0 {
		for m > y.GetMonthCount() {
			m -= y.GetMonthCount()
			y = y.Next(1)
		}
	} else {
		for m <= 0 {
			y = y.Next(-1)
			m += y.GetMonthCount()
		}
	}
	leap := false
	leapMonth := y.GetLeapMonth()
	if leapMonth > 0 {
		if m == leapMonth+1 {
			leap = true
		}
		if m > leapMonth {
			m--
		}
	}
	if leap {
		m = -m
	}
	month, _ := NewLunarMonth(y.GetYear(), m)
	return month
}

// ============ 农历日 ============

// LunarDay 农历日
type LunarDay struct {
	month LunarMonth
	day   int
}

// NewLunarDay 创建农历日
func NewLunarDay(year int, month int, day int) (LunarDay, error) {
	m, err := NewLunarMonth(year, month)
	if err != nil {
		return LunarDay{}, err
	}
	if day < 1 || day > m.GetDayCount() {
		return LunarDay{}, fmt.Errorf("非法农历日: %d年%d月%d日", year, month, day)
	}
	return LunarDay{month: m, day: day}, nil
}

func (o LunarDay) GetYear() int        { return o.month.GetYear() }
func (o LunarDay) GetMonth() int       { return o.month.GetMonthWithLeap() }
func (o LunarDay) GetMonthValue() int  { return o.month.GetMonth() }
func (o LunarDay) GetDay() int         { return o.day }
func (o LunarDay) Next(n int) LunarDay { return o.GetSolarDay().Next(n).GetLunarDay() }

// GetSolarDay 获取公历日
func (o LunarDay) GetSolarDay() SolarDay {
	return o.month.GetFirstJulianDay().Next(o.day - 1).GetSolarDay()
}

// ============ 农历节日 ============

// LunarFestivalNames 农历节日名称
var LunarFestivalNames = []string{"春节", "元宵节", "龙头节", "上巳节", "清明节", "端午节", "七夕节", "中元节", "中秋节", "重阳节", "冬至节", "腊八节", "除夕"}

// LunarFestivalData 农历节日数据 格式: @索引类型(0=日期,1=节气,2=除夕)数据
var LunarFestivalData = "@0000101@0100115@0200202@0300303@04107@0500505@0600707@0700715@0800815@0900909@10124@1101208@122"

// LunarFestival 农历节日
type LunarFestival struct {
	name string
}

// GetLunarFestivalByYmd 根据农历年月日获取农历节日
func GetLunarFestivalByYmd(year int, month int, day int) (*LunarFestival, error) {
	// 检查日期类型节日
	re, _ := regexp.Compile(fmt.Sprintf("@\\d{2}0%02d%02d", month, day))
	data := re.FindString(LunarFestivalData)
	if data != "" {
		index, _ := strconv.Atoi(data[1:3])
		return &LunarFestival{name: LunarFestivalNames[index]}, nil
	}

	// 检查节气类型节日（清明节）
	re, _ = regexp.Compile("@\\d{2}1\\d{2}")
	arr := re.FindAllString(LunarFestivalData, -1)
	for _, data := range arr {
		i, _ := strconv.Atoi(data[4:])
		solarTerm := NewSolarTermFromIndex(year, i)
		d := solarTerm.GetSolarDay().GetLunarDay()
		if d.GetYear() == year && d.GetMonth() == month && d.GetDay() == day {
			index, _ := strconv.Atoi(data[1:3])
			return &LunarFestival{name: LunarFestivalNames[index]}, nil
		}
	}

	// 检查除夕
	d, err := NewLunarDay(year, month, day)
	if err != nil {
		return nil, nil
	}
	nextDay := d.Next(1)
	if nextDay.GetMonthValue() == 1 && nextDay.GetDay() == 1 {
		return &LunarFestival{name: "除夕"}, nil
	}
	return nil, nil
}

// GetName 获取名称
func (o LunarFestival) GetName() string {
	return o.name
}
