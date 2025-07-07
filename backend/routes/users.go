package routes

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/youssefrag/SoundSphere/models"
	"github.com/youssefrag/SoundSphere/utils"
)

func signup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.Save()

	if err != nil {

		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not save user."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "User created successfully."})


}

func login(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = user.ValidateCredentials()

	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "Could not authenticate user."})
		return
	}

	// Refresh token logic

	// 1) Create a unique JTI for this session
	jti := uuid.NewString()

	// 2) Generate the refresh token with that JTI
	refreshToken, err := utils.GenerateRefreshToken(user.ID, jti)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not create refresh token"})
		return
	}

	if err := models.StoreRefreshToken(user.ID, jti, time.Now().Add(7*24*time.Hour)); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "could not persist refresh token"})
		return
	}

	context.SetCookie(
    "refresh_token",           // name
    refreshToken,              // value
    7*24*3600,                 // maxAge = 7 days in seconds
    "/",                       // path
    "",               // domain (or your real domain in prod)
    false,                     // secure (false if not HTTPS locally)
    true,                      // httpOnly
)




	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not authenticate user."})
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Login successful!",
		"token": token,  
		"user": gin.H{
    	"email": user.Email,
			"name": user.Name,
  	}, 
	})

}

func logout(c *gin.Context) {
  // pull the refresh token identifier (JTI) out of the HTTP-only cookie
  rtCookie, cookieErr := c.Cookie("refresh_token")
  if cookieErr != nil {
    // there was no cookie → nothing to revoke
    c.Status(http.StatusNoContent)
    return
  }

  // now delete exactly that token (or all for this user, if that's your goal)
  if err := models.RemoveRefreshTokenByJTI(rtCookie); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "could not revoke token"})
    return
  }

  // clear the cookie
  c.SetCookie("refresh_token", "", -1, "/", "", true, true)
  c.Status(http.StatusOK)
}


func refresh(c *gin.Context) {
  // 1) Read the incoming refresh cookie
  rt, err := c.Cookie("refresh_token")
  if err != nil {
    c.JSON(http.StatusUnauthorized, gin.H{"error": "no refresh token"})
    return
  }

  // 2) Parse & validate refresh JWT
  token, err := jwt.Parse(rt, func(token *jwt.Token) (interface{}, error) {
    return []byte(os.Getenv("REFRESH_SECRET")), nil
  })
  if err != nil || !token.Valid {
    c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid refresh token"})
    return
  }
  claims := token.Claims.(jwt.MapClaims)
  userID := int64(claims["sub"].(float64))
  oldJti := claims["jti"].(string)

  // 3) Revoke the old refresh token JTI
  if err := models.RevokeRefreshToken(userID, oldJti); err != nil {
    // log if you want, but don’t block the flow
  }

  // 4) Issue & store a brand‐new refresh token
  newJti := uuid.NewString()
  newRT, err := utils.GenerateRefreshToken(userID, newJti)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "could not rotate refresh token"})
    return
  }
  if err := models.StoreRefreshToken(userID, newJti, time.Now().Add(7*24*time.Hour)); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "could not persist refresh token"})
    return
  }

  // 5) Reset the HTTP‐only cookie
  c.SetCookie(
    "refresh_token",
    newRT,
    int(7*24*time.Hour.Seconds()), // 7 days
    "/",
    "localhost",  // use your domain in prod
    false,        // secure=true if HTTPS
    true,         // httpOnly
  )

  // 6) Load the user's public info
  userObj, err := models.GetUserByID(userID)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "could not load user"})
    return
  }

  // 7) Issue a fresh access token (include email claim)
  at, err := utils.GenerateToken(userObj.Email, userObj.ID)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{"error": "could not create access token"})
    return
  }

  // 8) Respond with the new access token and user info
  c.JSON(http.StatusOK, gin.H{
    "access_token": at,
    "user": gin.H{
      "id":    userObj.ID,
      "email": userObj.Email,
      "name":  userObj.Name,
    },
  })
}
