package schema

import (
	"time"
)

type JsonTime time.Time

func (jt JsonTime) MarshalJSON() ([]byte, error) {
	if time.Time(jt).IsZero() {
		return []byte(`""`), nil
	}
	return []byte(`"` + time.Time(jt).Format("2006-01-02 15:04:05") + `"`), nil
}

func (jt JsonTime) Now() JsonTime {
	return JsonTime(time.Now())
}

type PagingParam struct {
	Page int `form:"page,default=1"`
	Size int `form:"size,default=10" binding:"max=500"`
}

func (p *PagingParam) Offset() int {
	return (p.Page - 1) * p.Size
}
