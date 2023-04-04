package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/shuklaritvik06/chatongo/backend/database"
	"github.com/shuklaritvik06/chatongo/backend/models"
	"github.com/shuklaritvik06/chatongo/backend/utils"
	"go.mongodb.org/mongo-driver/bson"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	var result models.User
	database.GetDB().Database("user").Collection("user").FindOne(context.Background(), bson.D{{Key: "username", Value: user.Username}}).Decode(&result)
	if utils.CheckPasswordHash(user.Password, result.Password) {
		token, refresh, _ := utils.GetTokens(result.First_Name, result.Last_Name, result.Username)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{
			"token":    token,
			"refresh":  refresh,
			"username": user.Username,
		})
	} else {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(map[string]string{
			"message": "Unauthorized",
		})
	}
}

func Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	hash, _ := utils.HashPassword(user.Password)
	user.Password = hash
	user.Created_At = time.Now().UTC()
	database.GetDB().Database("user").Collection("user").InsertOne(context.Background(), user)
	token, refresh, _ := utils.GetTokens(user.First_Name, user.Last_Name, user.Username)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token":    token,
		"refresh":  refresh,
		"username": user.Username,
	})
}

func Refresh(w http.ResponseWriter, r *http.Request) {
	var user models.User
	_ = json.NewDecoder(r.Body).Decode(&user)
	utils.CheckAuthenticated(w, r)
	token, refresh, _ := utils.GetTokens(user.First_Name, user.Last_Name, user.Username)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]string{
		"token":   token,
		"refresh": refresh,
	})
}
