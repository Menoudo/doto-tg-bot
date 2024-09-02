package main

import (
	"log"
	"sync"

	"github.com/go-telegram/bot/models"
	"golang.org/x/exp/maps"
)

var logger = log.Default()

type MessageKeeper struct {
	doneMessages     map[int64]map[int]models.Message
	doneMessagesLock sync.RWMutex
}

func NewMessageKeeper() *MessageKeeper {
	return &MessageKeeper{make(map[int64]map[int]models.Message), sync.RWMutex{}}
}

func (c *MessageKeeper) GetChannels() []int64 {
	result := make([]int64, 0)
	c.doneMessagesLock.RLock()
	defer c.doneMessagesLock.RUnlock()
	for channel := range c.doneMessages {
		result = append(result, channel)
	}
	return result
}

func (c *MessageKeeper) GetMessagesCount() int {
	count := 0
	c.doneMessagesLock.RLock()
	defer c.doneMessagesLock.RUnlock()
	for channel := range c.doneMessages {
		count += len(c.doneMessages[channel])
	}
	return count
}

func (c *MessageKeeper) GetMessages(channel int64) []int {
	c.doneMessagesLock.RLock()
	defer c.doneMessagesLock.RUnlock()
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

func (c *MessageKeeper) DoneMessage(message *models.Message) {
	c.doneMessagesLock.Lock()
	defer c.doneMessagesLock.Unlock()
	if _, ok := c.doneMessages[message.Chat.ID]; !ok {
		c.doneMessages[message.Chat.ID] = make(map[int]models.Message)
	}
	c.doneMessages[message.Chat.ID][message.ID] = *message
}

func (c *MessageKeeper) UnDoneMessage(message *models.Message) {
	c.doneMessagesLock.Lock()
	defer c.doneMessagesLock.Unlock()
	chatID := message.Chat.ID
	messageID := message.ID
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
