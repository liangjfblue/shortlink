/**
 *
 * @author liangjf
 * @create on 2020/7/9
 * @version 1.0
 */
package etcdv3

import (
	"context"
	"time"

	"go.etcd.io/etcd/clientv3"
)

var (
	optTimeout    = time.Second
	_etcdV3Client *clientv3.Client
)

func Init(nodes []string) error {
	var err error
	_etcdV3Client, err = clientv3.New(clientv3.Config{
		Endpoints:   nodes,
		DialTimeout: time.Second * 3,
	})
	if err != nil {
		return err
	}
	return nil
}

func Create(key, val string) error {
	ctx, cancel := context.WithTimeout(context.TODO(), optTimeout)
	defer cancel()

	if _, err := _etcdV3Client.Put(ctx, key, val); err != nil {
		return err
	}
	return nil
}

func Delete(key string) error {
	ctx, cancel := context.WithTimeout(context.TODO(), optTimeout)
	defer cancel()

	if _, err := _etcdV3Client.Delete(ctx, key); err != nil {
		return err
	}
	return nil
}

func Update(key, val string) error {
	ctx, cancel := context.WithTimeout(context.TODO(), optTimeout)
	defer cancel()

	if _, err := _etcdV3Client.Put(ctx, key, val); err != nil {
		return err
	}
	return nil
}

func Get(key string, opts ...clientv3.OpOption) (string, error) {
	ctx, cancel := context.WithTimeout(context.TODO(), optTimeout)
	defer cancel()

	resp, err := _etcdV3Client.Get(ctx, key, opts...)
	if err != nil {
		return "", err
	}

	val := ""
	for _, kv := range resp.Kvs {
		val = string(kv.Value)
		break
	}
	return val, nil
}
