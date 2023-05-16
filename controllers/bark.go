package controllers

import (
	"PrometheusAlert/models"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type BarkMessage struct {
	Title             string `json:"title"`             // Notification title (font size would be larger than the body)
	Body              string `json:"body"`              // Notification content
	Category          string `json:"category"`          // Reserved field, no use yet
	DeviceKey         string `json:"device_key"`        // The key for each device
	Level             string `json:"level"`             // 	'active', 'timeSensitive', or 'passive'
	Badge             int    `json:"badge"`             // The number displayed next to App icon (Apple Developer)
	AutomaticallyCopy string `json:"automaticallyCopy"` // 	Must be 1
	Copy              string `json:"copy"`              // The value to be copied
	Sound             string `json:"sound"`             // Value from here
	Icon              string `json:"icon"`              // An url to the icon, available only on iOS 15 or later
	Group             string `json:"group"`             // The group of the notification
	IsArchive         string `json:"isArchive"`         // Value must be 1. Whether or not should be archived by the app
	Url               string `json:"url"`               //  that will jump when click notification
}

// SendBark 发送消息至iPhone
func SendBark(title, msg, AtSomeOne, logsign string) string {
	open := beego.AppConfig.String("open-bark")
	if open != "1" {
		logs.Info(logsign, "[bark]", "bark未配置未开启状态,请先配置open-bark为1")
		return "bark未配置未开启状态,请先配置open-bark为1"
	}
	barkServer := beego.AppConfig.String("BARK_URL")
	if barkServer == "" {
		barkServer = "https://api.day.app/push"
	}
	sendUser := beego.AppConfig.String("BARK_KEYS")
	if len(AtSomeOne) > 0 {
		sendUser = AtSomeOne
	}
	barkTitle := beego.AppConfig.String("BARK_TITLE")
	if len(title) > 0 {
		barkTitle = title
	}
	if len(barkTitle) == 0 {
		barkTitle = "Bark推送测试"
	}
	barkCopy := beego.AppConfig.String("BARK_COPY")
	barkArchive := beego.AppConfig.String("BARK_ARCHIVE")
	barkGroup := beego.AppConfig.String("BARK_GROUP")
	bm := BarkMessage{
		Title:             barkTitle,
		Body:              msg,
		AutomaticallyCopy: barkCopy,
		IsArchive:         barkArchive,
		Group:             barkGroup,
	}
	sendUsers := strings.Split(sendUser, "-")
	for _, u := range sendUsers {
		// 处理发送消息
		bm.DeviceKey = u
		b := new(bytes.Buffer)
		json.NewEncoder(b).Encode(bm)
		logs.Info(logsign, "[bark]", b)
		get, err := sendBark(barkServer, *b)
		if err != nil {
			logs.Error(logsign, "[bark]", fmt.Errorf("send to %s, err: %v", u, err))
		}
		if get.Code != 200 {
			logs.Error(logsign, "[bark]", fmt.Errorf("send to %s, get code: %d", u, get.Code))
		}
	}
	models.AlertToCounter.WithLabelValues("bark").Add(1)
	ChartsJson.Bark += 1
	logs.Info(logsign, "[bark]", "bark send ok.")
	return "bark send ok"
}

type responseMessage struct {
	Code    int64  `json:"code"`
	Data    string `json:"data"`
	Message string `json:"message"`
}

func sendBark(url string, body bytes.Buffer) (responseMessage, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Post(url, "application/json", &body)
	if err != nil {
		return responseMessage{}, err
	}
	defer resp.Body.Close()
	var buffer [512]byte
	result := bytes.NewBuffer(nil)
	for {
		n, err := resp.Body.Read(buffer[0:])
		result.Write(buffer[0:n])
		if err != nil && err == io.EOF {
			break
		} else if err != nil {
			return responseMessage{}, err
		}
	}

	message := responseMessage{}
	err = json.Unmarshal(result.Bytes(), &message)
	if err != nil {
		return responseMessage{}, err
	}

	return message, nil
}

func generateGetUrlPrefix(title, msg, userkey string) string {
	barkserver := beego.AppConfig.String("BARK_URL")
	barktitle := beego.AppConfig.String("BARK_TITLE")
	if len(title) > 0 {
		barktitle = title
	}
	if len(barktitle) == 0 {
		barktitle = "Bark推送测试"
	}
	buffer := bytes.Buffer{}
	buffer.WriteString(fmt.Sprintf("%s/%s/%s/%s", barkserver, userkey, barktitle, msg))
	return buffer.String()
}
