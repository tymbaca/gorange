package main

import (
	"fmt"
	"time"

	"github.com/tymbaca/gorange/cmd/gorm/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/logger"
)

func main() {
	db, err := New("host=localhost port=5433 user=user password=password dbname=feed-service-db sslmode=disable")
	if err != nil {
		panic(err)
	}

	// SaveLast(db, []entity.Stock{
	// 	{
	// 		ID:          39916817,
	// 		OfferID:     "12414",
	// 		AccountName: "DASFA",
	// 		LastStock:   10,
	// 		// NewStock:    0,
	// 		LastUpdated: time.Now(),
	// 		// NewUpdated:  time.Time{},
	// 	},
	// })
	err = SaveNew(db, []entity.Stock{
		{
			// ID:          39916817,
			OfferID:     "12414",
			AccountName: "DASFA",
			LastStock:   10,
			NewStock:    20,
			// LastUpdated: time.Now(),
			NewUpdated: time.Now(),
		},
		{
			// ID:          39916817,
			OfferID:     "12415",
			AccountName: "DASFA",
			LastStock:   10,
			NewStock:    25,
			// LastUpdated: time.Now(),
			NewUpdated: time.Now(),
		},
	})
	if err != nil {
		panic(err)
	}
}

func SaveLast(db *gorm.DB, stocks []entity.Stock) error {
	sql := db.ToSQL(func(tx *gorm.DB) *gorm.DB {
		return db.Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "offer_id"}, {Name: "account_name"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"last_stock",
				"last_updated",
			}),
		}).Create(&stocks)
	})

	fmt.Println(sql)
	return nil
}

func SaveNew(db *gorm.DB, stocks []entity.Stock) error {
	return db.Clauses(
		clause.OnConflict{
			Columns: []clause.Column{{Name: "offer_id"}, {Name: "account_name"}},
			DoUpdates: clause.AssignmentColumns([]string{
				"new_stock",
				"new_updated",
			}),
		}).Create(&stocks).Error
}

const (
	_defaultMaxIdleConns = 2
	_defaultMaxOpenConns = 0
	_defaultLogMode      = logger.Info
)

type config struct {
	translateError bool
	maxIdleConns   int
	maxOpenConns   int
	logMode        logger.LogLevel
	nowFunc        func() time.Time
}

func New(connString string, opts ...Option) (*gorm.DB, error) {
	defaultNowFunc := func() time.Time { return time.Now() }

	cfg := &config{
		maxIdleConns: _defaultMaxIdleConns,
		maxOpenConns: _defaultMaxOpenConns,
		logMode:      _defaultLogMode,
		nowFunc:      defaultNowFunc,
	}

	for _, opt := range opts {
		opt(cfg)
	}

	db, err := gorm.Open(postgres.Open(connString), cfg.toGormConfig(logger.Error)) // todo вынести в config logLevel
	if err != nil {
		return nil, fmt.Errorf("failed to init db session: %w", err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve sqlDB object: %w", err)
	}

	sqlDB.SetMaxIdleConns(cfg.maxIdleConns)
	sqlDB.SetMaxOpenConns(cfg.maxOpenConns)

	if err := sqlDB.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping db: %w", err)
	}

	return db, nil
}

type Option func(*config)

func NowFunc(f func() time.Time) Option {
	return func(c *config) {
		c.nowFunc = f
	}
}

func (c *config) toGormConfig(logLevel logger.LogLevel) *gorm.Config {
	return &gorm.Config{
		NowFunc:        c.nowFunc,
		TranslateError: c.translateError,
		Logger:         logger.Default.LogMode(logLevel),
	}
}

func SilentLogger() Option {
	return func(c *config) {
		c.logMode = logger.Silent
	}
}
