package repository
import (
	"database/sql"
	"errors"
	"time"
	"Quiz-3/structs"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
)
var jwtSecret = []byte("your_secret_key") 

func GenerateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), 
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenStr string) (int, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return jwtSecret, nil
	})
	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return int(claims["user_id"].(float64)), nil
	}

	return 0, jwt.ErrSignatureInvalid
}
func RegisterUser(db *sql.DB, username, password, createdBy string) error {
	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Get the latest ID from the database
	var lastID int
	err = db.QueryRow(`SELECT COALESCE(MAX(id), 0) FROM users`).Scan(&lastID)
	if err != nil {
		return errors.New("failed to get the last user ID: " + err.Error())
	}

	// Increment the ID for the new user
	newID := lastID + 1

	// Execute the INSERT query with the new ID
	_, err = db.Exec(`
		INSERT INTO users (id, username, password, created_at, created_by) 
		VALUES ($1, $2, $3, $4, $5)
	`, newID, username, string(hashedPassword), time.Now(), createdBy)

	if err != nil {
		return errors.New("failed to register user: " + err.Error())
	}

	return nil
}

func LoginUser(db *sql.DB, username, password string) (string, error) {
	var user structs.Users
	var hashedPassword string

	err := db.QueryRow("SELECT id, password FROM users WHERE username = $1", username).Scan(&user.ID, &hashedPassword)
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	// Compare passwords
	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return "", errors.New("invalid username or password")
	}

	token, err := GenerateToken(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
func GetUserByToken(db *sql.DB, token string) (structs.Users, error) {
	var user structs.Users

	userID, err := ParseToken(token)
	if err != nil {
		return user, errors.New("invalid token")
	}

	err = db.QueryRow("SELECT id, username, created_at, created_by, modified_at, modified_by FROM users WHERE id = $1", userID).
		Scan(&user.ID, &user.Username, &user.CreatedAt, &user.CreatedBy, &user.ModifiedAt, &user.ModifiedBy)

	if err != nil {
		return user, errors.New("user not found")
	}

	return user, nil
}
