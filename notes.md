## Stack

### Frontend: Vue/Typescript, Tailwind

### Backend: Go, Gin

### Devops: Docker

## Color Pallette

#98C970
#0E0E0F
#0DE27C
#FF4D42
#DCCA7F
#FBBF00
#FFA900
#FFFFFF
#4636FF

#055a32

# Database Schema

## User

- UserID (primary key)
- Name
- Email
- Password
- Picture Link

## Song

- SongID (primary key)
- UserID (foreign key)
- Song link
- Genre
- Rating Overall

## Comment

- CommentID (primary key)
- SongID (foreign key)
- UserID (foreign key)
- Content
- Rating
