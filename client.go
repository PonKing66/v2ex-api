package v2exapi

import (
	"strconv"
)

var prefix = "https://www.v2ex.com/api/v2/"

func (c *Client) Notifications(page int) ([]byte, error) {
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

	return body, nil
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

func (c *Client) Member() ([]byte, error) {
	body, err := c.request(prefix,
		"member",
		"GET",
		nil,
	)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) Token() ([]byte, error) {
	body, err := c.request(prefix,
		"token",
		"GET",
		nil,
	)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) NodesNodeName(nodeName string) ([]byte, error) {
	body, err := c.request(prefix,
		"nodes/"+nodeName,
		"GET",
		nil,
	)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) NodesNodeNameTopics(nodeName string, page int) ([]byte, error) {
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

	return body, nil
}

func (c *Client) Topics(topicId string) ([]byte, error) {
	body, err := c.request(prefix,
		"topics/"+topicId,
		"GET",
		nil,
	)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func (c *Client) TopicsReplies(topicId string, page int) ([]byte, error) {
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
	return body, nil
}
