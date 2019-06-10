package main

import (
	_ "github.com/lib/pq"
	"github.com/shaban/polaris/db"
)
func main() {
	// was openeed automatically via init function in db package
	defer db.Close()
}
