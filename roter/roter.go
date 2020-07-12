package roter

import (
	"errors"
	"strconv"
)

type Roter struct {
	Offset     int
	Values     []int
	IsOneRound bool
}

func (r *Roter) Plus() {
	if r.Offset >= 0 && r.Offset < 25 {
		//update offset
		r.Offset += 1

		//rotate
		r.rotate()

		//update one round flag
		r.IsOneRound = false
	} else if r.Offset == 25 {
		//update offset
		r.Offset = 0

		//rotate
		r.rotate()

		//update one round flag
		r.IsOneRound = true
	} else {
		panic(errors.New("Invalid offset: " + strconv.Itoa(r.Offset)))
	}
}

func (r *Roter) RepeatRotation(count int) {
	for i := 0; i < count; i++ {
		r.rotate()
	}
}

func (r *Roter) rotate() {
	start_value := r.Values[0]
	for index := range r.Values {
		if index+1 == len(r.Values) {
			break
		}
		r.Values[index] = r.Values[index+1]
	}
	r.Values[len(r.Values)-1] = start_value
}

func Initialize(values []int) *Roter {
	return &Roter{
		Offset:     0,
		Values:     values,
		IsOneRound: false,
	}
}
