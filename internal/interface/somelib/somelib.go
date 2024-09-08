package somelib

type doer interface {
	do(val int) (int, error)
}

func UseDoer(doer doer, val int) (int, error) {
	return doer.do(val)
}
