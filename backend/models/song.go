package models

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"time"

	"github.com/youssefrag/SoundSphere/cache"
	"github.com/youssefrag/SoundSphere/db"
)

type Song struct {
	ID int64           `json:"id"`
	Name string        `json:"name"                   binding:"required"`
	ArtistEmail string `json:"artistEmail,omitempty"  binding:"required"`
	Genre string       `json:"genre"                  binding:"required"`
	Duration int64     `json:"duration"               binding:"required"`
	SongUrl string     `json:"songUrl"                binding:"required"`
	StoragePath string `json:"storagePath,,omitempty" binding:required`

}

type SongDetails struct {
	Name string         `json:"name"       binding:"required"`
	Genre string        `json:"genre"      binding:"required"`
	SongUrl string      `json:"songUrl"    binding:"required"`
	Date time.Time      `json:"date"         binding:"required"`
	ArtistName string   `json:"artistName"   binding:"required"`
	ArtistImgUrl string `json:"artistImgUrl" binding:"required"`
}

type ArtistWithSongs struct {
  Name      string `json:"name"`
  Email     string `json:"email"`
  ImageUrl  string `json:"imageUrl"`
  Songs     []Song `json:"songs"`
}

// Redis cache keys & TTL for songs
const (
  allSongsKey    = "songs:all"
  cacheTTL        = 5 * time.Minute
)

func (s *Song) Save() error {
	
	const query = `
			INSERT INTO songs (name, genre, artist_id, duration, song_url, storage_path)
			VALUES (
				$1, 
				$2, 
				(SELECT id FROM users WHERE email = $3), 
				$4,
				$5,
				$6
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
			s.StoragePath,
	).Scan(&s.ID)
	if err != nil {
			return fmt.Errorf("saving song: %w", err)
	}

	return nil
}

func GetAllSongs() ([]ArtistWithSongs, error) {

	ctx := cache.Ctx

	// Try Redic cache

	if blob, err := cache.Client.Get(ctx, allSongsKey).Result(); err == nil {
		var list []ArtistWithSongs
		if err := json.Unmarshal([]byte(blob), &list); err == nil {
			fmt.Println("✅ Cache hit for all songs")
			return list, nil
		}
	}

	// Cache Miss
	fmt.Println("🔵 Cache miss — loading all songs from DB")
	
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
		}
		
		groups[artistEmail].Songs = append(groups[artistEmail].Songs, song)
	}

	if iterErr := rows.Err(); iterErr != nil {
		return nil, fmt.Errorf("rows iteration error: %w", iterErr)
}

  result := make([]ArtistWithSongs, 0, len(order))
  for _, email := range order {
    result = append(result, *groups[email])
  }

	// Save to cache
	if blob, err := json.Marshal(result); err == nil {
		cache.Client.Set(ctx, allSongsKey, blob, cacheTTL)
	}

  return result, nil
}

func DeleteSong(songId int64)(string, error) {

	const query = `
		DELETE FROM songs
		WHERE id = $1
		RETURNING storage_path;
	`

	var storagePath string

	err := db.DB.QueryRow(query, songId).Scan(&storagePath)

	if err != nil {
		return "", fmt.Errorf("could not delete song %d: %w", songId, err)
	}

	return storagePath, nil
}

func EditSong(songId int64, name string, genre string) error {

	const query = `
		UPDATE songs
			SET name = $1,
					genre = $2
		WHERE id = $3
	`
	_, err := db.DB.Exec(query, name, genre, songId)
  if err != nil {
    return fmt.Errorf("could not update song %d: %w", songId, err)
  }

	return nil

}

func GetSongDetails(songId int64) (SongDetails, error) {
	var details SongDetails
	query := `
  	SELECT
    	s.name,
    	s.genre,
    	s.song_url    AS "songUrl",
    	s.uploaded_at AS "date",
    	u.name        AS "artistName",
			COALESCE(
  			NULLIF(TRIM(u.imageurl), ''),
  			'https://images.unsplash.com/photo-1511671782779-c97d3d27a1d4?q=80&w=2070&auto=format&fit=crop&ixlib=rb-4.1.0&ixid=M3wxMjA3fDB8MHxwaG90by1wYWdlfHx8fGVufDB8fHx8fA%3D%3D'
			) AS "artistImgUrl"
  		FROM songs  AS s
  		JOIN users  AS u ON s.artist_id = u.id
  		WHERE s.id = $1;
	`



	if err := db.DB.QueryRow(query, songId).Scan(
    &details.Name,
    &details.Genre,
    &details.SongUrl,
    &details.Date,
    &details.ArtistName,
    &details.ArtistImgUrl,
	); err != nil {

		fmt.Println("🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑🛑", err)
    return SongDetails{}, err
	}

	return details, nil
}