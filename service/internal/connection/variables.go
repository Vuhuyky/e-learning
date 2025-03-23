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

var (
	dbPsql         *gorm.DB
	redisClient    *redis.Client
	rabbitmq       *amqp091.Connection
	upgraderSocket *websocket.Upgrader

	mapSocket      map[string]*websocket.Conn
	mapSocketEvent map[string]map[string]*websocket.Conn

	authSmtp smtp.Auth

	conn Connection

	// chanel job
	emailChan chan EmailJob_MessPayload

	// rate limit
	limiter *rate.Limiter

	// log
	logger *zap.Logger

	grpcClientQuizz    servicegrpc.QuizzServiceClient
	grpcClientStream   servicegrpc.StreamServiceClient
	grpcClientQuantity servicegrpc.QuantityServiceClient
)
