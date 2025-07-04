package models

import (
	"time"

	"github.com/youssefrag/SoundSphere/db"
)

// StoreRefreshToken saves a new refresh‐token record.
func StoreRefreshToken(userID int64, jti string, expires time.Time) error {
  const q = `
    INSERT INTO refresh_tokens (user_id, jti, expires_at)
    VALUES ($1, $2, $3)
  `
  _, err := db.DB.Exec(q, userID, jti, expires)
  return err
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