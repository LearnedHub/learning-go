package etcd

import (
	"context"
	"fmt"
	"time"

	"github.com/LearnHub/learning-go/pkg/learning"
	"go.etcd.io/etcd/client/v3"
)

type LearnEtcd struct {
	learning.Learning
}

func (learner LearnEtcd) Learn() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"localhost:2379", "localhost:22379", "localhost:32379"},
		DialTimeout: 5 * time.Second,
	})
	defer cli.Close()

	if err != nil {
		// handle error!
		fmt.Println("error ", err)
		panic(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 123)
	_, err = cli.Put(ctx, "sample_key", "sample_value")
	cancel()

	if err != nil {
		// handle error!
		fmt.Println("error ", err)
		panic(err)
	}
}

func (learner LearnEtcd) Kind() string {
	return "etcd"
}
