package chippy

import(
	_ "image/png"
	"image"
	"bytes"
	"fmt"
)

var defaultIcon image.Image

func init() {
	var err error

	buf := bytes.NewBuffer(defaultIconBytes)
	defaultIcon, _, err = image.Decode(buf)
	if err != nil {
		panic(fmt.Sprintf("Unable to decode default icon", err))
	}
}

