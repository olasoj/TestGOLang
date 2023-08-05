package data

import (
	"time"
)

type Movie struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Year      int32     `json:"year,omitempty"`
	// Use the Runtime type instead of int32. Note that the omitempty directive will
	// still work on this: if the Runtime field has the underlying value 0, then it will // be considered empty and omitted -- and the MarshalJSON() method we just made
	// won't be called at all.
	Runtime Runtime  `json:"runtime,omitempty"`
	Genres  []string `json:"genres,omitempty"`
	Version int32    `json:"version"`

	// Unique integer ID for the movie
	// Timestamp for when the movie is added to our database
	// Movie title
	// Movie release year
	// Movie runtime (in minutes)
	// Slice of genres for the movie (romance, comedy, etc.)
	// The version number starts at 1 and will be incremented each // time the movie information is updated
}
