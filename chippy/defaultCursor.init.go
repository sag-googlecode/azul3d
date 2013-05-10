package chippy

import(
	_ "image/png"
	"image"
	"bytes"
	"fmt"
)

var defaultCursor *Cursor

func init() {
	var err error

	defaultCursor = new(Cursor)
	buf := bytes.NewBuffer(defaultCursorBytes)
	defaultCursor.Image, _, err = image.Decode(buf)
	if err != nil {
		panic(fmt.Sprintf("Unable to decode default cursor", err))
	}
	defaultCursor.X = 10
	defaultCursor.Y = 6
}

