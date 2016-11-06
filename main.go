package main

import (
	_ "github.com/lib/pq"
	"github.com/pbk/kit-api/db"
	"github.com/pbk/kit-api/http/util"
)

func main() {
	util.NewServer(db.SetupDB())
}
