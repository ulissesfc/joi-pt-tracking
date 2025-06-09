package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/ulissesfc/joi-pt-tracking.git/internal/db"
)

func insert(dbcon *sql.DB, TripID string) {
	err := dbcon.Ping()
	if err != nil {
		panic(err)
	}

	queries := db.New(dbcon)

	ctx := context.Background()
	for i := range 200 {
		data, err := queries.InsertDataTrip(ctx, db.InsertDataTripParams{
			VehiclePrefix:  "2007",
			Latitude:       -26.310283,
			Longitude:      -48.860653,
			ShapeID:        "0041-0",
			TripID:         TripID,
			TripStatus:     "LIVE",
			TripProgress:   int16(i),
			TripDelay:      -126,
			StopDistance:   -257.49,
			StopOrder:      9,
			ReportTimeDiff: 7,
		})

		if err != nil {
			panic(err)
		}

		fmt.Println(data)

	}
}

func main() {
	dbcon, err := sql.Open("postgres", "user=admin password=admin123 dbname=trips host=localhost port=5433 sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer dbcon.Close()

	go insert(dbcon, "T-DB-CON")
	insert(dbcon, "T-DB-CON2")

	time.Sleep(time.Second * 15)

}
