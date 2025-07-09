package models

import (
	"database/sql"
	"fmt"

	"github.com/youssefrag/SoundSphere/db"
)

type Song struct {
	ID int64           `json:"id"`
	Name string        `json:"name"        binding:"required"`
	ArtistEmail string `json:"artistEmail,omitempty" binding:"required"`
	Genre string       `json:"genre"       binding:"required"`
	Duration int64     `json:"duration"       binding:"required"`
	SongUrl string     `json:"songUrl"     binding:"required"`
}

type ArtistWithSongs struct {
  Name      string `json:"name"`
  Email     string `json:"email"`
  ImageUrl  string `json:"imageUrl"`
  Songs     []Song `json:"songs"`
}

func (s *Song) Save() error {

	fmt.Println("🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣🟣")
	// fmt.Println(err)
	
	const query = `
			INSERT INTO songs (name, genre, artist_id, duration, song_url)
			VALUES (
				$1, 
				$2, 
				(SELECT id FROM users WHERE email = $3), 
				$4,
				$5
			)
			RETURNING id
	`

	// QueryRow returns the row with the new ID
	err := db.DB.QueryRow(
			query,
			s.Name,
			s.Genre,
			s.ArtistEmail,
			s.Duration,
			s.SongUrl,
	).Scan(&s.ID)
	if err != nil {
			return fmt.Errorf("saving song: %w", err)
	}

	return nil
}

func GetAllSongs() ([]ArtistWithSongs, error) {

	fmt.Println("🔵 Entering GetAllSongs()")
	
	const query = `
		SELECT
			u.name,
			u.email,
			u.imageurl,
			s.id,
			s.name,
			s.genre,
			s.duration,
			s.song_url
		FROM songs s
		JOIN users u ON s.artist_id = u.id
		ORDER BY u.id, s.uploaded_at DESC
	`

	fmt.Println("🔵 About to run query…")

	rows, err := db.DB.Query(query)

	if err != nil {
		return nil, fmt.Errorf("query grouped songs: %w", err)
  }
	
	
	defer rows.Close()
	
	groups := make(map[string]*ArtistWithSongs)
	order  := []string{}
	
	

	for rows.Next() {
		var (
			artistName string
			artistEmail string
			artistImageURLRaw sql.NullString
			song Song
		)

		

		if err := rows.Scan(
			&artistName,
			&artistEmail,
			&artistImageURLRaw,
			&song.ID,
			&song.Name,
			&song.Genre,
			&song.Duration,
			&song.SongUrl,
		); err != nil {
			fmt.Println("🔴 scan grouped song row error:", err)
			return nil, fmt.Errorf("scan grouped song row: %w", err)
		}

		artistImageUrl := ""
		if artistImageURLRaw.Valid {
  		artistImageUrl = artistImageURLRaw.String
		}

		
		if _, exists := groups[artistEmail]; !exists {
			
			groups[artistEmail] = &ArtistWithSongs{
				Name: artistName,
				Email: artistEmail,
				ImageUrl: artistImageUrl,
				Songs: []Song{},
			}

			order = append(order, artistEmail)
			fmt.Println("🟡 New artist:", artistEmail, "order so far:", order)
		}
		
		groups[artistEmail].Songs = append(groups[artistEmail].Songs, song)
		fmt.Println("🟡 Appended song for", artistEmail, "total songs:", len(groups[artistEmail].Songs))
	}

	if iterErr := rows.Err(); iterErr != nil {
		fmt.Println("⚠️ rows.Err() was non-nil:", iterErr)
		return nil, fmt.Errorf("rows iteration error: %w", iterErr)
}

	fmt.Println("🔵 Finished loop. Artists grouped:", len(order))

  result := make([]ArtistWithSongs, 0, len(order))
  for _, email := range order {
    result = append(result, *groups[email])
  }

	fmt.Println("🔵 Returning", len(result), "ArtistWithSongs structs")
  return result, nil
}

