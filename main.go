package main

import (
	"github.com/devKiratu/gator/internal/config"
	_ "github.com/lib/pq"
)

func main() {

	config.StartGator()

}
