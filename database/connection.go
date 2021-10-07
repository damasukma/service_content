package database

import (
	"context"
	"log"
	"time"

	"gorm.io/gorm"

	"github.com/service_content/model"
	"gorm.io/driver/mysql"
)

type (
	connection struct {
		DB *gorm.DB
		// ES *elastic.Client
	}

	EventStore struct {
		ID            string
		EventType     string
		AggregateID   string
		AggregateType string
		EventData     string
		Channel       string
	}

	EventStoreRepository interface {
		CreateEvent(ctx context.Context, params *model.EventParam) error
	}

	Article struct {
		ID        int32     `json:"id"`
		Author    string    `json:"author"`
		Title     string    `json:"title"`
		Body      string    `json:"body"`
		CreatedAt time.Time `json:"crated_at"`
	}
)

func NewConnection() *gorm.DB {

	dsn := "root:@tcp(127.0.0.1:3306)/service_article?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	// if errAr := db.Debug().AutoMigrate(&Article{}); errAr != nil {
	// 	log.Fatal(errAr)
	// }

	// if errEv := db.Debug().AutoMigrate(&EventStore{}); errEv != nil {
	// 	log.Fatal(errEv)
	// }

	return db
}

func NewEventConnection(db *gorm.DB) EventStoreRepository {
	return &connection{DB: db}
}

func (c *connection) CreateEvent(ctx context.Context, params *model.EventParam) error {
	log.Print(params)
	// query := c.DB.WithContext(ctx)

	// if err := query.Create(&EventStore{
	// 	ID:            params.EventId,
	// 	EventType:     params.EventType,
	// 	AggregateID:   params.AggregateId,
	// 	AggregateType: params.AggregateType,
	// 	EventData:     params.EventData,
	// 	Channel:       params.Channel,
	// }).Error; err != nil {
	// 	return err
	// }

	return nil
}
