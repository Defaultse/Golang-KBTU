package http

import (
	"api/internal/models"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"strings"
	"time"
)

const (
	salt = "jojozxCCMnb98736Jxz"
	signingKey = "qrkjk#4#%35FSFJlja#4353KSFjH"
	tokenTTL = 12 * time.Hour
	authHeader = "Authorization"
)

type UserClaims struct {
	UserID int `json:"user_id"`
	jwt.StandardClaims
}

func (s *Server) registration(w http.ResponseWriter, r *http.Request) {
	user := new(models.User)
	if err := json.NewDecoder(r.Body).Decode(user); err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}
	user.Password = generateHash(user.Password)
	if err := s.store.User().Create(r.Context(), user); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "DB err: %v", err)
		return
	}
}

func (s *Server) login(w http.ResponseWriter, r *http.Request) {
	credentials := new(models.User)
	if err := json.NewDecoder(r.Body).Decode(credentials); err != nil {
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	user, err := s.store.User().GetUser(r.Context(), credentials.Email, generateHash(credentials.Password))
	if err != nil {
		fmt.Fprintf(w, "Unknown err: %v", err)
		return
	}

	token, err := generateToken(user)
	fmt.Fprintf(w, "Token: %s", token)
}

func VerifyToken(w http.ResponseWriter, r *http.Request) int {
	reqToken := r.Header.Get(authHeader)
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	token, err := jwt.ParseWithClaims(reqToken, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})

	var userId int
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		userId = claims.UserID
		fmt.Printf("Token verifyed ID: %v, Expires: %v", claims.UserID, claims.StandardClaims.ExpiresAt)
	} else {
		fmt.Fprintf(w, "Unknown err: %v", err)
	}
	return userId
}

func generateToken(user *models.User) (string, error){
	var err error
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = user.ID
	atClaims["exp"] = time.Now().Add(tokenTTL).Unix()
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(signingKey))
	if err != nil {
		return "", err
	}
	return token, nil
}

func generateHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}