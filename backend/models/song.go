package models

import (
	"fmt"

	"github.com/youssefrag/SoundSphere/db"
)

type Song struct {
	ID int64           `json:"id"`
	Name string        `json:"name"        binding:"required"`
	ArtistEmail string `json:"artistEmail" binding:"required"`
	Genre string       `json:"genre"       binding:"required"`
	SongUrl string     `json:"songUrl"     binding:"required"`
}

func (s *Song) Save() error {

	fmt.Println("🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣")
	// fmt.Println(err)
	
	const query = `
			INSERT INTO songs (name, genre, artist_id, song_url)
			VALUES (
				$1, 
				$2, 
				(SELECT id FROM users WHERE email = $3), 
				$4
			)
			RETURNING id
	`

	// QueryRow returns the row with the new ID
	err := db.DB.QueryRow(
			query,
			s.Name,
			s.Genre,
			s.ArtistEmail,
			s.SongUrl,
	).Scan(&s.ID)
	if err != nil {
			return fmt.Errorf("saving song: %w", err)
	}

	return nil
}