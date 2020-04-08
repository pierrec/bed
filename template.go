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
func templateExec(w io.Writer, tmpl string, data map[string]string) error {
	const sep = "%"
	fragments := strings.Split(tmpl, sep)
	if len(fragments)%2 == 0 {
		return fmt.Errorf("imbalanced separator %q", sep)
	}
	var pos int // fragment position in the template
	for i, f := range fragments {
		if i%2 == 0 {
			// Text.
			if _, err := io.WriteString(w, f); err != nil {
				return err
			}
			pos += len(f) + 1
			continue
		}
		// Variable.
		ppos := pos
		pos += len(f) + 1
		if v, ok := data[f]; ok {
			// Variable defined.
			f = v
		} else {
			// Variable not found, keep it as is.
			f = tmpl[ppos-1 : pos]
		}
		if _, err := io.WriteString(w, f); err != nil {
			return err
		}
	}
	return nil
}
