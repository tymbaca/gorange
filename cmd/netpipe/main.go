package main

import (
	"errors"
	"fmt"
	"io"
	"net"
	"sync"
	"time"

	"github.com/charmbracelet/log"
)

type CalculateResponse struct {
	Result int
	Err    error
}

type calculator interface {
	CalculateSomehow(val int) (int, error)
}

type client struct {
	calculator calculator
}

// BatchCalculate вызывает метод [CalculateSomehow] у [calculator] для каждого значения
// В случае ошибки у какого-либо значения - положит эту ошибку в отдельный объект в возвращаемом массиве
// а не вернет общую ошибку.
func (c *client) BatchCalculate(vals []int) []CalculateResponse {
	responses := make([]CalculateResponse, 0, len(vals))

	for _, val := range vals {
		result, err := c.calculator.CalculateSomehow(val)
		if err != nil {
			responses = append(responses, CalculateResponse{Err: err})
			continue
		}

		responses = append(responses, CalculateResponse{Result: result})
	}

	return responses
}

func main() {
	// log.SetLevel(log.InfoLevel)
	w, r := net.Pipe()
	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer wg.Done()

		<-time.After(1 * time.Second)

		defer log.Info("exitings reader", "actor", "reader")

		buf := make([]byte, 3)
		for {
			n, err := r.Read(buf)
			if err != nil {
				if errors.Is(err, io.EOF) {
					return
				}
				log.Error(err.Error(), "actor", "reader")
				return
			}

			log.Info(fmt.Sprintf("got: '%s'", buf[:n]), "actor", "reader")
		}
	}()
	go func() {
		defer wg.Done()
		var err error
		var n int

		<-time.After(1 * time.Second)

		log.Info("starting to write 1st part", "actor", "writer")
		n, err = w.Write([]byte("hello "))
		if err != nil {
			log.Error(err.Error(), "actor", "writer")
			return
		}
		log.Info("wrote 1st part, n = "+fmt.Sprint(n), "actor", "writer")

		<-time.After(2 * time.Second)

		log.Info("starting to write 2nd part", "actor", "writer")
		n, err = w.Write([]byte("world"))
		if err != nil {
			log.Error(err.Error(), "actor", "writer")
			return
		}
		log.Info("wrote 2nd part, n = "+fmt.Sprint(n), "actor", "writer")
	}()
	go func() {
		defer wg.Done()

		<-time.After(2 * time.Second)
		r.Close()
		log.Info("closed conn", "actor", "closed")
	}()

	wg.Wait()
}
