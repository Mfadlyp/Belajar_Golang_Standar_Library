package main

import (
	"fmt"
	"sort"
)

// membuat struct user
type User struct {
	Name string
	Age  int
}

// membuat alias dri User
type UserSlice []User

// membuat kontrak dri interface sort
func (userSlice UserSlice) Len() int {
	return len(userSlice)
}

func (userSlice UserSlice) Less(i, j int) bool {
	return userSlice[i].Age < userSlice[j].Age
}

func (userSlice UserSlice) Swap(i, j int) {
	// menukar posisi
	userSlice[i], userSlice[j] = userSlice[j], userSlice[i]
}

func main() {
	users := []User{
		{"budya", 30},
		{"cici", 25},
		{"didi", 40},
		{"eko", 55},
		{"patrio", 10},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)
}
