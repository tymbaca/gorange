package main

import (
	"errors"
	"fmt"
	"io"
	"runtime/debug"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/tymbaca/gorange/cmd/io/encodereader/xml"
	"github.com/tymbaca/gorange/internal/helper/mem"
)

type RecordBook struct {
	Records []Record
}

type Record struct {
	ID   string
	Key  string
	Val  string
	Bulk [32]byte
}

func main() {
	debug.SetGCPercent(-1)

	const N = 100_000
	records := make([]Record, 0, N)
	for range 100_000 {
		records = append(records, Record{
			ID:  gofakeit.UUID(),
			Key: gofakeit.Name(),
			Val: gofakeit.Gender(),
		})
	}

	rb := RecordBook{
		Records: records,
	}

	fmt.Printf("just structs: %v\n", mem.Format(mem.MiB))

	// data, _ := xml.MarshalIndent(rb, "", "  ")
	// copyInterval(io.Discard, bytes.NewBuffer(data), 50*time.Millisecond)

	r := xml.NewEncoder(rb).Indent("", "  ").Encode()
	defer r.Close()
	copyInterval(io.Discard, r, 50*time.Millisecond)

	// marshal(rb)
	// encode(rb)

	// marshalJson(rb)
	// r := encodeJsonReader(rb)
	// copyInterval(io.Discard, r, 300*time.Millisecond)
}

// func encodeJsonReader(v any) io.ReadCloser {
// 	r, w := io.Pipe()
//
// 	go func() {
// 		e := json.NewEncoder(w)
// 		e.SetIndent("", "  ")
// 		err := e.Encode(v)
// 		if err != nil {
// 			_ = w.CloseWithError(err)
// 		}
// 		w.Close()
// 	}()
//
// 	return r
// }
//
// func encodeXmlReader(v any) io.ReadCloser {
// 	r, w := io.Pipe()
//
// 	go func() {
// 		e := xml.NewEncoder(w)
// 		e.Indent("", "	")
// 		err := e.Encode(v)
// 		if err != nil {
// 			_ = w.CloseWithError(err)
// 			return
// 		}
// 		w.Close()
// 	}()
//
// 	return r
// }
//
// func marshalXml(rb RecordBook) {
// 	data, err := xml.Marshal(rb)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	fmt.Fprintf(io.Discard, "%s", data)
// 	fmt.Printf("\nmem marshal: %v\n", mem.Format(mem.MiB))
// }
//
// func marshalJson(rb RecordBook) {
// 	data, err := json.Marshal(rb)
// 	if err != nil {
// 		panic(err)
// 	}
//
// 	fmt.Fprintf(io.Discard, "%s", data)
// 	fmt.Printf("\nmem marshal: %v\n", mem.Format(mem.MiB))
// }
//
// func encode(rb RecordBook) {
// 	r, w := io.Pipe()
//
// 	go func() {
// 		copyInterval(io.Discard, r, 50*time.Millisecond)
// 	}()
//
// 	xml.NewEncoder(w).Encode(rb)
// }

func copyInterval(dst io.Writer, src io.Reader, interval time.Duration) {
	buf := make([]byte, 1024)
	for ; true; time.Sleep(interval) {
		n, err := src.Read(buf)
		if n > 0 {
			dst.Write(buf[:n])
			fmt.Printf("\nmem encode: %v", mem.Format(mem.MiB))
		}
		if errors.Is(err, io.EOF) {
			break
		}
		if err != nil {
			panic(err)
		}
	}
}
