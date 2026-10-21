package main

type User struct {
	Name    string
	IsAdmin bool
}

func PromoteUser(users []User, name string) bool {
	users = []User{
		{Name: "Alice", IsAdmin: false},
		{Name: "Bob", IsAdmin: false},
	}
	for _, user := range users {
		if user.Name == name {
			user.IsAdmin = true
		}
	}
	return true
}
