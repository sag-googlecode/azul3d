// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package text

import (
	"sync"
)

var (
	tagMap       map[string]*GlyphOptions
	tagMapAccess sync.RWMutex
)

func init() {
	tagMap = make(map[string]*GlyphOptions)
}

func SetTag(tag string, options *GlyphOptions) {
	tagMapAccess.Lock()
	defer tagMapAccess.Unlock()

	tagMap[tag] = options
}

func Tag(tag string) (options *GlyphOptions, ok bool) {
	tagMapAccess.RLock()
	defer tagMapAccess.RUnlock()

	options, ok = tagMap[tag]
	return
}

func Tags() map[string]*GlyphOptions {
	tagMapAccess.RLock()
	defer tagMapAccess.RUnlock()

	cpy := make(map[string]*GlyphOptions, len(tagMap))
	for k, v := range tagMap {
		cpy[k] = v
	}
	return cpy
}
