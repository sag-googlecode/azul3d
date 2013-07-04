// Copyright 2012 Lightpoke. All rights reserved.
// This source code is subject to the terms and
// conditions defined in the "License.txt" file.

package config

import (
	"bytes"
	"io"
	"runtime"
	"strconv"
)

// Write writes this *Config out, it also writes out comments and the same number of line returns
// if this *Config was retrieved via config.Parse()
func (c *Config) Write(w io.Writer) error {
	c.access.RLock()
	defer c.access.RUnlock()

	buf := new(bytes.Buffer)

	if len(c.lines) > 0 {
		// Lines contain information like comments and newlines, we should opt for these
		for _, ln := range c.lines {
			if ln.comment {
				buf.WriteString(ln.value.(string))
				buf.WriteString("\n")
			} else if ln.newline {
				buf.WriteString("\n")
			} else {
				if !ln.comment && !ln.newline {
					buf.WriteString(ln.key)
					buf.WriteString(" = ")

					switch ln.value.(type) {
					case string:
						buf.WriteString(ln.value.(string))

					case bool:
						if ln.value.(bool) {
							buf.WriteString("true")
						} else {
							buf.WriteString("false")
						}

					case float64:
						buf.WriteString(strconv.FormatFloat(ln.value.(float64), 'f', -1, 64))
					}
				}
				buf.WriteString("\n")
			}
		}
	} else {
		for key, value := range c.values {
			buf.WriteString(key)
			buf.WriteString(" = ")
			switch value.(type) {
			case string:
				buf.WriteString(value.(string))

			case bool:
				if value.(bool) {
					buf.WriteString("true")
				} else {
					buf.WriteString("false")
				}

			case float64:
				buf.WriteString(strconv.FormatFloat(value.(float64), 'f', -1, 64))
			}
			if runtime.GOOS == "windows" {
				buf.WriteString("\r\n")
			} else {
				buf.WriteString("\n")
			}
		}
	}

	_, err := buf.WriteTo(w)
	if err != nil {
		return err
	}
	return nil
}
