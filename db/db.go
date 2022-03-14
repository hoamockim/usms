package db

import (
	"context"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"usms/db/models"
	"usms/pkg/cache"
	"usms/pkg/configs"
)

const (
	MySQLDialect = "mysql"
)

type Config struct {
	driver        string
	connection    string
	maxLifeTime   int32
	maxConnection int16
}

//dbFacade Build gorm run with raw query -> auto generate raw query
type dbFacade struct {
	app   string
	ctx   context.Context
	orm   *gorm.DB
	cache cache.Adapter
}

var fcd *dbFacade

func init() {
	var (
		orm *gorm.DB
		err error
	)

	if orm, err = gorm.Open(MySQLDialect, configs.DBConnectionString()); err != nil {
		panic(err)
	}
	orm.DB().SetMaxOpenConns(300)
	orm.DB().SetMaxIdleConns(10)
	fcd = new(dbFacade)
	fcd.orm = orm
}

//addDbContext
func addDbContext(key string, value interface{}) context.Context {
	ctx := context.Background()
	context.WithValue(ctx, key, value)
	return ctx
}

//GetById
func GetById(id int, tableName string, data interface{ models.ModelCache }) error {
	/*	if data.IsCached() {
		if err := fcd.cache.Get(fmt.Sprintf("%v:db:%v:%v", fcd.app, tableName, id), data); err != nil {
			//TODO: log err
		}
	}*/
	return fcd.orm.New().Table(tableName).
		Where("id = ?", id).
		First(data).Error
}

//Save data into db
func Save(tableName string, data interface{}) error {
	return fcd.orm.New().Table(tableName).Save(data).Error
}

func Update(tableName string, data interface{}, conditions map[string]interface{}) error {
	query := fcd.orm.New().Table(tableName)
	for key, val := range conditions {
		query = query.Where(fmt.Sprintf("%v = ?", key), val)
	}
	return query.Update(data).Error
}

//Filter
func Filter(tableName string, entities interface{}, conditions map[string]interface{}) error {
	query := fcd.orm.New().Table(tableName)
	for key, val := range conditions {
		query = query.Where(fmt.Sprintf("%v = ?", key), val)
	}
	return query.Find(entities).Error
}
