package connection

import (
	"flag"
	"log"
)

func init() {
	db := flag.Bool("db", false, "")

	flag.Parse()

	// connect
	// initJwt()
	loadYml()
	makeVariable()
	makeFolder()
	connectPostgresql(*db)
	connectRabbitmq()
	connectRedis()
	connectGrpcService()
	initSmptAuth()
	initLogger()

	log.Println("connect ok")
}
