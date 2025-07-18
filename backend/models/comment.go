package models

import (
	"fmt"
	"time"

	"github.com/youssefrag/SoundSphere/db"
)

type Comment struct {
	ID      int64  `json:"id"`
	UserId  int64  `json:"userId"`
	SongId  int64  `json:"songId"`
	Content string `json:"content"`
}

type CommentWithAuthor struct {
	ID             int64     `json:"id"`
	Content        string    `json:"content"`
	CreatedAt      time.Time `json:"createdAt"`
	UserName       string    `json:"userName"`
	UserEmail      string    `json:"userEmail"`
	UserImageUrl   string    `json:"userImageUrl"`
}

func CreateComment(userID, songID int64, content string) (int64, error) {
	var id int64
	err := db.DB.QueryRow(`
			INSERT INTO comments (content, song_id, user_id)
			VALUES ($1, $2, $3)
			RETURNING id
	`, content, songID, userID).Scan(&id)
	if err != nil {
			return 0, err
	}
	return id, nil
}

func GetCommentsBySongID(songID int64) ([]CommentWithAuthor, error) {
	rows, err := db.DB.Query(`
			SELECT
					c.id,
					c.content,
					c.created_at,
					u.name AS user_name,
					u.email AS user_email,
					COALESCE(
						NULLIF(TRIM(u.imageurl), ''),
						'https://images.unsplash.com/photo-1511671782779-c97d3d27a1d4?q=80&w=2070'
					) AS user_image_url
			FROM comments AS c
			JOIN users    AS u ON c.user_id = u.id
			WHERE c.song_id = $1
			ORDER BY c.created_at DESC;
	`, songID)
	if err != nil {
			return nil, err
	}
	defer rows.Close()

	var out []CommentWithAuthor
	for rows.Next() {
			var c CommentWithAuthor
			if err := rows.Scan(
					&c.ID,
					&c.Content,
					&c.CreatedAt,
					&c.UserName,
					&c.UserEmail,
					&c.UserImageUrl,
			); err != nil {
					return nil, err
			}
			out = append(out, c)
	}
	if err := rows.Err(); err != nil {
			return nil, err
	}
	return out, nil
}

func DeleteCommentById(commentId int64) error {
	const query = `
			DELETE FROM comments
			WHERE id = $1
	`

	result, err := db.DB.Exec(query, commentId)
	if err != nil {
			return fmt.Errorf("could not execute delete for comment %d: %w", commentId, err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
			return fmt.Errorf("could not fetch rows affected after deleting comment %d: %w", commentId, err)
	}
	if rowsAffected == 0 {
			return fmt.Errorf("no comment found with id %d", commentId)
	}

	return nil
}