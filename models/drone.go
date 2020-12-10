package models

import (
	"time"

	"github.com/Peace-for-ever/peace-api/eventpb"
	"github.com/brianvoe/gofakeit"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GenerateEvent() *eventpb.Event {
	d := eventpb.Event{}
	d.Message = gofakeit.Sentence(10)
	city := gofakeit.Number(0, 10595)
	d.Latitude = Cities[city].Latitude
	d.Longitude = Cities[city].Longitude
	d.Country = Cities[city].Country
	d.Citizen = gofakeit.Person().FirstName + " " + gofakeit.Person().LastName
	d.Date = timestamppb.New(gofakeit.DateRange(time.Now().AddDate(0, 0, -7), time.Now()))
	d.Battery = int32(gofakeit.Number(0, 100))
	d.Temperature = int32(gofakeit.Number(0, 40))

	return &d
}
