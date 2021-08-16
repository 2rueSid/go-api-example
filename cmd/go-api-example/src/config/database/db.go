// Database contains defined database connection instance and helper function.
package database

import (
	"context"
	"log"

	"github.com/2rueSid/go-api/cmd/go-api-example/prisma/db"
)

// Variables which used in database queries.
var (
	// Initialize Database connection instance.
	Client = db.NewClient()
	// Get context that use within the prisma connections.
	Context = context.Background()
)

// Connect connects to the DB, and return client instance.
func Connect() db.PrismaClient {
	if err := Client.Prisma.Connect(); err != nil {
		log.Fatalf("Error while connection to DB \n%v", err)
	}

	return *Client
}

// Disconnect close current connection.
func Disconnect() {
	if err := Client.Prisma.Disconnect(); err != nil {
		log.Fatalf("Error while closing DB connection \n%v", err)
	}
}
