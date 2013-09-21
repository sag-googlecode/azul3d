// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package ui

type StateType int

const (
	Default StateType = iota
	Hover
	Click
	Active
)

func (s StateType) Valid() bool {
	if s == Default || s == Hover || s == Click || s == Active {
		return true
	}
	return false
}
