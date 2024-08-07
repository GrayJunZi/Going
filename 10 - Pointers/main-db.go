package main

import "fmt"

type Database struct {
	user string
}

type Server struct {
	db *Database
}

func (s *Server) GetUserFromDB() string {
	if s.db == nil {
		panic("database == nil hence, is not initialized.")
	}
	return s.db.user
}

func main() {
	s := &Server{
		// db: &Database{
		// 	user: "Test",
		// },
	}
	fmt.Println(s.GetUserFromDB())
}
