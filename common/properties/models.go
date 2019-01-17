package properties

import "time"

type Location struct {
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Address   string  `json:"address"`
}

type Properties struct {
	ID           string    `json:"_id"`
	UserID       string    `json:"user_id"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	DateCreated  time.Time `json:"created_at"`
	DateUpdated  time.Time `json:"updated_at"`
	Price        float64   `json:"price"`
	Location     Location  `json:"location"`
	ToiletCount  int       `json:"toilet_count"`
	RoomCount    int       `json:"room_count"`
	PropertyType string    `json:"property_type"`
	Amenities    []string  `json:"amenities"`
}
