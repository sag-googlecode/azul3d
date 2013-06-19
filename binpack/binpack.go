// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package binpack

import (
	"errors"
)

type Block struct {
	Width, Height uint
	X, Y          uint
}

type Grid struct {
	Width, Height uint
}

// Pack is an 2D bin packing algorithm based on the work of Jake Gordon, described on his blog at:
//
//   http://codeincomplete.com/posts/2011/5/7/bin_packing/
//
// And implemented by him in JavaScript here:
//
//   https://github.com/jakesgordon/bin-packing
//
// The returned grid will reflect how large of an grid is required to fit all blocks.
//
// The returned slice of blocks will have their X and Y coordinates updated to reflect where the
// block can be placed efficiently on the grid.
//
// The algorithm does not start with an fixed width and height, instead it starts with the width
// and height of the first block in the slice and then grows as neccessary to fit each block. As
// the grid is grown it attempts to maintain an roughly square ratio by making 'smart' choices
// about whether to grow right or down.
//
// When growing, the algorithm can only grow to the right OR down. Therefore, if the new block is
// BOTH wider and taller than the current target then it will be rejected. This makes it very
// important to initialize with a sensible starting width and height. If you are providing sorted
// input (largest first) then this will not be an issue.
//
// A potential way to solve this limitation would be to allow growth in BOTH directions at once,
// but this requires maintaining a more complex tree with 3 children (down, right and center) and
// that complexity can be avoided by simply chosing a sensible starting block.
//
// Best results occur when the input blocks are sorted by height, or even better when sorted by
// max(width,height).
//
// This function will panic if len(blocks) == 0, or if any of the *Block in the slice are nil.
func Pack(blocks []*Block) (*Block, error) {
	if len(blocks) == 0 {
		panic("Pack(): Cannot pack zero blocks!")
	}

	root := &node{
		x:      0,
		y:      0,
		width:  blocks[0].Width,
		height: blocks[0].Height,
	}

	for _, block := range blocks {
		if block == nil {
			panic("Pack(): Block is nil!")
		}

		node := root.find(block.Width, block.Height)
		if node != nil {
			node = node.split(block.Width, block.Height)

			// Update block in-place
			block.X = node.x
			block.Y = node.y

		} else {
			newRoot, grown := root.grow(block.Width, block.Height)
			if newRoot == nil {
				return nil, errors.New("Pack(): Could not grow node, first block size is not large enough.")
			}

			// Update block in-place
			block.X = grown.x
			block.Y = grown.y

			root = newRoot
		}
	}

	bounding := &Block{
		Width:  root.width,
		Height: root.height,
	}

	return bounding, nil
}

type node struct {
	x, y, width, height uint
	right, down         *node
}

func (n *node) find(width, height uint) *node {
	if n.right != nil || n.down != nil {
		right := n.right.find(width, height)
		if right != nil {
			return right
		}
		return n.down.find(width, height)
	} else if width <= n.width && height <= n.height {
		return n
	}
	return nil
}

func (n *node) split(width, height uint) *node {
	n.down = &node{
		x:      n.x,
		y:      n.y + height,
		width:  n.width,
		height: n.height - height,
	}

	n.right = &node{
		x:      n.x + width,
		y:      n.y,
		width:  n.width - width,
		height: height,
	}

	return n
}

func (n *node) grow(width, height uint) (root, grown *node) {
	canGrowDown := width <= n.width
	canGrowRight := height <= n.height

	// attempt to keep square-ish by growing right when height is much greater than width
	shouldGrowRight := canGrowRight && (n.height >= (n.width + width))

	// attempt to keep square-ish by growing down when width is much greater than height
	shouldGrowDown := canGrowDown && (n.width >= (n.height + height))

	if shouldGrowRight {
		return n.growRight(width, height)
	} else if shouldGrowDown {
		return n.growDown(width, height)
	} else if canGrowRight {
		return n.growRight(width, height)
	} else if canGrowDown {
		return n.growDown(width, height)
	}

	// need to ensure sensible root starting size to avoid this happening
	return nil, nil
}

func (n *node) growRight(width, height uint) (root, grown *node) {
	newRoot := &node{
		x:      0,
		y:      0,
		width:  n.width + width,
		height: n.height,
		down:   n,
		right: &node{
			x:      n.width,
			y:      0,
			width:  width,
			height: n.height,
		},
	}

	node := newRoot.find(width, height)
	if node != nil {
		return newRoot, node.split(width, height)
	}
	return nil, nil
}

func (n *node) growDown(width, height uint) (root, grown *node) {
	newRoot := &node{
		x:      0,
		y:      0,
		width:  n.width,
		height: n.height + height,
		down: &node{
			x:      0,
			y:      n.height,
			width:  n.width,
			height: height,
		},
		right: n,
	}

	node := newRoot.find(width, height)
	if node != nil {
		return newRoot, node.split(width, height)
	}
	return nil, nil
}
