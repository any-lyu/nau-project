package model

import "time"

type Pic struct {
	Url string `json:"url"`
}

type Pics []Pic

const (
	OssAccessKey    = "******"
	OssAccessSecret = "******"
	OssEndPoint     = "******"
	OssBucket       = "******"
	OssImageDomain  = "******"
)

var (
	// ShanghaiLocationOffset 是东八区的 offset
	ShanghaiLocationOffset = 8 * 60 * 60

	// ShanghaiLocation 是东八区的 time.Location
	ShanghaiLocation = time.FixedZone("Asia/Shanghai", ShanghaiLocationOffset)
)
