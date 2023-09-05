package v2exapi

/**
  {
     "id": 20639167,
     "member_id": 627282,
     "for_member_id": 558600,
     "text": "<a href=\"/member/wanderingaround\"><strong>wanderingaround</strong></a> \u6536\u85cf\u4e86\u4f60\u53d1\u5e03\u7684\u4e3b\u9898 \u203a <a href=\"/t/923238#reply82\" class=\"topic-link\">\u71ac\u4e0d\u4f4f\u4e86\uff0c\u60f3\u8f9e\u804c\u4e86</a>",
     "payload": null,
     "payload_rendered": "",
     "created": 1693659177,
     "member": {
       "username": "wanderingaround"
     }
   },
*/

type Notifications struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  []struct {
		Id              int64  `json:"id"`
		MemberId        int    `json:"member_id"`
		ForMemberId     int    `json:"for_member_id"`
		Text            string `json:"text"`
		Payload         string `json:"payload"`
		PayloadRendered string `json:"payload_rendered"`
		Created         int64  `json:"created"`
		Member          struct {
			UserName string `json:"user_name"`
		} `json:"member"`
	} `json:"result"`
}

type Member struct {
	Success bool `json:"success"`
	Result  struct {
		ID           int         `json:"id"`
		Username     string      `json:"username"`
		URL          string      `json:"url"`
		Website      string      `json:"website"`
		Twitter      interface{} `json:"twitter"`
		Psn          interface{} `json:"psn"`
		Github       interface{} `json:"github"`
		Btc          interface{} `json:"btc"`
		Location     string      `json:"location"`
		Tagline      string      `json:"tagline"`
		Bio          string      `json:"bio"`
		AvatarMini   string      `json:"avatar_mini"`
		AvatarNormal string      `json:"avatar_normal"`
		AvatarLarge  string      `json:"avatar_large"`
		Created      int         `json:"created"`
		LastModified int         `json:"last_modified"`
	} `json:"result"`
}

type Token struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  struct {
		Token       string `json:"token"`
		Scope       string `json:"scope"`
		Expiration  int    `json:"expiration"`
		GoodForDays int    `json:"good_for_days"`
		TotalUsed   int    `json:"total_used"`
		LastUsed    int    `json:"last_used"`
		Created     int    `json:"created"`
	} `json:"result"`
}

type NodesNodeName struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  struct {
		ID           int    `json:"id"`
		URL          string `json:"url"`
		Name         string `json:"name"`
		Title        string `json:"title"`
		Header       string `json:"header"`
		Footer       string `json:"footer"`
		Avatar       string `json:"avatar"`
		Topics       int    `json:"topics"`
		Created      int    `json:"created"`
		LastModified int    `json:"last_modified"`
	} `json:"result"`
}

type NodesNodeNameTopics struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  []struct {
		ID              int    `json:"id"`
		Title           string `json:"title"`
		Content         string `json:"content"`
		ContentRendered string `json:"content_rendered"`
		Syntax          int    `json:"syntax"`
		URL             string `json:"url"`
		Replies         int    `json:"replies"`
		LastReplyBy     string `json:"last_reply_by"`
		Created         int    `json:"created"`
		LastModified    int    `json:"last_modified"`
		LastTouched     int    `json:"last_touched"`
	} `json:"result"`
}

type Topics struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  struct {
		ID              int    `json:"id"`
		Title           string `json:"title"`
		Content         string `json:"content"`
		ContentRendered string `json:"content_rendered"`
		Syntax          int    `json:"syntax"`
		URL             string `json:"url"`
		Replies         int    `json:"replies"`
		LastReplyBy     string `json:"last_reply_by"`
		Created         int    `json:"created"`
		LastModified    int    `json:"last_modified"`
		LastTouched     int    `json:"last_touched"`
		Member          struct {
			ID       int         `json:"id"`
			Username string      `json:"username"`
			Bio      string      `json:"bio"`
			Website  string      `json:"website"`
			Github   interface{} `json:"github"`
			URL      string      `json:"url"`
			Avatar   string      `json:"avatar"`
			Created  int         `json:"created"`
		} `json:"member"`
		Node struct {
			ID           int    `json:"id"`
			URL          string `json:"url"`
			Name         string `json:"name"`
			Title        string `json:"title"`
			Header       string `json:"header"`
			Footer       string `json:"footer"`
			Avatar       string `json:"avatar"`
			Topics       int    `json:"topics"`
			Created      int    `json:"created"`
			LastModified int    `json:"last_modified"`
		} `json:"node"`
		Supplements []interface{} `json:"supplements"`
	} `json:"result"`
}

type TopicsReplies struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Result  []struct {
		ID              int    `json:"id"`
		Content         string `json:"content"`
		ContentRendered string `json:"content_rendered"`
		Created         int    `json:"created"`
		Member          struct {
			ID       int    `json:"id"`
			Username string `json:"username"`
			Bio      string `json:"bio"`
			Website  string `json:"website"`
			Github   string `json:"github"`
			URL      string `json:"url"`
			Avatar   string `json:"avatar"`
			Created  int    `json:"created"`
		} `json:"member"`
	} `json:"result"`
}
