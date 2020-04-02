package store

import usersapi "github.com/xshoji/go-rest-openapi/usersapi/go"

var data map[int32]*usersapi.User

func init() {
	data = make(map[int32]*usersapi.User, 10)
}

func Put(user *usersapi.User) {
	data[user.Id] = user
}

func Get(id int32) *usersapi.User {
	user, ok := data[id]
	if !ok {
		return nil
	}
	return user
}

func GetAll() []*usersapi.User {
	var users []*usersapi.User
	for _, user := range data {
		users = append(users, user)
	}
	return users
}
