package handlers

var userContribute map[string][]string

func (historyService) SetUserContribute(userId, content string) {
	if userContribute == nil {
		userContribute = make(map[string][]string)
	}

	userContribute[userId] = append(userContribute[userId], content)
}

func (historyService) GetUserContribute(userId string) []string {
	if userContribute == nil {
		return nil
	}

	return userContribute[userId]
}
