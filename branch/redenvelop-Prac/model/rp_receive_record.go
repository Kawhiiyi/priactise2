package model

import "time"

type RpReceiveRecord struct {
	Id            int64
	UserId        string
	GroupChatId   string
	ReceiveAmount int64
	password      string
	RpId          string
	CreateTime    time.Time
	NotifyTime    time.Time
}
