package logic

import (
	"fmt"
	"math"

	"time"

	entitie "github.com/ninja-dark/QSOFT-task/internal/entities"
)

type ServiceLogic struct{
	
}
func New() *ServiceLogic{
	return &ServiceLogic{}
}
//GetCountDays - функция, которая возвращает количество дней
func (l *ServiceLogic)GetCountDays(year int)(entitie.Days, error){
	month, day := 1,1
	t1 := l.Date(year, month, day)
	t2 := time.Now()
	fmt.Println(t1)
	fmt.Println(t2)

	days:= math.Abs(t2.Sub(t1).Hours()/24)
	// получаем количество дней
	isT1AfterT2 := t1.After(t2)
	if isT1AfterT2{
		s := "Days left:"
		return entitie.Days{
			Message: s,
			NumberOfDays: int(days),
		}, nil
	}else{
		s := "Days done:"
		return entitie.Days{
			Message: s,
			NumberOfDays: int(days),
		}, nil
	}
}
//Date - функция, которая возвращает год в формате time.Time
func(l *ServiceLogic)Date (year, month, day int) time.Time {
	return time.Date(year, time.Month(month), day, 0,0,0,0, time.UTC)
}