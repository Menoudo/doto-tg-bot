package main_test

import (
	"testing"

	bot "todo-tg-bot"

	"github.com/stretchr/testify/require"
)

func TestKeeper_empty(t *testing.T) {
	c := bot.New()
	ch := c.GetChannels()
	require.Empty(t, ch)
}

func TestKeeper_one(t *testing.T) {
	c := bot.New()
	c.DoneMessage("channel_1", 1000)
	ch := c.GetChannels()
	require.Len(t, ch, 1)
	c.DoneMessage("channel_1", 1001)
	require.Len(t, ch, 1)
}

func TestKeeper_many(t *testing.T) {
	c := bot.New()
	c.DoneMessage("channel_1", 1000)
	c.DoneMessage(10000, 1001)
	ch := c.GetChannels()
	require.Len(t, ch, 2)
	c.DoneMessage("channel_2", 1002)
	ch = c.GetChannels()
	require.Len(t, ch, 3)
}

func TestKeeper_undone(t *testing.T) {
	c := bot.New()
	c.DoneMessage("channel_1", 1000)
	c.DoneMessage("channel_3", 1000)
	c.DoneMessage("channel_3", 1001)
	ch := c.GetChannels()
	require.Len(t, ch, 2)

	c.UnDoneMessage("channel_3", 1001)
	ch = c.GetChannels()
	require.Len(t, ch, 2)

	c.UnDoneMessage("channel_3", 1000)
	ch = c.GetChannels()
	require.Len(t, ch, 1)

	c.UnDoneMessage("channel_1", 776)
	ch = c.GetChannels()
	require.Len(t, ch, 1)

	c.UnDoneMessage("channel_1", 1000)
	ch = c.GetChannels()
	require.Empty(t, ch)
}

func TestKeeper_messages(t *testing.T) {
	c := bot.New()
	c.DoneMessage("channel_1", 1000)
	c.DoneMessage("channel_1", 1001)
	ch := c.GetChannels()
	require.Equal(t, len(ch), 1)

	messages := c.GetMessages("channel_1")
	require.Len(t, messages, 2)
	require.Contains(t, messages, 1000)
	require.Contains(t, messages, 1001)
}
