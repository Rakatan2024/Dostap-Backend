package entity

import (
	"time"

	"github.com/golang-jwt/jwt"
)

var (
	JWTKey                  = []byte("sercet_key")
	EmtpyString             = ""
	VerificationLinkURL     = "http://92.38.48.85:80/auth/confirmUserAccount?link="
	NilID               int = -1
)

const (
	BucketName                   = "test-bucket-dostap"
	VerificationSecretCodeLength = 16
	Charset                      = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	PicturesLinkNameLength       = 32
)

type User struct {
	ID              int
	FirstName       string `json:"first_name"`
	LastName        string `json:"last_name"`
	EncryptedPass   string `json:"password"`
	AvatarLink      string `json:"avatar_link"`
	Gender          string `json:"gender"`
	Age             int64  `json:"age"`
	PhoneNum        string `json:"phone_number"`
	ResidenceCity   string `json:"city_of_residence"`
	Description     string `json:"description"`
	Email           string `json:"email"`
	IsEmailVerified bool
	Username        string `json:"username"`
	Interests       []int  `json:"interests"`
}

type Event struct {
	ID                 int
	CreatorID          int
	OrganizerUsernames []string  `json:"organizerUsernames"`
	EventName          string    `json:"eventName"`
	FormatID           int       `json:"formatID"`
	Address            string    `json:"address"`
	CoordinateX        float64   `json:"coordinateX"`
	CoordinateY        float64   `json:"coordinateY"`
	Capacity           int       `json:"capacity"`
	Link               string    `json:"link"`
	Links              []string  `json:"links"`
	Description        string    `json:"description"`
	PrivacyID          int       `json:"privacyID"`
	InterestIDs        []int     `json:"interestIDs"`
	StartTime          time.Time `json:"startTime"`
	EndTime            time.Time `json:"endTime"`
	City               string    `json:"city"`
}

type Email struct {
	ID         int
	UserID     int
	Email      string
	SecretCode string
	IsUsed     bool
	ExpTime    time.Time
}

type Credentials struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Interest struct {
	ID       int
	Name     string
	Category string
}
type Claims struct {
	Email    string `json:"email"`
	Sub      int    `json:"sub"`
	Level    string `json:"level"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// struct for jsopn decoding for update user info
type UpdateJson struct {
	Token    TokenData `json:"jwtToken"`
	UserInfo User      `json:"userInfo"`
}

type TokenData struct {
	Token string
}

type ResponseJSON struct {
	Message string
}

type FriendRequest struct {
	ID          int
	SenderID    int `json:"senderID"`
	RecipientID int `json:"recipientID"`
	IsAccepted  bool
}
