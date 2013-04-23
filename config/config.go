// Package config implements an simple configuration file reader and writer.
//
// The configuration file format implemented by this package is designed to be very user friendly
// specifically towards non computer-oriented people, there are no 'sections' like in INI files,
// and likewise it is mostly line-based in order to recover from bad user input more easily.
//
// The configuration file format looks like the following:
//
//  # this is an comment, anything (#===#) may be here.
//
//  # string value
//  uniqueKey = value
//
//  # string value
//  uniqueKey2 = strings can contain anything (#===#) and run until the end of the line.
//
//  # override previous uniqueKey2 value
//  uniqueKey2 = Comments following an string on the same line. #Are part of the string
//
//  # float64 value
//  uniqueKey3= 1.0
//
//  # Also float64 value
//  1 =1
//
//  # bool value
//  doSomething=true
//
// Things you should have noticed:
//
// 1. White space is ignored except in strings.
//
// 2. Once an comment (#) is encountered, the rest of the line from that character on is ignored.
//
// 3. An value is determined to be an string when it cannot be parsed as an float64 or bool.
//
// 4. Keys cannot contain equals (=) signs inside them, but can contain spaces and tabs.
package config

import (
	"strconv"
	"sync"
)

type line struct {
	comment, newline bool
	key              string
	value            interface{}
}

// Config represents access to an configuration file, it can be manipulated or wrote to an
// io.Writer
type Config struct {
	access sync.RWMutex
	values map[string]interface{}
	lines  []*line
}

// Set specifies the value of an existing (or non existing) key within this Config.
func (c *Config) Set(key string, value interface{}) {
	c.access.Lock()
	defer c.access.Unlock()

	c.values[key] = value
}

// Get returns the value related to the key, or returns nil and false if no such key exists within
// this Config.
func (c *Config) Get(key string) (value interface{}, ok bool) {
	c.access.RLock()
	defer c.access.RUnlock()

	value, ok = c.values[key]
	return
}

// Keys returns all keys withing this Config
func (c *Config) Keys() []string {
	k := make([]string, len(c.values))
	i := 0
	for key, _ := range c.values {
		i++
		k[i] = key
	}
	return k
}

// StringVar stores the string of the specified key in the specified store variable.
//
// Returns false if the specified key does not exist.
func (c *Config) StringVar(store *string, key string) bool {
	v, ok := c.Get(key)
	if !ok {
		return false
	}

	switch v.(type) {
	case string:
		*store = v.(string)
		return true

	case bool:
		if v.(bool) {
			*store = "true"
		} else {
			*store = "false"
		}
		return true

	case float64:
		*store = strconv.FormatFloat(v.(float64), 'f', -1, 64)
		return true
	}
	return false
}

// String is equivilent to the following code:
//
//  var value string
//  ok := c.StringVar(&value, key)
//
func (c *Config) String(key string) (value string, ok bool) {
	ok = c.StringVar(&value, key)
	return
}

// Float64Var stores the float64 of the specified key in the specified store variable.
//
// Returns false if the specified key does not exist.
func (c *Config) Float64Var(store *float64, key string) bool {
	v, ok := c.Get(key)
	if !ok {
		return false
	}
	*store = v.(float64)
	return true
}

// Float64 is equivilent to the following code:
//
//  var value float64
//  ok := c.Float64Var(&value, key)
//
func (c *Config) Float64(key string) (value float64, ok bool) {
	ok = c.Float64Var(&value, key)
	return
}

// BoolVar stores the bool of the specified key in the specified store variable.
//
// Returns false if the specified key does not exist.
func (c *Config) BoolVar(store *bool, key string) bool {
	v, ok := c.Get(key)
	if !ok {
		return false
	}
	*store = v.(bool)
	return true
}

// Bool is equivilent to the following code:
//
//  var value bool
//  ok := c.BoolVar(&value, key)
//
func (c *Config) Bool(key string) (value bool, ok bool) {
	ok = c.BoolVar(&value, key)
	return
}
