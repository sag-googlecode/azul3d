// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package ui

type LayoutType int

const (
	Horizontal LayoutType = iota
	Vertical
)

func (l LayoutType) Valid() bool {
	if l == Horizontal || l == Vertical {
		return true
	}
	return false
}
