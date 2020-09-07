package ext

import (
	"log"

	tg "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/robfig/cron"
)

// InitializeScheduler starts new cron instance with predifined schedules
func InitializeScheduler(bot *tg.BotAPI) *cron.Cron {
	c := cron.New()
	c.AddFunc("0 9 * * 1-5", func() { SendIncompleteTasks(bot) })
	c.AddFunc("0 14 * * 1-5", func() { SendIncompleteTasks(bot) })
	c.AddFunc("0 19 * * 1-5", func() { SendIncompleteTasks(bot) })

	c.Start()
	log.Print("Scheduler initialized")
	return c
}
