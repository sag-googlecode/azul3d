// Copyright 2014 The Azul3D Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
//
// Code generated by this program is also under
// the above license.

// glwrap is a tool for generating Go OpenGL wrappers.
package main

import (
	"encoding/xml"
)

type Registry struct {
	XMLName    xml.Name   `xml:"registry"`
	Comment    Comment    `xml:"comment"`
	Types      Types      `xml:"types"`
	Groups     Groups     `xml:"groups"`
	Enums      []Enums    `xml:"enums"`
	Commands   Commands   `xml:"commands"`
	Feature    []Feature  `xml:"feature"`
	Extensions Extensions `xml:"extensions"`
}

type Comment struct {
	String string `xml:",chardata"`
}

type Types struct {
	Type []Type `xml:"type"`
}

type Type struct {
	Raw        string `xml:",innerxml"`
	Definition string `xml:",chardata"`
	Name       string `xml:"name"`
	NameAttr   string `xml:"name,attr"`
	Comment    string `xml:",comment"`
	API        string `xml:"api,attr,omitempty"`
}

type Groups struct {
	Group []Group `xml:"group"`
}

type Group struct {
	Name string `xml:"name,attr"`
	Enum []Enum `xml:"enum"`
}

type Enum struct {
	Name    string `xml:"name,attr"`
	Value   string `xml:"value,attr,omitempty"`
	Comment string `xml:"comment,attr,omitempty"`
}

type Unused struct {
	Start   string `xml:"start,attr"`
	End     string `xml:"end,attr"`
	Comment string `xml:"comment,attr"`
}

type Enums struct {
	Namespace string   `xml:"namespace,attr"`
	Group     string   `xml:"group,attr"`
	Type      string   `xml:"type,attr"`
	Comment   string   `xml:"comment,attr"`
	Start     string   `xml:"start,attr"`
	End       string   `xml:"end,attr"`
	Vendor    string   `xml:"vendor,attr"`
	Enum      []Enum   `xml:"enum"`
	Unused    []Unused `xml:"unused"`
}

type Commands struct {
	Namespace string    `xml:"namespace,attr"`
	Command   []Command `xml:"command"`
	GLX       []GLX     `xml:"glx"`
}

type Command struct {
	Name  string  `xml:"name,attr,omitempty"`
	Proto Proto   `xml:"proto"`
	Param []Param `xml:"param"`
}

type Proto struct {
	Raw        string    `xml:",innerxml"`
	Definition string    `xml:",chardata"`
	Name       Name      `xml:"name,omitempty"`
	Returns    ParamType `xml:"ptype"`
}

type Name struct {
	String string `xml:",chardata"`
}

type Param struct {
	Raw        string    `xml:",innerxml"`
	Definition string    `xml:",chardata"`
	Group      string    `xml:"group,attr"`
	Len        string    `xml:"len,attr"`
	Name       Name      `xml:"name"`
	ParamType  ParamType `xml:"ptype"`
}

type ParamType struct {
	String string `xml:",chardata"`
}

type GLX struct {
	Type   string `xml:"render,attr"`
	Opcode string `xml:"opcode,attr"`
}

type Feature struct {
	API     string    `xml:"api,attr"`
	Name    string    `xml:"name,attr"`
	Number  string    `xml:"number,attr"`
	Require []Require `xml:"require"`
	Remove  []Remove  `xml:"remove"`
}

type Require struct {
	Comment string    `xml:"comment,attr"`
	Profile string    `xml:"profile,attr"`
	Type    []Type    `xml:"type"`
	Command []Command `xml:"command"`
	Enum    []Enum    `xml:"enum"`
}

type Remove struct {
	Comment string    `xml:"comment,attr"`
	Profile string    `xml:"profile,attr"`
	Command []Command `xml:"command"`
	Enum    []Enum    `xml:"enum"`
}

type Extensions struct {
	Extension []Extension `xml:"extension"`
}

type Extension struct {
	Name      string    `xml:"name,attr"`
	Supported string    `xml:"supported,attr"`
	Require   []Require `xml:"require"`
}
