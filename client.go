package v2exapi

import (
	"encoding/json"
	"strconv"
)

var prefix = "https://www.v2ex.com/api/v2/"

func (c *Client) Notifications(page int) (*Notifications, error) {
	body, err := c.request(prefix,
		"notifications",
		"GET",
		map[string]string{
			"p": strconv.Itoa(page),
		},
	)
	if err != nil {
		return nil, err
	}
	var res *Notifications
	err = json.Unmarshal(body, &res)
	return res, err
}

func (c *Client) DeleteNotifications(notificationId string) ([]byte, error) {
	body, err := c.request(prefix,
		"notifications/"+notificationId,
		"DELETE",
		nil,
	)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) Member() (*Member, error) {
	body, err := c.request(prefix,
		"member",
		"GET",
		nil,
	)
	if err != nil {
		return nil, err
	}

	var res *Member
	err = json.Unmarshal(body, &res)

	return res, nil
}

func (c *Client) Token() (*Token, error) {
	body, err := c.request(prefix,
		"token",
		"GET",
		nil,
	)
	if err != nil {
		return nil, err
	}
	var res *Token
	err = json.Unmarshal(body, &res)

	return res, nil
}

func (c *Client) NodesNodeName(nodeName string) (*NodesNodeName, error) {
	body, err := c.request(prefix,
		"nodes/"+nodeName,
		"GET",
		nil,
	)
	if err != nil {
		return nil, err
	}
	var res *NodesNodeName
	err = json.Unmarshal(body, &res)

	return res, nil
}

func (c *Client) NodesNodeNameTopics(nodeName string, page int) (*NodesNodeNameTopics, error) {
	body, err := c.request(prefix,
		"nodes/"+nodeName+"/topics",
		"GET",
		map[string]string{
			"p": strconv.Itoa(page),
		},
	)
	if err != nil {
		return nil, err
	}

	var res *NodesNodeNameTopics
	err = json.Unmarshal(body, &res)

	return res, nil
}

func (c *Client) Topics(topicId string) (*Topics, error) {
	body, err := c.request(prefix,
		"topics/"+topicId,
		"GET",
		nil,
	)
	if err != nil {
		return nil, err
	}

	var res *Topics
	err = json.Unmarshal(body, &res)

	return res, nil
}

func (c *Client) TopicsReplies(topicId string, page int) (*TopicsReplies, error) {
	body, err := c.request(prefix,
		"topics/"+topicId+"/replies",
		"GET",
		map[string]string{
			"p": strconv.Itoa(page),
		},
	)
	if err != nil {
		return nil, err
	}
	var res *TopicsReplies
	err = json.Unmarshal(body, &res)

	return res, nil
}
