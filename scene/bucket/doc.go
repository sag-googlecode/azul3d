// Copyright 2014 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

// Package bucket provides utilities for controlling the render order of nodes.
//
// Controlling the render order of nodes is an important process for various
// rendering effects (i.e. alpha-blending).
//
// At the most basic level nodes are collected into so-called buckets and by
// default nodes are collected into the Opaque bucket (which is state-sorted
// for efficiency). Each bucket full of nodes are sorted in a specific manner
// (e.g. based on distance to the camera for alpha-blending, based on render
// state, or something else entirely user implemented by filling out the
// BucketSorter interface).
//
// The pre-defined buckets in this package will fulfill most purposes, but
// users can define their own method of sorting a bucket of nodes for rendering
// by filling out the Sorter interface.
//
// Render order vs visual order
//
// The order that nodes are rendered in does not directly affect the order in
// which objects appear on the screen relative to one another (i.e. weather an
// object is behind or in front of another object) but instead contributes to
// it in addition to other factors such as depth testing.
//
// Misunderstanding the difference between render order and visual order is a
// common pitfall for beginners.
package bucket
