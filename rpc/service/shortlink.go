/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package service

import (
	"context"
	"shortlink/rpc/proto"
)


type ShortLink struct {
}

func (s *ShortLink) Get(ctx context.Context, req *proto.GetRequest) (resp *proto.GetRespond, err error) {
	return
}

func (s *ShortLink) Change(ctx context.Context, req *proto.ChangeRequest) (resp *proto.ChangeRespond, err error) {
	return
}
