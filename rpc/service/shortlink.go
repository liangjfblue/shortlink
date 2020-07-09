/**
 *
 * @author liangjf
 * @create on 2020/7/8
 * @version 1.0
 */
package service

import (
	"context"
	"errors"
	"shortlink/models/db"
	"shortlink/rpc/proto"
	"shortlink/service/convert"
	"shortlink/service/core"
	"strings"

	"github.com/sirupsen/logrus"
)

type ShortLink struct{}

func (s *ShortLink) Info(ctx context.Context, req *proto.InfoRequest) (resp *proto.InfoRespond, err error) {
	select {
	case <-ctx.Done():
		return resp, errors.New("ctx done")
	default:
	}

	resp = &proto.InfoRespond{
		Code: 0,
		Msg:  "ok",
	}

	strList := strings.Split(req.ShortLink, "/")
	shortCode := strList[len(strList)-1]
	if shortCode == "" {
		logrus.WithField("shortLink", req.ShortLink).Error("shortCode is empty")
		resp.Msg = "shortCode is empty"
		return resp, errors.New("shortCode is empty")
	}

	shortId := convert.Decode(shortCode).(int64)

	tBShortLink, err := db.GetTBShortLink(map[string]interface{}{"short_id": shortId})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"shortId":   shortId,
			"shortCode": shortCode,
			"shortLink": req.ShortLink,
		}).Error("tBShortLink is empty")
		resp.Msg = "shortCode no longLink"
		return resp, errors.New("shortCode no longLink")
	}

	resp.Code = 1
	resp.ShortLink = &proto.ShortLinkInfo{
		ShortId:   int64(tBShortLink.ShortId),
		ShortCode: tBShortLink.ShortCode,
		LongLink:  tBShortLink.LongUrl,
		CreatedAt: tBShortLink.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt: tBShortLink.UpdatedAt.Format("2006-01-02 15:04:05"),
	}

	return
}

func (s *ShortLink) Shorten(ctx context.Context, req *proto.ShortenRequest) (resp *proto.ShortenRespond, err error) {
	select {
	case <-ctx.Done():
		return resp, errors.New("ctx done")
	default:
	}

	resp = &proto.ShortenRespond{
		Code: 0,
		Msg:  "ok",
	}

	if req.LongLink == "" {
		logrus.Errorf("longLink empty")
		resp.Msg = "longLink empty"
		err = errors.New("longLink empty")
		return
	}

	var shortLink string
	shortLink, err = core.GetShortLinkByLongLink(req.LongLink)
	if err != nil {
		logrus.Errorf("get longLink err:%s", err.Error())
		resp.Msg = "shorten longLink to shortLink error"
		return
	}

	resp.Code = 1
	resp.LongLink = shortLink

	return
}
