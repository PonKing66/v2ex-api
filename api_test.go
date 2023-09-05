package v2exapi

import (
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
	DebugMode: false,
}

func TestNotifications(t *testing.T) {
	v := newClient(DefaultSetting)
	json, _ := v.Notifications(1)
	t.Logf("Json:%v", json)
}

func TestDeleteNotifications(t *testing.T) {
	v := newClient(DefaultSetting)
	json, _ := v.DeleteNotifications("test123123")
	t.Logf("Json:%v", json)
}

func TestMember(t *testing.T) {
	v := newClient(DefaultSetting)
	res, _ := v.Member()
	t.Logf("Json:%v", res)
}

func TestToken(t *testing.T) {
	v := newClient(DefaultSetting)
	json, _ := v.Token()
	t.Logf("Json:%v", *json)
}

func TestNodesNodeName(t *testing.T) {
	v := newClient(DefaultSetting)
	json, _ := v.NodesNodeName("jobs")
	t.Logf("Json:%v", *json)
}

func TestNodesNodeNameTopics(t *testing.T) {
	v := newClient(DefaultSetting)
	json, _ := v.NodesNodeNameTopics("jobs", 2)
	t.Logf("Json:%v", *json)
}

func TestTopics(t *testing.T) {
	v := newClient(DefaultSetting)
	json, _ := v.Topics("970564")
	t.Logf("Json:%v", *json)
}

func TestTopicsReplies(t *testing.T) {
	v := newClient(DefaultSetting)
	json, _ := v.TopicsReplies("970564", 1)
	t.Logf("Json:%v", *json)
}

func unicode(raw []byte) []byte {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		panic(err)
	}
	return []byte(str)
}
