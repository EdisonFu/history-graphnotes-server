package handlers

import (
	"history-graph-notes-server/model"
	"time"
)

var userContribute map[string][]*model.ContributeTime

func (historyService) SetUserContribute(userId, content string) {
	if userContribute == nil {
		userContribute = make(map[string][]*model.ContributeTime)
	}

	var contribute = new(model.ContributeTime)
	loc, _ := time.LoadLocation("Asia/Shanghai")
	contribute.Date = time.Now().In(loc).Format("2006-01-02 15:04:05")
	contribute.Text = content

	userContribute[userId] = append(userContribute[userId], contribute)
}

func (historyService) GetUserContribute(userId string) interface{} {
	if userContribute == nil {
		return nil
	}

	return userContribute[userId]
}
