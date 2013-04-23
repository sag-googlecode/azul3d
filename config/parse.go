package config

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

const (
	lf = '\n'
	cr = '\r'
)

func (l *line) String() string {
	if l.comment {
		return fmt.Sprintf("comment(%q)", l.value)
	} else if l.newline {
		return fmt.Sprintf("newline()")
	}
	return fmt.Sprintf("keyvalue(%q, %q)", l.key, l.value)
}

// Parse parses the data and returns an initiliazed *Config
func Parse(data []byte) *Config {
	c := new(Config)
	c.values = make(map[string]interface{})

	sdata := string(data)
	sdata = strings.Replace(sdata, "\r\n", "\n", -1)
	sdata = strings.Replace(sdata, "\r", "\n", -1)

	lineStrings := strings.Split(sdata, "\n")

	comment := new(bytes.Buffer)
	key := new(bytes.Buffer)
	value := new(bytes.Buffer)
	var readingValue bool

	for _, lineString := range lineStrings {
		ln := new(line)

		comment.Reset()
		key.Reset()
		value.Reset()
		readingValue = false

		for _, c := range lineString {
			if ln.comment {
				comment.WriteRune(c)
				continue
			}
			if c == '#' && !readingValue && key.Len() == 0 {
				ln.comment = true
				continue
			}

			if !ln.comment {
				// If not an comment, it must be either an dull, useless line, or an key=value pair
				if (c == ' ' || c == '\t') && !readingValue && key.Len() == 0 {
					// Drop all prefixed white space
				} else {
					if readingValue {
						value.WriteRune(c)
						continue
					} else {
						if c == '=' {
							noSpace := strings.TrimRight(string(key.Bytes()), " \t")
							key.Reset()
							key.WriteString(noSpace)
							//fmt.Printf("%q\n", string(key.Bytes()))
							readingValue = true
							continue
						}
						key.WriteRune(c)
					}
				}
			}
		}

		if ln.comment {
			// Comment
			ln.value = string(comment.Bytes())
		} else {
			// Either new line, or key=value pair
			keyBytes := key.Bytes()
			if len(keyBytes) == 0 {
				// New line
				ln.newline = true
			} else {
				// key=value pair
				ln.key = string(keyBytes)
				ln.value = string(value.Bytes())
				if len(ln.value.(string)) > 0 {
					ln.value = strings.TrimLeft(ln.value.(string), " \t")

					if ln.value.(string) == "true" {
						ln.value = true
					} else if ln.value.(string) == "false" {
						ln.value = false
					} else {
						fv, err := strconv.ParseFloat(ln.value.(string), 64)
						if err == nil {
							ln.value = fv
						}
					}
				} else {
					// Since it's invalid, simply replace it with an new line
					ln.newline = true
				}
			}
		}

		c.lines = append(c.lines, ln)
	}

	for _, ln := range c.lines {
		if !ln.comment && !ln.newline {
			c.values[ln.key] = ln.value
		}
	}
	return c
}
