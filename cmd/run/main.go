package main

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/charmbracelet/log"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/tymbaca/gorange/internal/config"
	"github.com/tymbaca/gorange/internal/helper/mem"
)

func main() {
	cfg := config.New()

	log.SetLevel(log.DebugLevel)

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	pool, err := pgxpool.New(ctx, fmt.Sprintf("postgres://%s:%s@%s:%s/%s?pool_max_conns=%d", //nolint
		cfg.PostgresUser,
		cfg.PostgresPassword,
		cfg.PostgresHost,
		cfg.PostgresPort,
		cfg.PostgresDB,
		cfg.PostgresMaxConn))
	if err != nil {
		log.Fatalf("err: %s", err)
	}
	defer pool.Close()

	//--------------------------------------------------------------------------------------------------

	type killer struct {
		Name      string    `db:"name"`
		Kills     uint      `db:"kills"`
		City      string    `db:"city"`
		CreatedAt time.Time `db:"created_at"`
	}

	log.Info("starting killer generation", "mem", mem.FormatMem(mem.MiB))

	count := 5_000_000
	killers := make([]killer, 0, count)
	for range count {
		killers = append(killers, killer{
			Name:      gofakeit.Name(),
			Kills:     gofakeit.UintN(1000),
			City:      gofakeit.City(),
			CreatedAt: gofakeit.PastDate(),
		})
	}

	log.Infof("generated all killers, gonna copy to repo, mem: %s", mem.FormatMem(mem.MiB))

	_, err = pool.CopyFrom(
		ctx,
		pgx.Identifier{"killers"},
		[]string{"name", "kills", "city", "created_at"},
		pgx.CopyFromSlice(len(killers), func(i int) ([]any, error) {
			return []any{killers[i].Name, killers[i].Kills, killers[i].City, killers[i].CreatedAt}, nil
		}),
	)
	if err != nil {
		log.Fatal(err)
	}

	log.Infof("successfully copied to repo, mem: %s", mem.FormatMem(mem.MiB))

	runtime.GC()
	log.Infof("after GC, mem: %s", mem.FormatMem(mem.MiB))
}
