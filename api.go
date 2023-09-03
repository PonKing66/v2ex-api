package v2exapi

import (
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

type Client struct {
	debug  bool
	client *http.Client
	logger *log.Logger
	token  string
}

type ClientSetting struct {
	Client *http.Client
	// Debug
	DebugMode bool
	// Logger Prefix
	Prefix string
	// v2ex token
	Token string
}

func newClient(setting *ClientSetting) *Client {
	client := setting.Client

	if client == nil {
		client = http.DefaultClient
	}

	return &Client{
		debug:  setting.DebugMode,
		client: client,
		token:  setting.Token,
		logger: log.New(os.Stdout, setting.Prefix, log.LstdFlags),
	}
}

func (h *Client) doRequest(req *http.Request, v interface{}) ([]byte, error) {
	resp, err := h.client.Do(req)
	if err != nil {
		return nil, err
	}
	resp.Close = true
	defer resp.Body.Close()

	raw, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if h.debug {
		h.logger.Printf("%s %s  %+v", req.Method, req.URL, v)
		h.logger.Printf("%s", string(raw))
	}

	return raw, nil
}

func (h *Client) request(prefix, suffix, method string, payload map[string]string) ([]byte, error) {
	var (
		req *http.Request
		err error
	)

	data := url.Values{}
	for k, v := range payload {
		data.Add(k, v)
	}

	link := prefix + suffix
	switch method {
	case http.MethodDelete:
		if req, err = http.NewRequest(method, link, nil); err != nil {
			return nil, err
		}
		req.URL.RawQuery = data.Encode()
	case http.MethodGet:
		if req, err = http.NewRequest(method, link, nil); err != nil {
			return nil, err
		}
		req.URL.RawQuery = data.Encode()
	case http.MethodPost:
		if req, err = http.NewRequest(method, link, strings.NewReader(data.Encode())); err != nil {
			return nil, err
		}
	}

	req.Header.Add("Authorization", "Bearer "+h.token)

	return h.doRequest(req, payload)
}
