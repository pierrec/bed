package serializer

import (
	"fmt"
	"io"
	"strings"
)

// templateExec writes the template tmpl instantiated with the variables defined in data:
//   - variables are enclosed in the % character
//   - variables not defined in data are returned as is, allowing nested templates
//   - use %% for the % character
//   - using fmt words returns an error
func templateExec(w io.Writer, tmpl string, data map[string]interface{}) error {
	fragments := strings.Split(tmpl, "%")
	s := fragments[0]
	var args []interface{}
	for i, f := range fragments[1:] {
		switch f {
		case "":
			continue
		case "t", "b", "c", "d", "o", "O", "q", "x", "X", "U", "e", "E", "f", "F", "g", "G", "s", "p", "T":
			return fmt.Errorf("cannot use fmt verb: %s", f)
		}
		if v, ok := data[f]; ok {
			s += "%v"
			args = append(args, v)
			continue
		}
		if i%2 == 0 {
			f = "%%" + f + "%%"
		}
		s += f
	}
	_, err := fmt.Fprintf(w, s, args...)
	return err
}
