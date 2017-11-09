package testexample

func DivMod(dvdn, dvsr int) (q, r int) {
	r = dvdn
	for r >= dvsr {
		q += 1
		r = r -dvsr
	}
	return
}