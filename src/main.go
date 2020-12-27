package main

import (
	"context"
	"etcd"
	"fmt"
	"go.etcd.io/etcd/clientv3"
	"go.etcd.io/etcd/etcdserver/api/v3rpc/rpctypes"
	"log"
	"time"
)

func main(){
	info := etcd.ServiceInfo{
		Name: "s1",
		IP: "localhost:3001",
	}
	endpoints := []string{"192.168.3.142:2379"}

	Watch(endpoints)
	//Put(endpoints)
	service, err := etcd.NewService(info, endpoints)
	if err != nil {
		log.Fatal(err)
	}
	err = service.Start()
	if err != nil {
		log.Fatal(err)
	}
}
const dialTimeout    = 5 * time.Second
const requestTimeout = 10 * time.Second
func Watch(endpoints []string) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	rch := cli.Watch(context.Background(), "/kk")
	for wresp := range rch {
		for _, ev := range wresp.Events {
			fmt.Printf("%s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}

func Put(endpoints []string) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		log.Fatal(err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	_, err = cli.Put(ctx, "/kk", "sample_value")
	cancel()
	if err != nil {
		switch err {
		case context.Canceled:
			fmt.Printf("ctx is canceled by another routine: %v\n", err)
		case context.DeadlineExceeded:
			fmt.Printf("ctx is attached with a deadline is exceeded: %v\n", err)
		case rpctypes.ErrEmptyKey:
			fmt.Printf("client-side error: %v\n", err)
		default:
			fmt.Printf("bad cluster endpoints, which are not etcd servers: %v\n", err)
		}
	}
	//cli, err := clientv3.New(clientv3.Config{
	//	Endpoints:   endpoints,
	//	DialTimeout: dialTimeout,
	//})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer cli.Close()
	//
	//ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	//_, err = cli.Put(ctx, "sample_key", "sample_value")
	//cancel()
	//if err != nil {
	//	log.Fatal(err)
	//}
}