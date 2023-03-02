package level

import "time"

type ScreenText struct {
	x, y        int
	text        string
	expiredTime time.Time
}

func NewScreenText(x, y int, text string, expiredTime time.Time) ScreenText {
	return ScreenText{
		x:           x,
		y:           y,
		text:        text,
		expiredTime: expiredTime,
	}
}

func (s *ScreenText) X() int {
	return s.x
}

func (s *ScreenText) Y() int {
	return s.y
}

func (s *ScreenText) Text() string {
	return s.text
}

func (s *ScreenText) ExpiredTime() time.Time {
	return s.expiredTime
}
