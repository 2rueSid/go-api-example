// package to define config
package config

import (
	"context"
	"log"

	"github.com/2rueSid/go-api-example/prisma/db"
)

// Variables which used in database queries
var (
	Client  = db.NewClient()
	Context = context.Background()
)

// Connect to the DB, and return client instance
func Connect() db.PrismaClient {
	if err := Client.Prisma.Connect(); err != nil {
		log.Fatalf("Error while connection to DB \n%v", err)
	}

	return *Client
}

// Close current connection
func Disconnect() {
	if err := Client.Prisma.Disconnect(); err != nil {
		log.Fatalf("Error while closing DB connection \n%v", err)
	}
}
