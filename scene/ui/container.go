// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package ui

import (
	"fmt"
)

type Options map[OptionType]interface{}

func (o Options) Validate() error {
	for opt, value := range o {
		if !opt.Valid() {
			return fmt.Errorf("Invalid option specified.")
		}
		if !opt.ValidValue(value) {
			return fmt.Errorf("Invalid option value specified.")
		}
	}
	return nil
}

type StateOptions map[StateType]Options

func (s StateOptions) Validate() error {
	for state, options := range s {
		if !state.Valid() {
			return fmt.Errorf("Invalid state type specified.")
		}

		err := options.Validate()
		if err != nil {
			return err
		}
	}
	return nil
}
