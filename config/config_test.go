// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package config

import (
	"bytes"
	"testing"
)

func TestConfig(t *testing.T) {
	c := Parse([]byte(`
	# this is an comment, anything (#===#) may be here.

	# string value
	uniqueKey = value

	# string value
	uniqueKey2 = strings can contain anything (#===#) and run until the end of the line.

	# override previous uniqueKey2 value
	uniqueKey2 = Comments following an string on the same line. #Are part of the string

	# float64 value
	uniqueKey3= 1.0

	# Also float64 value
	1 =1

	# bool value
	doSomething=true
	`))

	var (
		uniqueKey, uniqueKey2 string
		one, uniqueKey3       float64
		ok, doSomething       bool
	)

	ok = c.StringVar(&uniqueKey, "uniqueKey")
	if !ok || uniqueKey != "value" {
		t.Logf("%q\n", uniqueKey)
		t.Log("String parsing failed")
		t.Fail()
	}

	ok = c.StringVar(&uniqueKey2, "uniqueKey2")
	if !ok || uniqueKey2 != "Comments following an string on the same line. #Are part of the string" {
		t.Logf("%q\n", uniqueKey2)
		t.Log("String parsing failed")
		t.Fail()
	}

	ok = c.Float64Var(&uniqueKey3, "uniqueKey3")
	if !ok || uniqueKey3 != 1.0 {
		t.Log("Float64 parsing failed")
		t.Fail()
	}

	ok = c.Float64Var(&one, "1")
	if !ok || one != 1 {
		t.Log("Float64 parsing failed")
		t.Fail()
	}

	ok = c.BoolVar(&doSomething, "doSomething")
	if !ok || doSomething != true {
		t.Log("Bool parsing failed")
		t.Fail()
	}
}

func TestMessyConfig(t *testing.T) {
	c := Parse([]byte(`
###########################comment
key with spaces = 1.0

key#with#hash = 1.0

properPair1 = 1

keyWithNoValue=

properPair2 = 2

keyWithNoEqual

properPair3 = 3

overTwoLines
=
oopsmybad

properPair4 = 4
	`))

	if v, ok := c.Float64("key with spaces"); !ok || v != 1.0 {
		t.Log("key with spaces parsing failed!", ok)
		t.Fail()
	}

	if v, ok := c.Float64("key#with#hash"); !ok || v != 1.0 {
		t.Log("key#with#hash parsing failed!")
		t.Fail()
	}

	if v, ok := c.Float64("properPair1"); !ok || v != 1 {
		t.Log("Hash key parsing failed!")
		t.Fail()
	}

	if v, ok := c.Float64("properPair2"); !ok || v != 2 {
		t.Log("No value key parsing failed!")
		t.Fail()
	}

	if v, ok := c.Float64("properPair3"); !ok || v != 3 {
		t.Log("No equals key parsing failed!")
		t.Fail()
	}

	if v, ok := c.Float64("properPair4"); !ok || v != 4 {
		t.Log("Over two lines parsing failed!")
		t.Fail()
	}
}

func TestWriteValidAndCleanup(t *testing.T) {
	data := []byte(`
###########################comment
key with spaces = 1.0

key#with#hash = 1.0

properPair1 = 1

keyWithNoValue=

properPair2 = 2

keyWithNoEqual

properPair3 = 3

overTwoLines
=
oopsmybad

properPair4 = 4
	`)

	c := Parse(data)

	buf := new(bytes.Buffer)
	c.Write(buf)

	t.Log("(Please compare below)")
	t.Log("* Before *")
	t.Log("\n" + string(data))

	t.Log("* After *")
	t.Log("\n" + string(buf.Bytes()))
}
