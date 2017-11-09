package testexample

// DivMod performs a Eucledan division producing a quotient and remainder.
// This version only works if dividend and divisor > 0.
func DivMod(dvdn, dvsr int) (q, r int) {
	r = dvdn
	for r >= dvsr {
		q += 1
		r = r -dvsr
	}
	return
}