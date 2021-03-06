// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/line/line-bot-sdk-go/linebot"
)

var bot *linebot.Client

func main() {
	var err error
	bot, err = linebot.New(os.Getenv("ChannelSecret"), os.Getenv("ChannelAccessToken"))
	log.Println("Bot:", bot, " err:", err)	//在log中加入紀錄
	http.HandleFunc("/callback", callbackHandler)
	port := os.Getenv("PORT")
	addr := fmt.Sprintf(":%s", port)
	http.ListenAndServe(addr, nil)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
	events, err := bot.ParseRequest(r)

	if err != nil {
		if err == linebot.ErrInvalidSignature {
			w.WriteHeader(400)
		} else {
			w.WriteHeader(500)
		}
		return
	}



	for _, event := range events {
		
		if event.Type == linebot.EventTypeMessage {
			switch  message := event.Message.(type){
				
			case *linebot.TextMessage:
				msgcontent := message.Text+" 您好!"
// 				if message.Text == "幹" {
// 					err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text+" 幹!")).Do()
// 				}
// 				else {
// 					err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(message.Text+" OK!")).Do()
// 				}
// 				if err != nil {
// 					log.Print(err)
// 				}
				if message.Text == "幹" {
					msgcontent = "牟利洗列幹三小啦！"
				}
				_, err = bot.ReplyMessage(event.ReplyToken, linebot.NewTextMessage(msgcontent)).Do();
				if err != nil {
					log.Print(err)
				}
			}
		}
	}
}
