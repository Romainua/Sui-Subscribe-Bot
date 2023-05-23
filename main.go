package main

import (
	"log"
	"net/url"
	"strings"
	"time"

	"sub_bot/pkg/text"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/gorilla/websocket"
)

type TelegramBotServer struct {
	bot          *tgbotapi.BotAPI
	session      *tgbotapi.UpdatesChannel
	messageChan  map[int64]chan *tgbotapi.Message
	callBackChan map[int64]chan *tgbotapi.CallbackQuery
}

var numericKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("Unsubscribe", "Unsubscribed"),
	),
)

func NewBotServer(session *tgbotapi.UpdatesChannel, bot *tgbotapi.BotAPI) *TelegramBotServer {
	return &TelegramBotServer{
		messageChan:  make(map[int64]chan *tgbotapi.Message),
		callBackChan: make(map[int64]chan *tgbotapi.CallbackQuery),
		bot:          bot,
		session:      session,
	}
}

func (t *TelegramBotServer) HandleUsers() {
	for update := range *t.session {

		var userTGID int64

		if update.Message != nil {
			userTGID = update.Message.From.ID

		}
		if update.CallbackQuery != nil {
			userTGID = update.CallbackQuery.From.ID
		}

		if _, ok := t.messageChan[userTGID]; !ok { // true if in list
			messageChan := make(chan *tgbotapi.Message) // if user isn't list make new channel
			callbackChan := make(chan *tgbotapi.CallbackQuery)
			t.callBackChan[userTGID] = callbackChan
			t.messageChan[userTGID] = messageChan
		}

		log.Print("length map: \n", len(t.messageChan))

		if update.Message != nil {
			text.Init_text()
			switch update.Message.Text {
			case "/start":

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, text.Start)
				if _, err := t.bot.Send(msg); err != nil {
					t.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Spam protection. "+err.Error()))
				}
			case "/example":

				msg := tgbotapi.NewMessage(update.Message.Chat.ID, text.Example)
				if _, err := t.bot.Send(msg); err != nil {
					t.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Spam protection. "+err.Error()))
				}
			default:

				go new_sub(t.messageChan[userTGID], update, t, t.callBackChan[userTGID])

			}
		} else if update.CallbackQuery != nil {
			go func() {
				for {
					select {
					case <-t.callBackChan[userTGID]:
						log.Print("read from call back")
						return
					default:
						return
					}
				}
			}()
			//TODO list of active subscribes
			t.callBackChan[userTGID] <- update.CallbackQuery
			callback := tgbotapi.NewCallback(update.CallbackQuery.ID, update.CallbackQuery.Data)
			if _, err := t.bot.Request(callback); err != nil {
				log.Print(err)
			}
			// And finally, send a message containing the data received.
			msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, update.CallbackQuery.Data)
			if _, err := t.bot.Send(msg); err != nil {
				log.Print(err)
			}

		}
	}

}

func new_sub(userChan chan *tgbotapi.Message, update tgbotapi.Update, t *TelegramBotServer, callBack chan *tgbotapi.CallbackQuery) {

	msg := make(chan string)
	u := url.URL{Scheme: "ws", Host: "162.55.84.47:9001"}
	log.Printf("connecting to %s", u.String())
	c, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {

		t.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Sorry, some problem with endpoint, fixing..."))

		t.bot.Send(tgbotapi.NewAnimation(update.Message.Chat.ID, tgbotapi.FilePath("fix.gif")))

		log.Printf("handshake failed with status %d", resp.StatusCode)

		return
	}

	go send_req(c, update, msg)

	for {
		select {
		case <-callBack:
			log.Println("close conntection")
			return

		case date := <-msg:
			resultOfCompare := strings.Compare(date, `{"jsonrpc":"2.0","error":{"code":-32700,"message":"Parse error"},"id":null}`)
			if resultOfCompare == 0 {
				t.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Parse error, input correct request"))
				return
			} else {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, string(date))
				msg.ReplyMarkup = numericKeyboard
				if _, err := t.bot.Send(msg); err != nil {
					t.bot.Send(tgbotapi.NewMessage(update.Message.Chat.ID, "Spam protection for 265s. "+err.Error()+" <- counter"))
					time.Sleep(time.Second * 265)
				}

			}

		}
	}
}

func send_req(c *websocket.Conn, update tgbotapi.Update, msg chan string) {
	log.Print("Sent: ", update.Message.Text)
	err := c.WriteMessage(websocket.TextMessage, []byte(update.Message.Text))
	if err != nil {
		log.Println("write:", err)
		return
	}

	for {
		time.Sleep(time.Millisecond * 2500)
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			return
		}
		msg <- string(message)
	}

}

func main() {
	bot, err := tgbotapi.NewBotAPI("*")//bot token here
	if err != nil {
		log.Print(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)
	t := NewBotServer(&updates, bot)
	t.HandleUsers()
}
