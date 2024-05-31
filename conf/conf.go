package conf

import (
	"github.com/gin-gonic/gin"
	logging "github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

var (
	AppMode  string
	HttpPort string

	Db         string
	DbHost     string
	DbPort     string
	DbUser     string
	DbPassWord string
	DbName     string

	RedisDb     string
	RedisAddr   string
	RedisPw     string
	RedisDbName string

	RabbitMQ         string
	RabbitMQUser     string
	RabbitMQPassWord string
	RabbitMQHost     string
	RabbitMQPort     string

	EtcdHost string
	EtcdPort string

	UserServiceAddress    string
	TaskServiceAddress    string
	GatewayServiceAddress string

	RabbitMqTaskQueue string = "rabbitmq-task-queue"
)

func Init() {
	// 读取配置
	file, err := ini.Load("./conf/config.ini")
	if err != nil {
		panic(err)
		return
	}
	LoadAppMOde(file)
	LoadRedis(file)
	LoadMysql(file)
	LoadRabbitmq(file)
	LoadEtcd(file)
	LoadServer(file)

	if AppMode == "release" {
		logging.SetLevel(logging.ErrorLevel)
		gin.SetMode(gin.ReleaseMode)
	}

}

func LoadMysql(file *ini.File) {
	Db = file.Section("mysql").Key("Db").String()
	DbHost = file.Section("mysql").Key("DbHost").String()
	DbPort = file.Section("mysql").Key("DbPort").String()
	DbUser = file.Section("mysql").Key("DbUser").String()
	DbPassWord = file.Section("mysql").Key("DbPassWord").String()
	DbName = file.Section("mysql").Key("DbName").String()
}

func LoadAppMOde(file *ini.File) {
	AppMode = file.Section("appmode").Key("AppMode").String()
	HttpPort = file.Section("appmode").Key("HttpPort").String()
}

func LoadRedis(file *ini.File) {
	RedisDb = file.Section("redis").Key("RedisDb").String()
	RedisAddr = file.Section("redis").Key("RedisAddr").String()
	RedisDbName = file.Section("redis").Key("RedisDbName").String()
	RedisPw = file.Section("redis").Key("RedisPw").String()
}

func LoadRabbitmq(file *ini.File) {
	RabbitMQ = file.Section("rabbitmq").Key("RabbitMQ").String()
	RabbitMQUser = file.Section("rabbitmq").Key("RabbitMQUser").String()
	RabbitMQPassWord = file.Section("rabbitmq").Key("RabbitMQPassWord").String()
	RabbitMQHost = file.Section("rabbitmq").Key("RabbitMQHost").String()
	RabbitMQPort = file.Section("rabbitmq").Key("RabbitMQPort").String()
}

func LoadEtcd(file *ini.File) {
	EtcdHost = file.Section("etcd").Key("EtcdHost").String()
	EtcdPort = file.Section("etcd").Key("EtcdPort").String()
}

func LoadServer(file *ini.File) {
	GatewayServiceAddress = file.Section("server").Key("GatewayServiceAddress").String()
	UserServiceAddress = file.Section("server").Key("UserServiceAddress").String()
	TaskServiceAddress = file.Section("server").Key("TaskServiceAddress").String()
}
