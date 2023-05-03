package conf

import (
	"crypto/rsa"
	"fmt"
	"loan-back-services/pkg/logger"
	"os"

	"github.com/golang-jwt/jwt"
	"gopkg.in/yaml.v3"
)

type (
	app struct {
		Name   string `yaml:"name"`
		Domain string `yaml:"domain"`
	}

	mongo struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	}

	redis_ struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	}

	mysql struct {
		Host     string `yaml:"host"`
		Port     string `yaml:"port"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
	}

	svc struct {
		Hosts []string `yaml:"hosts"`
		Host  string   `yaml:"host"`
		Name  string   `yaml:"name"`
	}

	rpcSVC struct {
		MarketRPC svc `yaml:"marketRPC"`
		EtherRPC  svc `yaml:"etherRPC"`
		TronRPC   svc `yaml:"tronRPC"`
	}

	rsa_ struct {
		Private string `yaml:"private"`
		Public  string `yaml:"public"`
		Secret  string `yaml:"secret"`

		PublicKey  *rsa.PublicKey
		PrivateKey *rsa.PrivateKey
	}
)

var (
	_c struct {
		App    app    `yaml:"app"`
		Mongo  mongo  `yaml:"mongo"`
		Redis  redis_ `yaml:"redis"`
		Mysql  mysql  `yaml:"mysql"`
		RPCSvc rpcSVC `yaml:"rpcSVC"`
		Rsa    rsa_   `yaml:"rsa"`
	}
)

func init() {
	data, err := os.ReadFile("./conf/conf.yaml")
	if err != nil {
		panic(err)
	}

	if err := yaml.Unmarshal(data, &_c); err != nil {
		panic(err)
	}

	//dir, err := os.Getwd()
	//if err != nil {
	//	logger.Sugar.Error("Error on getting directory : ", err.Error())
	//}

	privateBytes, err := os.ReadFile(_c.Rsa.Private)
	if err != nil {
		logger.Sugar.Fatalf("error on reading english env : %v\n", err)
	}

	prvKey, err := jwt.ParseRSAPrivateKeyFromPEM(privateBytes)
	if err != nil {
		logger.Sugar.Fatalf("error on parsing private key : %v\n", err)
	}
	_c.Rsa.PrivateKey = prvKey

	publicBytes, err := os.ReadFile(_c.Rsa.Public)
	if err != nil {
		logger.Sugar.Fatalf("error on loading public key : %v\n", err)
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(publicBytes)
	if err != nil {
		logger.Sugar.Fatalf("error on parsing public key : %v\n", err)
	}
	_c.Rsa.PublicKey = pubKey
}

func App() *app {
	return &_c.App
}

func MongoDSN() string {
	if _c.Mongo.Username == "" && _c.Mongo.Password == "" {
		return fmt.Sprintf("mongodb://%s:%s", _c.Mongo.Host, _c.Mongo.Port)
	} else {
		return fmt.Sprintf(
			"mongodb://%s:%s@%s/%s",
			_c.Mongo.Username,
			_c.Mongo.Password,
			_c.Mongo.Host,
			_c.Mongo.Database,
		)
	}
}

func Redis() *redis_ {
	return &_c.Redis
}

func MysqlDNS() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		_c.Mysql.Username,
		_c.Mysql.Password,
		_c.Mysql.Host,
		_c.Mysql.Port,
		_c.Mysql.Database,
	)
}

func Mysql() *mysql {
	return &_c.Mysql
}

func Rsa() *rsa_ {
	return &_c.Rsa
}
