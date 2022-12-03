package db

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/jaswdr/faker"
	_ "github.com/lib/pq"
	"github.com/ramonamaltan/go-api/internal/models"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "pguser"
	password = "localtest"
	dbname   = "aroundhome"
)

func Init() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	var err error
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}
	fmt.Println("Successfully connected!")

	_, err = insertDummyData(db)
	if err != nil {
		log.Fatal("failed to insert dummy data")
	}
	return db
}

func insertDummyData(db *sql.DB) ([]models.Partner, error) {
	clearDB := `delete from "partners"`
	_, err := db.Exec(clearDB)
	if err != nil {
		return nil, err
	}

	rand.Seed(time.Now().UnixNano())
	fake := faker.New()
	materials := []string{"wood", "tiles", "carpet", "wood, tiles", "wood, tiles, carpet", "carpet, wood"}
	services := []string{"flooring", "other"}
	queries := models.New(db)
	var partners []models.Partner
	for i := 0; i < 100; i++ {
		randomRating := rand.Intn(5)
		randomRadius := rand.Intn(100)
		randMatI := rand.Intn(len(materials))
		material := materials[randMatI]
		randSerI := rand.Intn(len(services))
		service := services[randSerI]
		lat := rand.Float64() + 52
		long := rand.Float64() + 13
		partner, err2 := queries.CreatePartner(context.Background(), models.CreatePartnerParams{
			Partnername: fake.Person().Name(),
			Servicename: service,
			Latitude:    lat,
			Longitude:   long,
			Material:    material,
			Radius:      int32(randomRadius),
			Rating:      float64(randomRating),
		})
		if err2 != nil {
			return nil, err2
		}
		partners = append(partners, partner)
	}

	return partners, nil
}
