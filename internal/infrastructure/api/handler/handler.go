package handler

import (
	"context"
	"fmt"

	"github.com/ninja-dark/QSOFT-task/internal/logic"
)


type Handlers struct {
	l *logic.ServiceLogic
}

func NewHandlers(l *logic.ServiceLogic) *Handlers {
	r := &Handlers{
		l:l,
	}
	return r
}
type Days struct{
	Message string
	Count int
}


func (rt *Handlers) GetCountDays(ctx context.Context, year int)(Days, error){
	date, err := rt.l.GetCountDays(year)
	if err != nil{
		return Days{}, fmt.Errorf("can't count days: %w", err)
	}
	return Days{
		Message: date.Message,
		Count: date.NumberOfDays,
	}, nil
} 