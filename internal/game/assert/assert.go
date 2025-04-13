package assert

import "log"

func True(cond bool, fmt string, args ...any) {
	if !cond {
		log.Panicf(fmt, args...)
	}
}

func Panic(fmt string, args ...any) {
	log.Panicf(fmt, args...)
}
