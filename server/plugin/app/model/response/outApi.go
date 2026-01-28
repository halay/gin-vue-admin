package response

type (
	ImgResult struct {
		Index    int    `json:"index"`
		Url      string `json:"url"`
		Object   string `json:"object"`
		Duration string `json:"duration"`
		Error    string `json:"error,omitempty"`
	}

	CozeExecute struct {
		ID       string `json:"id"`
		Request  string `json:"request,omitempty"`
		Output   string `json:"output"`
		Error    string `json:"error"`
		Duration string `json:"duration"`
	}

	CozeResult struct {
		State   string        `json:"state"`
		Results []CozeExecute `json:"results"`
		Error   string        `json:"error"`
		Url     string        `json:"url"`
	}
)
