package main

type User struct {
	ID       string
	Email    string
	Password string
}

type Account struct {
	ID      string
	UserID  string
	Balance int64
}
