package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/gofrs/uuid"
)

type Foo struct {
	count int
}

func (f *Foo) add(n int) {
	f.count = f.count + n
}

func main() {
	var f Foo
	fmt.Println(f)
	f.add(10)
	fmt.Println(f)
	// start := time.Now()
	// time.Sleep(time.Second)
	// end := time.Now()

	// diff := end - start
}

type aservice struct {
	logBuf []auditmodel.Audit
	mu     sync.Mutex
}

const drainLogInterval = 10 * time.Second

func (s *aservice) run() {
	ticker := time.NewTicker(drainLogInterval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			s.mu.Lock()
			logEntries := s.logBuf
			s.logBuf = nil
			s.mu.Unlock()

			s.dao.NewAutoDistributionQuery(s.ctx).InsertAudits(logEntries)
		case <-s.ctx.Done():
			close(s.logChannel)
			return
		}
	}
}

// Audit - метод для записи логов. Использует канал для отправки логов.
func (s *aservice) Audit(taskID uuid.UUID, reserveID *uuid.UUID, actor string, reason string, status string) {
	s.mu.Lock()
	s.logBuf = append(s.logBuf, auditmodel.Audit{
		TaskID:    taskID,
		ReserveID: reserveID,
		Event: auditmodel.Event{
			Reason:    reason,
			Actor:     actor,
			Status:    status,
			CreatedAt: time.Now(),
		},
	})
	s.mu.Unlock()
}
