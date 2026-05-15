package store

import "sync/atomic"

type User struct {
	Email string `json:"email"`
	Id    int    `json:"id"`
	Name  string `json:"name"`
}

var database = []User{
	{
		Email: "user1@example.com",
		Id:    1,
		Name:  "User One",
	},
	{
		Email: "user2@example.com",
		Id:    2,
		Name:  "User Two",
	},
}

var idCounter int32

func init() {
	var max int
	for _, u := range database {
		if u.Id > max {
			max = u.Id
		}
	}
	atomic.StoreInt32(&idCounter, int32(max))
}

func GetUserById(clientId int) (User, bool) {
	for _, user := range database {
		if user.Id == clientId {
			return user, true
		}
	}
	return User{}, false
}

func UpdateUserById(clientId int, updatedUser User) bool {
	for i, user := range database {
		if user.Id == clientId {
			updatedUser.Id = clientId
			database[i] = updatedUser
			return true
		}
	}
	return false
}

func GetAllUsers() []User {
	var users []User
	for _, user := range database {
		users = append(users, user)
	}
	return users
}

func CreateUser(newUser User) User {
	// Concurrency-safe incremental ID generation
	newId := int(atomic.AddInt32(&idCounter, 1))
	newUser.Id = newId
	database = append(database, newUser)

	return newUser
}
