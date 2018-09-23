package lib

import "github.com/xuyu/goredis"
import "github.com/astaxie/beego"

var client *goredis.Redis

// connects to redis
func ConnectToCache() {

	// declare error object
	var err error

	// read redis ip and port from config
	strRedisIpPort := beego.AppConfig.String("redisIpPort")

	// connect to redis
	client, err = goredis.Dial(&goredis.DialConfig{Address: strRedisIpPort})
	HandleError(err)

}

// gets value from redis corresponding to a key
// input - string key
// output - byte value
func GetDataFromRedis(strKey string) []byte {

	// get the data from redis
	value, err := client.Get(strKey)

	// return null if key not found
	if err != nil {
		return nil
	}

	// return the value
	return value

}

// sets key value pair in redis
// input - two strings, key and value
func SetDataInRedis(strKey string, strValue string) {

	// set the data
	client.ExecuteCommand("SET", strKey, strValue)

}
