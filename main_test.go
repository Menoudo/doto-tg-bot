package main_test

import (
	"testing"

	bot "todo-tg-bot"

	"github.com/go-telegram/bot/models"
	"github.com/stretchr/testify/require"
)

func TestKeeper_empty(t *testing.T) {
	c := bot.NewMessageKeeper()
	ch := c.GetChannels()
	require.Empty(t, ch)
}

func TestKeeper_one(t *testing.T) {
	c := bot.NewMessageKeeper()
	msg1 := &models.Message{Chat: models.Chat{ID: 1}, ID: 1000}
	c.DoneMessage(msg1)
	ch := c.GetChannels()
	require.Len(t, ch, 1)
	msg1.ID = 1001
	c.DoneMessage(msg1)
	require.Len(t, ch, 1)
}

func TestKeeper_many(t *testing.T) {
	c := bot.NewMessageKeeper()

	c.DoneMessage(&models.Message{Chat: models.Chat{ID: 1}, ID: 1000})
	c.DoneMessage(&models.Message{Chat: models.Chat{ID: 10000}, ID: 1001})
	ch := c.GetChannels()
	require.Len(t, ch, 2)

	c.DoneMessage(&models.Message{Chat: models.Chat{ID: 2}, ID: 1002})
	ch = c.GetChannels()
	require.Len(t, ch, 3)
}

func TestKeeper_undone(t *testing.T) {
	c := bot.NewMessageKeeper()
	c.DoneMessage(&models.Message{Chat: models.Chat{ID: 1}, ID: 1000})
	c.DoneMessage(&models.Message{Chat: models.Chat{ID: 3}, ID: 1000})
	c.DoneMessage(&models.Message{Chat: models.Chat{ID: 3}, ID: 1001})
	ch := c.GetChannels()
	require.Len(t, ch, 2)

	c.UnDoneMessage(&models.Message{Chat: models.Chat{ID: 3}, ID: 1001})
	ch = c.GetChannels()
	require.Len(t, ch, 2)

	c.UnDoneMessage(&models.Message{Chat: models.Chat{ID: 3}, ID: 1000})
	ch = c.GetChannels()
	require.Len(t, ch, 1)

	c.UnDoneMessage(&models.Message{Chat: models.Chat{ID: 1}, ID: 776})
	ch = c.GetChannels()
	require.Len(t, ch, 1)

	c.UnDoneMessage(&models.Message{Chat: models.Chat{ID: 1}, ID: 1000})
	ch = c.GetChannels()
	require.Empty(t, ch)
}

func TestKeeper_messages(t *testing.T) {
	c := bot.NewMessageKeeper()
	c.DoneMessage(&models.Message{Chat: models.Chat{ID: 1}, ID: 1000})
	c.DoneMessage(&models.Message{Chat: models.Chat{ID: 1}, ID: 1001})
	ch := c.GetChannels()
	require.Equal(t, len(ch), 1)

	messages := c.GetMessages(1)
	require.Len(t, messages, 2)
	require.Contains(t, messages, 1000)
	require.Contains(t, messages, 1001)
}
