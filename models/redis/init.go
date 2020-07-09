/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package redis

import (
	"time"

	"github.com/chasex/redis-go-cluster"
)

var (
	_cluster *redis.Cluster
)

func Init(nodes []string) error {
	var err error
	_cluster, err = redis.NewCluster(
		&redis.Options{
			StartNodes:   nodes,
			ConnTimeout:  50 * time.Millisecond,
			ReadTimeout:  50 * time.Millisecond,
			WriteTimeout: 50 * time.Millisecond,
			KeepAlive:    16,
			AliveTime:    60 * time.Second,
		})
	if err != nil {
		return err
	}

	return nil
}
