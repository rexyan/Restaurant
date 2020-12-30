package tool

import (
	"Restaurant/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"log"
)

// 抛出全局 DBEngine
var DBEngine *Orm

type Orm struct {
	*xorm.Engine
}

func OrmEngine() (*Orm, error) {
	// 获取数据库相关配置
	dbConfig := GetConfig().DataBase
	// 拼接连接地址
	conn := dbConfig.User + ":" + dbConfig.Password + "@tcp(" + dbConfig.Host + ":" + dbConfig.Port + ")/" + dbConfig.DBName + "?charset=" + dbConfig.Charset
	// 获取新 Engine
	engine, err := xorm.NewEngine(dbConfig.Driver, conn)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	// 设置显示 SQL
	engine.ShowSQL(dbConfig.ShowSQL)

	// 根据 Model 逆向创建数据库表
	if err := engine.Sync2(
		new(model.Member),
		new(model.FoodCategory),
		new(model.Shop),
		new(model.Service),
		new(model.ShopService),
		); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	// 将 Engine 信息，赋值给自定义的 Orm struct
	orm := new(Orm)
	orm.Engine = engine

	// 更新 DBEngine
	DBEngine = orm
	return orm, nil
}
