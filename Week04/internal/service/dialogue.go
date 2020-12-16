package service

import (
	"context"
	"log"

	dto "dialogue/api/dialogue/service/v1"
	"dialogue/internal/biz"

	"github.com/google/wire"
)

var Provider = wire.NewSet(biz.NewDialogueManage, new(*DialogueService))

type DialogueService struct {
	dm *biz.DialogueManage
}

func NewDialogueService(dm *biz.DialogueManage) (d *DialogueService, closeFunc func(), err error) {
	d = &DialogueService{
		dm: dm,
	}

	closeFunc = func() {
		if err := d.dm.Close(); err != nil {
			log.Printf("close error(%v)", err)
		}
	}

	return
}

func (d *DialogueService) AddQA(ctx context.Context, req *dto.AddQARequest) (*dto.AddQAReply, error) {
	d.dm.Learn(req.Question, req.Answers)
	return &dto.AddQAReply{}, nil
}

func (d *DialogueService) Talk(ctx context.Context, req *dto.TalkRequest) (*dto.TalkReply, error) {
	reply := d.dm.SmartReply(req.Query)
	return &dto.TalkReply{Tts: reply}, nil
}
