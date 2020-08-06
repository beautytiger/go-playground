package main
// https://medium.com/@felipedutratine/leader-election-in-go-with-etcd-2ca8f3876d79

import (
	"context"
	"flag"
	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/clientv3/concurrency"
	"log"
	"time"
)
// https://kubernetes.io/blog/2019/08/30/announcing-etcd-3-4/
// https://etcd.io/docs/v3.4.0/dev-guide/interacting_v3/
// https://etcd.io/docs/v3.4.0/op-guide/container/
// docker run --rm -e ALLOW_NONE_AUTHENTICATION=yes -p2379:2379 -p2380:2380 bitnami/etcd:3.4.9
func main() {
	var name = flag.String("name", "", "instance name for this process")
	flag.Parse()

	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"localhost:2379"},
	})

	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	s, err := concurrency.NewSession(cli)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	e := concurrency.NewElection(s, "/test-group-leader/")
	ctx := context.Background()
	// 定时查询leader是谁
	go func(ctx context.Context, e *concurrency.Election) {
		for {
			resp, err := e.Leader(ctx)
			if err != nil {
				if err != concurrency.ErrElectionNoLeader {
					log.Println("leader returned error", err)
					return
				}
				time.Sleep(2*time.Second)
				continue
			}
			log.Println("current leader is ", string(resp.Kvs[0].Value))
			time.Sleep(2*time.Second)
			continue
		}
	}(ctx, e)
	// 执行选主循环
	for {
		if err := e.Campaign(ctx, *name); err != nil {
			log.Fatal(err)
		}
		log.Println("win leader election for ", *name)
		log.Println("Do some work as leader in ", *name)
		time.Sleep(5*time.Second)
		// 如果leader不主动退位且实例已经挂掉的话，会有问题，导致无主
		if err := e.Resign(ctx); err != nil {
			log.Fatal(err)
		}
		log.Println("leader resign ", *name)
	}
}
