package ext

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/adlio/trello"
	tg "github.com/go-telegram-bot-api/telegram-bot-api"
)

// GetChecklist returns checklist instance by ID
func GetChecklist(c *trello.Client, checklistID string, args trello.Arguments) (*trello.Checklist, error) {
	path := fmt.Sprintf("checklists/%s", checklistID)
	checklist := trello.Checklist{}
	err := c.Get(path, trello.Defaults(), &checklist)
	if err != nil {
		panic(err)
	}
	return &checklist, err
}

// FilterCompleteTasks returns array of task names with status "incomplete"
func FilterCompleteTasks(tasks []trello.CheckItem) (result []string) {
	for _, item := range tasks {
		if item.State == "incomplete" {
			result = append(result, item.Name)
		}
	}
	return
}

// GetUncompletedTasks returns text message with incomplete tasks
func GetUncompletedTasks(checklistID string) (composedMessage string) {
	token := os.Getenv("TRELLO_TOKEN")
	apiKey := os.Getenv("TRELLO_APP_KEY")
	client := trello.NewClient(apiKey, token)

	checklist, err := GetChecklist(client, checklistID, trello.Defaults())
	if err != nil {
		log.Fatal("Couldn't retrieve Checklist")
	}

	incompleteTasks := FilterCompleteTasks(checklist.CheckItems)
	message := strings.Join(incompleteTasks, "\n")

	return "List of incomplete tasks:\n\n" + message
}

// SendIncompleteTasks sends incomplete task list to bot
func SendIncompleteTasks(bot *tg.BotAPI) {
	tasksID := os.Getenv("TASKS_ID")
	ownerID, _ := strconv.ParseInt(os.Getenv("OWNER_ID"), 10, 64)
	incompleteTasks := GetUncompletedTasks(tasksID)
	msg := tg.NewMessage(ownerID, incompleteTasks)
	bot.Send(msg)
	log.Print("Sent message with incomplete task list")
}
