package entities

type ChatReq struct {
	Messages []*LlamaBody `json:"messages"`
}

type LlamaBody struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type LlamaResp struct {
	Response string `json:"response"`
}
