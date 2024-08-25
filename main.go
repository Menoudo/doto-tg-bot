package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
)

var massageCount int = 0

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithMessageTextHandler("/stat", bot.MatchTypeExact, stat_handler),
		bot.WithMessageTextHandler("/start", bot.MatchTypeExact, start_handler),
		bot.WithCallbackQueryDataHandler("done", bot.MatchTypeContains, callbackHandler),
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(os.Getenv("TG_BOT_TOKEN"), opts...)
	if err != nil {
		panic(err)
	}

	b.Start(ctx)
}

func GetTodoKeyboard(isDone bool) *models.InlineKeyboardMarkup {
	kb := &models.InlineKeyboardMarkup{
		InlineKeyboard: [][]models.InlineKeyboardButton{
			{
				{Text: "✔️ Выполнить", CallbackData: "done"},
			},
		},
	}
	if isDone {
		kb.InlineKeyboard[0] = []models.InlineKeyboardButton{
			{Text: "Выполнено", CallbackData: "undone"},
		}
	}
	return kb
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.ChannelPost == nil {
		return
	}
	massageCount++
	origMessage := update.ChannelPost
	logger.Print(origMessage.Text)

	_, errEdit := b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      origMessage.Chat.ID,
		MessageID:   origMessage.ID,
		Text:        origMessage.Text,
		ReplyMarkup: GetTodoKeyboard(false),
	})
	if errEdit != nil {
		logger.Printf("error edit message: %v", errEdit)
		return
	}
}

func start_handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:    update.Message.Chat.ID,
		Text:      bot_information,
		ParseMode: models.ParseModeMarkdown,
	})
	if err != nil {
		logger.Printf("error edit message: %v", err)
		return
	}
}

func stat_handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   fmt.Sprintf("Message processed: %d", massageCount),
	})
	if err != nil {
		logger.Printf("error edit message: %v", err)
		return
	}
}

func callbackHandler(ctx context.Context, b *bot.Bot, update *models.Update) {
	_, err := b.AnswerCallbackQuery(ctx, &bot.AnswerCallbackQueryParams{
		CallbackQueryID: update.CallbackQuery.ID,
		ShowAlert:       false,
	})
	if err != nil {
		logger.Printf("error AnswerCallbackQuery: %v", err)
		return
	}

	origMessage := update.CallbackQuery.Message.Message

	isDone := update.CallbackQuery.Data == "done"
	user := update.CallbackQuery.From.Username

	_, errEdit := b.EditMessageText(ctx, &bot.EditMessageTextParams{
		ChatID:      origMessage.Chat.ID,
		MessageID:   origMessage.ID,
		Text:        UpdateTest(origMessage.Text, user),
		ParseMode:   models.ParseModeHTML,
		ReplyMarkup: GetTodoKeyboard(isDone),
	})
	if errEdit != nil {
		logger.Printf("error edit message: %v", errEdit)
		return
	}
}
