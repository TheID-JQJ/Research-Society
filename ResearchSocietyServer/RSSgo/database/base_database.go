package database

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"hkc.ink/rss/model"
	"hkc.ink/rss/utils"
)

var db *gorm.DB
var viperUtil utils.ViperUtil

// 初始化数据库
func InitDB() *gorm.DB {
	//读取配置文件，加载配置
	mysqlConfig := viperUtil.Read("config.yml", "mysql")
	username := mysqlConfig["username"].(string)
	password := mysqlConfig["password"].(string)
	address := mysqlConfig["address"].(string)
	databases := mysqlConfig["databases"].(string)
	params := mysqlConfig["params"].(string)

	dsn := username + ":" + password + "@tcp(" + address + ")/" + databases + "?" + params

	//连接数据库
	if initDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		log.Println("failed to connect database:" + err.Error())
		return nil
	} else {
		db = initDB
	}

	// 迁移 schema
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.File{})
	db.AutoMigrate(
		&model.UserAccount{}, &model.UserAppoint{}, &model.UserDonate{}, &model.UserFriend{}, &model.UserGroup{}, &model.UserHistory{},
		&model.UserInformation{}, &model.UserOperateInformation{}, &model.UserOrders{}, &model.UserPassword{}, &model.UserRealName{}, &model.UserRent{},
		&model.UserRole{}, &model.UserSeat{}, &model.UserShoppingCart{}, &model.UserStar{}, &model.UserTarget{},
	)
	db.AutoMigrate(&model.GoodsInformation{}, &model.GoodsKind{})
	db.AutoMigrate(&model.Seat{}, &model.SeatKind{})
	db.AutoMigrate(&model.PublicResource{}, &model.PublicResourceKind{})
	db.AutoMigrate(&model.Finance{}, &model.FinanceKind{})
	db.AutoMigrate(&model.Feedback{}, &model.Notice{})
	db.AutoMigrate(&model.TargetType{})
	db.AutoMigrate(&model.Group{})

	return db
}

// 获取gorm.DB对象
func GetDB() *gorm.DB {
	return db
}

// 关闭数据库连接
func CloseDB(db *gorm.DB) {
	//获得sql.DB对象
	sqlDB, err1 := db.DB()
	if err1 != nil {
		log.Println("failed to close database:" + err1.Error())

		return
	}

	//关闭数据库连接
	err2 := sqlDB.Close()
	if err2 != nil {
		log.Println("failed to close database:" + err2.Error())

		return
	}
}
