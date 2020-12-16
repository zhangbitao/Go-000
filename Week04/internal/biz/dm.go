package biz

import (
	"dialogue/internal/data"

	"github.com/google/wire"
)

var Provider = wire.NewSet(NewDialogueManage)

type DialogueManage struct {
	repository data.Cacher
	db         *data.DB
	es         *data.ES
}

func NewDialogueManage(cache data.Cacher) *DialogueManage {
	return &DialogueManage{
		repository: cache,
	}
}

func (dm *DialogueManage) Learn(question string, answers []string) {
	if question == "" || len(answers) == 0 {
		return
	}

	dm.repository.Set(question, answers)
	dm.db.SaveQA(question, answers)
	dm.es.SaveQA(question, answers)
}

func (dm *DialogueManage) SmartReply(query string) string {
	if query == "" {
		return "请告诉我你的问题"
	}

	reply := dm.repository.Get(query)
	if reply != "" {
		return reply
	}

	replyList := dm.es.SearchQA(query)
	if len(replyList) == 0 {
		return "对不起，我还需要学习，暂时不能回答你的问题"
	}

	return replyList[0]
}

func (dm *DialogueManage) Close() error {
	return dm.db.Close()
}
