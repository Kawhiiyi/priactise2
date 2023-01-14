package dal

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"log"
	"redenvelop-Prac/model"
	"time"
)

var rdb *gorm.DB

var tableName = "rp_receive_record"

func InitDB() {
	dsn := "root:root2023@tcp(127.0.0.1:3306)/tech?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	rdb, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil : err *

	sqlDB, _ := rdb.DB()

	sqlDB.SetMaxIdleConns(10)

	sqlDB.SetMaxOpenConns(100)

	sqlDB.SetConnsMaxLifetime(time.Hour)



}

func QueryByUserId(useId string) (*model.RpReceiveRecord, error){
	var record model.RpReceiveRecord
	err := rdb.Table("rp_receive_record").Where("user_id= ? ", userId).First(&record).Error
	if err != nil {
		log.Printf("Can't find userId amount. userId : %v err : %v", userId, err)
		return nil, err
	}
	return &record, nil
}

func InsertRecord(record *model.RpReceiveRecord) (int64, error) {
	err := rdb.Table("rp_receive_record").Create(&record).Error
	if err != nil {
		log.Printf("insert data err: #{err}\n")
		return 0, err
	}
	return record.Id, nil
}