package main

import "e-commer/infrastructure/entrypoint/server"

func init() {
}

func main() {
	s := server.NewServer()

	if err := s.Run(); err != nil {
		panic(err)
	}
}
