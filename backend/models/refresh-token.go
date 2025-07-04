package models

import (
	"time"

	"github.com/youssefrag/SoundSphere/db"
)

func StoreRefreshToken(userID int64, jti string, expires time.Time) error {
	// start a transaction
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// 1) insert the fresh token
	const insertQ = `
		INSERT INTO refresh_tokens (user_id, jti, expires_at)
		VALUES ($1, $2, $3)
	`
	if _, err := tx.Exec(insertQ, userID, jti, expires); err != nil {
		return err
	}

	// 2) delete any of this user's tokens that are expired or revoked
	const cleanupQ = `
		DELETE FROM refresh_tokens
		WHERE user_id = $1
		  AND (expires_at < now() OR revoked = TRUE)
	`
	if _, err := tx.Exec(cleanupQ, userID); err != nil {
		return err
	}

	// commit both operations together
	return tx.Commit()
}

// IsRefreshTokenValid returns true if the given jti exists,
// is not expired, and is not revoked.
func IsRefreshTokenValid(userID int64, jti string) bool {
  const q = `
    SELECT COUNT(*) 
    FROM refresh_tokens
    WHERE user_id = $1
      AND jti = $2
      AND revoked = FALSE
      AND expires_at > now()
  `
  var cnt int
  _ = db.DB.QueryRow(q, userID, jti).Scan(&cnt)
  return cnt == 1
}

// RevokeRefreshToken marks the given jti as revoked so it
// can’t be used again.
func RevokeRefreshToken(userID int64, jti string) error {
  const q = `
    UPDATE refresh_tokens
    SET revoked = TRUE
    WHERE user_id = $1 AND jti = $2
  `
  _, err := db.DB.Exec(q, userID, jti)
  return err
}