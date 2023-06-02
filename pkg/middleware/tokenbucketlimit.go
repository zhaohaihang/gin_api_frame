package middleware

import (
	"fmt"
	"gin_api_frame/pkg/e"
	"gin_api_frame/pkg/redis"
	"gin_api_frame/pkg/utils/tokenutil"
	"time"

	"github.com/gin-gonic/gin"
	redigo "github.com/gomodule/redigo/redis"
)

const (
	KeyTokenBucketLimitActivityUser = "tokenbucketlimit:activity:username:%s"
)

func BucketLimit() gin.HandlerFunc {
	return func(c *gin.Context) {
		var data interface{}
		code := e.SUCCESS

		claims := tokenutil.GetTokenClaimsFromContext(c)
		if claims == nil {
			code = e.ErrorUserActivityLimit
		} else {
			username := claims.Username
			key := fmt.Sprintf(KeyTokenBucketLimitActivityUser, username)
			conn := redis.Pool.Get() 
			defer conn.Close()
			rate := 1                                                     
			capacity := 1                                                
			tokens, _ := redigo.Int(conn.Do("hget", key, "tokens"))       
			lastTime, _ := redigo.Int64(conn.Do("hget", key, "lastTime")) 
			now := time.Now().Unix()
			
			existKey, _ := redigo.Int(conn.Do("exists", key))
			if existKey != 1 {
				tokens = capacity
				conn.Do("hset", key, "lastTime", now)
			}
			deltaTokens := int(now-lastTime) * rate 
			if deltaTokens > 1 {
				tokens = tokens + deltaTokens 
			}
			if tokens < 1 {
				code = e.ErrorUserActivityLimit
			} else {
				if tokens > capacity {
					tokens = capacity
				}
				tokens-- 
				conn.Do("hset", key, "lastTime", now)
				conn.Do("hset", key, "tokens", tokens)
				c.Next()
			}
		}
		if code != e.SUCCESS {
			c.JSON(200, gin.H{
				"status": code,
				"msg":    e.GetMsg(code),
				"data":   data,
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
