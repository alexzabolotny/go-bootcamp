package main

import "fmt"

// User is here
type User struct {
	ID             int
	Name, Location string
}

// Greetings is here
func (u *User) Greetings() string {
	return fmt.Sprintf("Hi %s from %s",
		u.Name, u.Location)
}

// NewPlayer is player constructor
func NewPlayer(id int, name string, location string, gameid int) *Player {
	return &Player{
		&User{id, name, location},
		gameid,
	}
}

// Player is here
type Player struct {
	*User
	GameID int
}

func main() {
	p := NewPlayer(1, "Jim", "LA", 1)
	fmt.Println(p.Greetings())
}
