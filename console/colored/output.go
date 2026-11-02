package colored

import "io"

// output returns w, or a writer discarding everything when w is nil.
func output(w io.Writer) io.Writer {
	if w == nil {
		return io.Discard
	}

	return w
}
