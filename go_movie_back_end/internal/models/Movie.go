package modles

import "time"
type Movie struct {
ID int `json:"id"`
Title string `json:"title"`
ReleaseDate time.Time `json:"release_date"`
RunTime int `json:"runtime"`
MPAARating string `json:"mpaa_rating"`
Description string `json:"description"`
Image string `json:"image"`
CreatedAt time.Time `json:"-"`
UpdatedField time.Time `json:"-"`//ignore this using"-""

}