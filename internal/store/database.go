package store

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
	// Incremental ID generation (for simplicity)
	newId := len(database) + 1
	newUser.Id = newId
	database = append(database, newUser)

	return newUser
}
