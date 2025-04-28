package mem

import (
	"fmt"
	"runtime"
)

type Unit int

const (
	KiB Unit = 1024       // 1024 B
	MiB Unit = 1024 * KiB // 1024 KiB
)

func (u Unit) String() string {
	switch u {
	case KiB:
		return "KiB"
	case MiB:
		return "MiB"
	}

	return fmt.Sprintf("Unit(%d)", u)
}

// Heap gets bytes of allocated heap objects
func Heap() uint64 {
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	return mem.HeapAlloc
}

// Stack gets bytes of stack size
func Stack() uint64 {
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	return mem.StackSys
}

func Format(unit Unit) string {
	return fmt.Sprintf("stack: %.2f %s, heap: %.2f %s", float32(Stack())/float32(unit), unit, float32(Heap())/float32(unit), unit) //nolint
}
