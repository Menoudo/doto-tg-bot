package main

import (
	"log"

	"golang.org/x/exp/maps"
)

var logger = log.Default()

type MessageKeeper struct {
	doneMessages map[any]map[int]struct{}
}

func New() MessageKeeper {
	return MessageKeeper{make(map[any]map[int]struct{})}
}

func (c *MessageKeeper) GetChannels() []any {
	result := make([]any, 0)
	for channel := range c.doneMessages {
		result = append(result, channel)
	}
	return result
}

func (c *MessageKeeper) GetMessagesCount() int {
	count := 0
	for channel := range c.doneMessages {
		count += len(c.doneMessages[channel])
	}
	return count
}

func (c *MessageKeeper) GetMessages(channel any) []int {
	if messages, ok := c.doneMessages[channel]; ok {
		result := make([]int, len(messages))
		for i, v := range maps.Keys(messages) {
			result[i] = v
		}
		return result
	}
	logger.Print("GetMessages from empty channel")
	return make([]int, 0)
}

func (c *MessageKeeper) DoneMessage(chatID any, messageID int) {
	if _, ok := c.doneMessages[chatID]; !ok {
		c.doneMessages[chatID] = make(map[int]struct{})
	}
	c.doneMessages[chatID][messageID] = struct{}{}
}

func (c *MessageKeeper) UnDoneMessage(chatID any, messageID int) {
	if _, channel_ok := c.doneMessages[chatID]; channel_ok {
		if _, message_ok := c.doneMessages[chatID][messageID]; message_ok {
			delete(c.doneMessages[chatID], messageID)
		} else {
			logger.Print("remove non existing message")
		}
		if len(c.doneMessages[chatID]) == 0 {
			delete(c.doneMessages, chatID)
		}
	} else {
		logger.Print("remove from non existing chat")
	}
}
