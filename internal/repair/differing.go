//ff:func feature=repair type=helper control=iteration dimension=1
//ff:what differing — the checks whose stored rows differ from their occurrence (any verdict but same)
package repair

// differing returns the checks whose verdict is not Same: their stored rows
// are to be replaced.
func differing(checks []HashCheck) []HashCheck {
	var out []HashCheck
	for _, h := range checks {
		if h.Verdict != Same {
			out = append(out, h)
		}
	}
	return out
}
