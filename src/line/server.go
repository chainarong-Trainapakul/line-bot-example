package line

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/line/line-bot-sdk-go/linebot"
	"github.com/line/line-bot-sdk-go/linebot/httphandler"
	"sql" //sql wrapper
)

type Server struct {
	db *sql.SQLdb
	lc *lineConfig
}

func NewServer(db *sql.SQLdb, lc *lineConfig) *Server {
	return &Server{
		db: db,
		lc: lc,
	}
}

func (s *Server) Run() {
	handler, err := httphandler.New(s.lc.chanSecret, s.lc.chanToken)
	if err != nil {
		log.Fatalln(err)
	}
	err = s.db.Connect()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("start server !!!")
	handler.HandleEvents(func(events []*linebot.Event, r *http.Request) {
		bot, err := handler.NewClient()
		if err != nil {
			log.Print(err)
			return
		}
		for _, event := range events {
			fmt.Println(event)
			if event.Type == linebot.EventTypeMessage {
				switch message := event.Message.(type) {
				case *linebot.TextMessage:
					msg := s.checkTextMessage(message.Text)
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(msg)).Do(); err != nil {
						//if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text)).Do(); err != nil {
						log.Print(err)
					}
				case *linebot.ImageMessage:
					// not implement yet
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("ImageMessage")).Do(); err != nil {
						log.Print(err)
					}
				case *linebot.VideoMessage:
					// not implement yet
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("VideoMessage")).Do(); err != nil {
						log.Print(err)
					}
				case *linebot.AudioMessage:
					// not implement yet
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("AudioMessage")).Do(); err != nil {
						log.Print(err)
					}
				case *linebot.FileMessage:
					// not implement yet
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("FileMessage")).Do(); err != nil {
						log.Print(err)
					}
				case *linebot.LocationMessage:
					// not implement yet
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("LocationMessage")).Do(); err != nil {
						log.Print(err)
					}
				case *linebot.StickerMessage:
					// not implement yet
					if _, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage("StickerMessage")).Do(); err != nil {
						log.Print(err)
					}
				}
			}
		}
	})
	http.Handle("/callback", handler)
	if err := http.ListenAndServe(":6969", nil); err != nil {
		log.Fatalln(err)
	}
}

func (s *Server) checkTextMessage(msg string) string {
	msgSp := strings.Split(msg, " ")
	if msgSp[0] == "-Profile" {
		switch strings.Replace(msgSp[1], " ", "", 1) {
		case "add":
			return s.addProfile(msgSp)
		case "mod":
			return s.modifyProfile(msgSp)
		case "del":
			return s.deleteProfile(msgSp)
		case "search":
			return s.searchProfile(msgSp)
		default:
			return helpProfile(msgSp)
		}
	}
	return "unknown command"
}

func (s *Server) addProfile(msg []string) string {
	if len(msg) == 4 {
		err := s.db.Add(strings.Replace(msg[2], " ", "", -1), strings.Replace(msg[3], " ", "", -1))
		if err != nil {
			log.Print(err)
			return "add error please try again"
		}
		return "add success:" + msg[2] + " " + msg[3]
	}
	return addHelpProfile()
}

func addHelpProfile() string {
	return "-Profile add 'name' 'surname'\nPlease check format and try again"
}

func (s *Server) deleteProfile(msg []string) string {
	if len(msg) == 4 {
		err := s.db.Delete(strings.Replace(msg[3], " ", "", -1))
		if err != nil {
			return "delete error\nPlease try again"
		}
		return "delete id: " + msg[3] + " success"
	}
	return deleteHelpProfile()
}

func deleteHelpProfile() string {
	return "-Profile delete id 'your-id' \nPlease check format and try again"
}

func (s *Server) modifyProfile(msg []string) string {
	return modifyHelpProfile()
}

func modifyHelpProfile() string {
	return "not implement yet"
}

func helpProfile(msg []string) string {
	return "-Profile add|del|search"
}

func (s *Server) searchProfile(msg []string) string {
	var c string
	if len(msg) == 4 {
		iden := strings.Replace(msg[2], " ", "", -1)
		if iden == "name" {
			c = "select * from info where name = '" + msg[3] + "'"
		} else if iden == "surname" {
			c = "select * from info where surname = '" + msg[3] + "'"
		} else {
			c = "select * from info where id = " + msg[3]
		}
		rs, err := s.db.SelectQuery(c)
		if err != nil {
			return "query err \n Please try again"
		}
		return rs
	}
	return searchHelpProfile()
}

func searchHelpProfile() string {
	return "-Profile search id|name|surname 'yourname|yoursurname'\nPlease check format and try again"
}
