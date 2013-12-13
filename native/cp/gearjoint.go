// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package cp

/*
#include "chipmunk/chipmunk.h"
*/
import "C"

import (
	"unsafe"
)

// Check if a constraint is a ratchet joint.
func (c *Constraint) IsGearJoint() bool {
	return goBool(C.cpConstraintIsGearJoint(c.c))
}

// Allocate and initialize a gear joint.
func GearJointNew(a, b *Body, phase, ratio float64) *Constraint {
	c := new(Constraint)
	c.c = C.cpGearJointNew(
		(*C.cpBody)(unsafe.Pointer(a)),
		(*C.cpBody)(unsafe.Pointer(b)),
		C.cpFloat(phase),
		C.cpFloat(ratio),
	)
	if c.c == nil {
		return nil
	}
	C.cpConstraintSetUserData(c.c, C.cpDataPointer(unsafe.Pointer(c)))
	return c
}

// Get the phase offset of the gears.
func (c *Constraint) GearJointPhase() float64 {
	return float64(C.cpGearJointGetPhase(c.c))
}

// Set the phase offset of the gears.
func (c *Constraint) GearJointSetPhase(phase float64) {
	C.cpGearJointSetPhase(c.c, C.cpFloat(phase))
}

// Get the angular distance of each ratchet.
func (c *Constraint) GearJointRatio() float64 {
	return float64(C.cpGearJointGetRatio(c.c))
}

// Set the ratio of a gear joint.
func (c *Constraint) GearJointSetRatio(ratio float64) {
	C.cpGearJointSetRatio(c.c, C.cpFloat(ratio))
}
