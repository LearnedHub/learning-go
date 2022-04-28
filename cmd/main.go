package main

import (
	"github.com/LearnHub/learning-go/pkg/learning"
	"github.com/LearnHub/learning-go/pkg/learning/redis"
)

func main() {
	learning.Do(new(redis.RedisLearning))
}
