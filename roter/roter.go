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
	if r.Offset >= 0 && r.Offset < 3 {
		//update offset
		r.Offset += 1

		//rotate
		start_value := r.Values[0]
		for index := range r.Values {
			if index+1 == len(r.Values) {
				break
			}
			r.Values[index] = r.Values[index+1]
		}
		r.Values[len(r.Values)-1] = start_value

		//update one round flag
		r.IsOneRound = false
	} else if r.Offset == 3 {
		//update offset
		r.Offset = 0

		//rotate
		start_value := r.Values[0]
		for index := range r.Values {
			if index+1 == len(r.Values) {
				break
			}
			r.Values[index] = r.Values[index+1]
		}
		r.Values[len(r.Values)-1] = start_value

		//update one round flag
		r.IsOneRound = true
	} else {
		panic(errors.New("Invalid offset: " + strconv.Itoa(r.Offset)))
	}
}

func Initialize(values []int) *Roter {
	return &Roter{
		Offset:     0,
		Values:     values,
		IsOneRound: false,
	}
}
