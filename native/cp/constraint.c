// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

#include "_cgo_export.h"

void pre_go_chipmunk_constraint_pre_solve_func(cpConstraint *constraint, cpSpace *space) {
	go_chipmunk_constraint_pre_solve_func((void*)constraint, (void*)space);
}

void pre_go_chipmunk_constraint_post_solve_func(cpConstraint *constraint, cpSpace *space) {
	go_chipmunk_constraint_post_solve_func((void*)constraint, (void*)space);
}

