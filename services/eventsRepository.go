package services

import (
	"fmt"

	models "github.com/evileric-com/events-svc/models"
	"github.com/go-redis/redis"
)

const BASE_PATH = "events/%s/%s"

type EventRepository struct {
	DB *redis.Client
}

type ConnectionSettings struct {
	Host     string
	Password string
}

func NewEventRepository(settings ConnectionSettings) EventRepository {
	db := redis.NewClient(&redis.Options{
		Addr:     settings.Host,
		Password: settings.Password,
		DB:       0,
	})

	return EventRepository{
		DB: db,
	}
}

func SaveEvent(repo EventRepository, event *models.Event) {
	fmt.Println("Creating Event", event.Id)

	key := fmt.Sprintf(BASE_PATH, event.Id, "id")
	fmt.Println("Id Key", key, event.Id)
	err := repo.DB.Set(key, event.Id, 0).Err()
	if err != nil {
		fmt.Println(40, err)
	}

	key = fmt.Sprintf(BASE_PATH, event.Id, "name")
	fmt.Println("Name Key", key, event.Name)
	err = repo.DB.Set(key, event.Name, 0).Err()
	if err != nil {
		fmt.Println(47, err)
	}

	key = fmt.Sprintf(BASE_PATH, event.Id, "date")
	fmt.Println("Date Key", key, event.Date)
	err = repo.DB.Set(key, event.Date, 0).Err()
	if err != nil {
		fmt.Println(54, err)
	}

	key = fmt.Sprintf(BASE_PATH, event.Id, "photos")
	fmt.Println("Photos Key", key, event.Photos)
	if len(event.Photos) != 0 {
		err = repo.DB.LSet(key, 0, event.Photos).Err()
		if err != nil {
			fmt.Println(62, err)
		}
	}
}

func GetEvent(repo EventRepository, id string) models.Event {
	eId, errId := repo.DB.Get(fmt.Sprintf(BASE_PATH, id, "id")).Result()
	name, errName := repo.DB.Get(fmt.Sprintf(BASE_PATH, id, "name")).Result()
	date, errDate := repo.DB.Get(fmt.Sprintf(BASE_PATH, id, "date")).Result()
	photos := repo.DB.LRange(fmt.Sprintf(BASE_PATH, id, "photos"), 0, 0).Val()

	fmt.Println("Event ID", eId, errId)
	fmt.Println("Name", name, errName)
	fmt.Println("Date", date, errDate)

	return models.Event{
		Id:     id,
		Name:   name,
		Date:   date,
		Photos: photos,
	}
}
