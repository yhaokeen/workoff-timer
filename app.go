package main

import (
	"context"
	"fmt"

	"workoff-timer/internal/festival"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// FestivalInfo 节日信息结构（返回给前端）
type FestivalInfo struct {
	Name string `json:"name"`
	Days int    `json:"days"`
	Type string `json:"type"`
}

// GetNextFestival 获取下一个节日
func (a *App) GetNextFestival() *FestivalInfo {
	// 从今天开始查找60天内最近的节日
	f := festival.Today().GetNearestFestival(60)
	if f == nil {
		return &FestivalInfo{
			Name: "无",
			Days: 0,
			Type: "",
		}
	}

	// 计算距离天数
	today := festival.Today()
	days := int(f.SolarDay.GetJulianDay().Subtract(today.GetJulianDay()))

	return &FestivalInfo{
		Name: f.Name,
		Days: days,
		Type: f.Type.String(),
	}
}
