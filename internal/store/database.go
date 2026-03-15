package store

import "strconv"

type User struct {
	Email string
	Id    string
	Name  string
}

var database = map[string]User{
	"1": {
		Email: "user1@example.com",
		Id:    "1",
		Name:  "User One",
	},
	"2": {
		Email: "user2@example.com",
		Id:    "2",
		Name:  "User Two",
	},
}

func GetUserById(clientId string) (User, bool) {
	user, ok := database[clientId]
	return user, ok
}

func UpdateUserById(clientId string, updatedUser User) bool {
	_, ok := database[clientId]
	if !ok {
		return false
	}
	database[clientId] = updatedUser
	return true

}

func GetAllUsers() []User {
	var users []User
	for _, user := range database {
		users = append(users, user)
	}
	return users
}

func CreateUser(newUser User) {
	// Incremental ID generation (for simplicity)
	newId := len(database) + 1
	newUser.Id = strconv.Itoa(newId)
	database[newUser.Id] = newUser
}
