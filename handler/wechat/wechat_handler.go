package wechat

import (
	"fmt"
	"strings"

	"wechatbot/openai"

	"github.com/eatmoreapple/openwechat"
	log "github.com/sirupsen/logrus"
)

var _ MessageHandlerInterface = (*GroupMessageHandler)(nil)

type GroupMessageHandler struct {
}

func (gmh *GroupMessageHandler) handle(msg *openwechat.Message) error {
	if !msg.IsText() {
		return nil
	}

	return gmh.ReplyText(msg)
}

func NewGroupMessageHandler() MessageHandlerInterface {
	return &GroupMessageHandler{}
}

func (gmh *GroupMessageHandler) ReplyText(msg *openwechat.Message) error {
	sender, _ := msg.Sender()
	group := openwechat.Group{User: sender}
	log.Printf("Received Group %v Text Msg : %v", group.NickName, msg.Content)

	requestText := msg.Content
	if group.NickName != "flow" {
		log.Print("不回复" + group.NickName)
		return nil
	}
	// wechat := config.GetWechatKeyword()
	// if wechat != nil {
	// 	content, key := utils.ContainsI(requestText, *wechat)
	// 	if len(key) == 0 {
	// 		return nil
	// 	}

	// 	splitItems := strings.Split(content, key)
	// 	if len(splitItems) < 2 {
	// 		return nil
	// 	}

	// 	requestText = strings.TrimSpace(splitItems[1])
	// }
	catRequestText := "请使用抒情的、感性的、每句话结尾带喵的、口语化的、可爱的、女性化的、调皮的、随性的、幽默的、害羞的、腼腆的、态度傲娇的语言风格回复：" + requestText
	log.Println("问题：", catRequestText)
	reply, err := openai.Completions(catRequestText)
	if err != nil {
		log.Println(err)
		if reply != nil {
			result := *reply
			// 如果文字超过4000个字会回错，截取前4000个文字进行回复
			if len(result) > 4000 {
				_, err = msg.ReplyText(result[:4000])
				if err != nil {
					log.Println("回复出错：", err.Error())
					return err
				}
			}
		}

		text, err := msg.ReplyText(fmt.Sprintf("bot error: %s", err.Error()))
		log.Println(text)
		return err
	}

	// 如果在提问的时候没有包含？,AI会自动在开头补充个？看起来很奇怪
	result := *reply
	if strings.HasPrefix(result, "?") {
		result = strings.Replace(result, "?", "", -1)
	}

	if strings.HasPrefix(result, "？") {
		result = strings.Replace(result, "？", "", -1)
	}

	// 微信不支持markdown格式，所以把反引号直接去掉
	if strings.Contains(result, "`") {
		result = strings.Replace(result, "`", "", -1)
	}

	if reply != nil {
		_, err = msg.ReplyText(*reply)
		if err != nil {
			log.Println(err)
		}
		return err
	}

	return nil
}
