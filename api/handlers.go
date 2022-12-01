package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/go-redis/redis/v8"
	"github.com/takez0o/honestwork-api/utils/crypto"
	"github.com/takez0o/honestwork-api/utils/schema"
)

// todo: fix error handling
// todo: move validation to middleware
func getUserFromAddress(redis *redis.Client, address string) schema.User {
	var user schema.User
	data, err := redis.Do(redis.Context(), "JSON.GET", address).Result()
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = json.Unmarshal([]byte(fmt.Sprint(data)), &user)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return user
}

func HandleGetUser(redis *redis.Client, address string) schema.User {
	user := getUserFromAddress(redis, address)
	return user
}

func HandleUserUpdate(redis *redis.Client, address string, signature string, body []byte) string {
	result := crypto.VerifySignature("post", address, signature)
	if !result {
		return "Wrong signature."
	}

	// new user
	var user schema.User
	err := json.Unmarshal(body, &user) // todo: validate input after unmarshal
	if err != nil {
		fmt.Println("Error:", err)
	}

	// current user in db
	var user_db schema.User
	data, err := redis.Do(redis.Context(), "JSON.GET", address).Result()
	if err != nil {
		fmt.Println("Error:", err)
	}
	err = json.Unmarshal([]byte(fmt.Sprint(data)), &user_db)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// set skills in input json
	user.Skills = user_db.Skills

	// marshal back to bytes
	new_data, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error:", err)
	}

	redis.Do(redis.Context(), "JSON.SET", address, "$", new_data)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return "success"
}

func HandleGetSkills(redis *redis.Client, address string) []schema.Skill {
	user := getUserFromAddress(redis, address)
	return user.Skills
}

func HandleGetSkill(redis *redis.Client, address string, slot string) schema.Skill {
	s, _ := strconv.Atoi(slot)
	user := getUserFromAddress(redis, address)
	return user.Skills[s]
}

func HandleAddSkill(redis *redis.Client, address string, signature string, body []byte) string {
	result := crypto.VerifySignature("post", address, signature)
	if !result {
		return "Wrong signature."
	}

	var skill schema.Skill
	err := json.Unmarshal(body, &skill)
	if err != nil {
		fmt.Println("Error:", err)
	}

	user := getUserFromAddress(redis, address)
	user.Skills = append(user.Skills, skill)
	updated_user, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error:", err)
	}

	redis.Do(redis.Context(), "JSON.SET", address, "$", updated_user)
	if err != nil {
		fmt.Println("Error:", err)
	}
	return "success"
}

// func HandleUpdateSkill(redis *redis.Client, address string, signature string, slot string, body []byte) schema.Skill {
// 	var skill schema.Skill
// 	return skill
// }
