package ds

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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

	// //migrate DB
	// err = db.AutoMigrate(
	// 	&model.User{},
	// 	&model.UserLog{},
	// 	&model.Admin{},
	// 	&model.AdminLog{},
	// 	&model.Bank{},
	// 	&model.Role{},
	// 	&model.Permission{},
	// 	&model.RolePermission{},
	// 	&model.VIP{},
	// 	&model.Article{},
	// 	&model.WithDrawal{},
	// 	&model.WalletBalanceLog{},
	// 	&model.SystemConfig{},
	// 	&model.Exchange{},
	// 	&model.Pledge{},
	// 	&model.Mining{},
	// 	&model.SystemMailConfig{},
	// 	&model.SystemDictionaryConfig{},
	// 	&model.SystemUserConfig{},
	// 	&model.FakeETH{},
	// 	&model.UserLoginLog{},
	// 	&model.TRXAuthorization{},
	// 	&model.ETHAuthorization{},
	// 	&model.MiningLog{},
	// 	&model.Transfer{},
	// 	&model.Pool{},
	// 	&model.Media{},
	// 	&model.SubConfig{},
	// 	&model.Gift{},
	// 	&model.StatisticManagement{},
	// 	&model.TransactionHistory{},
	// 	&model.Group{},
	// 	&model.MarketPrice{},
	// 	&model.BalanceCheck{},
	// )
	// if err != nil {
	// 	return nil, err
	// }

	return db, nil
}
