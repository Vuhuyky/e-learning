package connection

import (
	"app/generated/proto/servicegrpc"
	"net/smtp"

	"github.com/gorilla/websocket"
	"github.com/rabbitmq/amqp091-go"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"golang.org/x/time/rate"
	"gorm.io/gorm"
)

func GetPsql() *gorm.DB {
	return dbPsql
}

func GetConnect() Connection {
	return conn
}

func GetRedisClient() *redis.Client {
	return redisClient
}

func GetUpgraderSocket() *websocket.Upgrader {
	return upgraderSocket
}

func GetMapSocket() map[string]*websocket.Conn {
	return mapSocket
}

func GetRabbitmq() *amqp091.Connection {
	return rabbitmq
}

func GetSmtpPort() string {
	return conn.Smpt.Port
}

func GetSmtpHost() string {
	return conn.Smpt.Host
}

func GetAuthSmtp() smtp.Auth {
	return authSmtp
}

func GetSocketEvent() map[string]map[string]*websocket.Conn {
	return mapSocketEvent
}

func GetLimiter() *rate.Limiter {
	return limiter
}

func GetConnectionInfo() Connection {
	return conn
}

func GetLogger() *zap.Logger {
	return logger
}

func GetGrpcClientQuizz() servicegrpc.QuizzServiceClient {
	return grpcClientQuizz
}

func GetGrpcClientStream() servicegrpc.StreamServiceClient {
	return grpcClientStream
}

func GetGrpcClientQuantity() servicegrpc.QuantityServiceClient {
	return grpcClientQuantity
}
