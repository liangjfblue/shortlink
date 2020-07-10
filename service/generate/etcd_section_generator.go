/**
 *
 * @author liangjf
 * @create on 2020/7/9
 * @version 1.0
 */
package generate

import (
	"encoding/json"
	"shortlink/common"
	"shortlink/models/etcdv3"
	"sync"

	"github.com/sirupsen/logrus"
)

/**
基于etcd+发号服务组成的分段分发
etcd作为元数据中心存储已分发id的id段区间(比如区间是10000)
各个发号服务到etcd取id区间, 并且更新etcd的id段区间为下一阶段(同一个事务)
所以, 即可有发号服务挂了也只是丢失一段id区间, id还是保持递增
*/
type etcdSectionGenerate struct {
	opts Options

	curNum int64
	endNum int64

	rw     sync.RWMutex
	onceDo sync.Once
}

type idSection struct {
	StartNum int64 `json:"startNum"`
	EndNum   int64 `json:"endNum"`
}

func NewEtcdSectionGenerate() IGenerate {
	return &etcdSectionGenerate{
		curNum: -1,
		endNum: 0,
	}
}

//生成短码id
func (g *etcdSectionGenerate) Create(opts ...Option) (int64, error) {
	g.opts = defaultOptions
	for _, o := range opts {
		o(&g.opts)
	}

	//TODO 分布式锁
	//TODO defer 释放分布式锁

	//本地没有id段/已用完, 去请求etcd获取
	if g.curNum < 0 || g.curNum == g.endNum {
		if err := g.getRange(opts...); err != nil {
			return 0, nil
		}
	} else {
		g.curNum++
	}

	return g.curNum, nil
}

func (g *etcdSectionGenerate) getRange(opts ...Option) error {
	val, err := etcdv3.Get(common.EtcdCreateIdDir)
	if err != nil {
		logrus.Errorf("get etcd id section err:%s", err.Error())
		return err
	}

	var idSection idSection
	if err := json.Unmarshal([]byte(val), &idSection); err != nil {
		logrus.Errorf("get etcd id section json err:%s", err.Error())
		return err
	}

	idSection.StartNum = idSection.EndNum + 1
	idSection.EndNum = idSection.EndNum + common.ShortIdIncrStep

	data, err := json.Marshal(idSection)
	if err != nil {
		return err
	}

	if err := etcdv3.Update(common.EtcdCreateIdDir, string(data)); err != nil {
		return err
	}

	//更新当前节点的id段
	g.curNum = idSection.StartNum
	g.endNum = idSection.EndNum

	return nil
}
