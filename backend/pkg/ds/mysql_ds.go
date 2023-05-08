package ds

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"loan-back-services/conf"
	"loan-back-services/pkg/logger"
	"loan-back-services/pkg/model"
)

func LoadDB() (*gorm.DB, error) {
	//host := os.Getenv("MYSQL_HOST")
	//port := os.Getenv("MYSQL_PORT")
	//user := os.Getenv("MYSQL_USER")
	//pass := os.Getenv("MYSQL_PASS")
	//name := os.Getenv("MYSQL_NAME")
	//
	//dsn := fmt.Sprintf(
	//	"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
	//	user, pass, host, port, name,
	//)

	db, err := gorm.Open(mysql.Open(conf.MysqlDNS()), &gorm.Config{
		// Logger: logger.Default.LogMode(logger.Info),
		// Logger: logger.Default.LogMode(logger.Error),
	})
	if err != nil {
		return nil, err
	}

	logger.Sugar.Info("Successfully connected to MySQL")

	//migrate DB
	err = db.AutoMigrate(
		&model.Admin{},
		&model.AdminLog{},
		&model.User{},
		&model.LoanPackage{},
		&model.LoanPackageLog{},
		&model.UserLoan{},
		&model.UserLoanLog{},
	)
	if err != nil {
		return nil, err
	}

	return db, nil
}
