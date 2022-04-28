package main

import (
	"github.com/LearnHub/learning-go/pkg/learning"
	"github.com/LearnHub/learning-go/pkg/learning/mysql"
	"github.com/LearnHub/learning-go/pkg/learning/oop"
	"github.com/LearnHub/learning-go/pkg/learning/redis"
)

func main() {
	learning.Do(new(oop.LearnOop))
	learning.Do(new(redis.LearnRedis))
	learning.Do(new(mysql.LearnMySQL))
	//learning.Do(new(etcd.LearnEtcd))
}
