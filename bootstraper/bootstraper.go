package bootstraper

import (
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"marketdataservice/model"
)

type Bootstrapper struct {
	PgConn       *gorm.DB
	Logger       logx.Logger
	CurrencyPair model.CurrencyPair
}

func NewBootstrapper(dbConfig Config) (*Bootstrapper, error) {
	var logger logx.Logger
	conn, err := Connect(dbConfig)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to the database: %w", err)
	}

	CurrencyPair := model.NewCurrencyPairModel(conn)
	return &Bootstrapper{PgConn: conn, CurrencyPair: CurrencyPair,
		Logger: logger,
	}, nil
}

type Config struct {
	Host     string
	Password string
	Port     int
	Name     string
	SSL      string
	User     string
}

func Connect(c Config) (*gorm.DB, error) {
	dsn := fmt.Sprintf("user=%v password=%v host=%v port=%v dbname=%v sslmode=%v", c.User, c.Password, c.Host, c.Port, c.Name, c.SSL)
	// Connect to the PostgreSQL database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("unable to bootstrap gorm: %w", err)
	}

	err = migrate(db)
	if err != nil {
		return nil, fmt.Errorf("unable to bootstrap gorm: %w", err)
	}

	return db, err
}

func migrate(conn *gorm.DB) error {
	err := conn.AutoMigrate(&model.CurrencyPairManagement{})
	if err != nil {
		return fmt.Errorf("unable to auto-migrate: %w", err)
	}
	return nil
}
