package v2exapi

import (
	"github.com/tidwall/gjson"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"testing"
)

var f = func() string {
	file, err := os.ReadFile("token.txt")
	if err != nil {
		panic(err)
	}
	return string(file)
}

var proxyURL, _ = url.Parse("http://127.0.0.1:7890")

var proxy = http.ProxyURL(proxyURL)
var transport = &http.Transport{Proxy: proxy}
var client = &http.Client{Transport: transport}

var DefaultSetting = &ClientSetting{
	Token:     f(),
	Client:    client,
	DebugMode: true,
}

func TestNotifications(t *testing.T) {
	v := newClient(DefaultSetting)
	json, _ := v.Notifications(1)
	t.Logf("Json:%v", gjson.ParseBytes(json).String())
}

func TestDeleteNotifications(t *testing.T) {
	v := newClient(DefaultSetting)
	json, _ := v.DeleteNotifications("test123123")
	t.Logf("Json:%v", gjson.ParseBytes(json).String())
}

func TestMember(t *testing.T) {
	v := newClient(DefaultSetting)
	json, _ := v.Member()
	t.Logf("Json:%v", gjson.ParseBytes(json).String())
}

func TestToken(t *testing.T) {
	v := newClient(DefaultSetting)
	json, _ := v.Token()
	t.Logf("Json:%v", gjson.ParseBytes(json).String())
}

func TestNodesNodeName(t *testing.T) {
	v := newClient(DefaultSetting)
	json, _ := v.NodesNodeName("jobs")
	t.Logf("Json:%v", gjson.ParseBytes(json).String())
}

func TestNodesNodeNameTopics(t *testing.T) {
	v := newClient(DefaultSetting)
	json, _ := v.NodesNodeNameTopics("jobs", 2)
	t.Logf("Json:%v", gjson.ParseBytes(json).String())
}

func TestTopics(t *testing.T) {
	v := newClient(DefaultSetting)
	json, _ := v.Topics("970564")
	t.Logf("Json:%v", gjson.ParseBytes(json).String())
}

func TestTopicsReplies(t *testing.T) {
	v := newClient(DefaultSetting)
	json, _ := v.TopicsReplies("970564", 1)
	t.Logf("Json:%v", gjson.ParseBytes(json).String())
}

func unicode(raw []byte) []byte {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		panic(err)
	}
	return []byte(str)
}
