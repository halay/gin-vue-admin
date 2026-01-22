package response

import "time"

type (
	ImgResult struct {
		Index    int           `json:"index"`
		Url      string        `json:"url"`
		Duration time.Duration `json:"duration"`
		Error    string        `json:"error,omitempty"`
	}

	CozeExecute struct {
		ID       string        `json:"id"`
		Output   string        `json:"output"`
		Error    string        `json:"error"`
		Duration time.Duration `json:"duration"`
	}

	CozeTask struct {
		ID       string        `json:"id"`
		Error    string        `json:"error"`
		Output   string        `json:"output"`
		Duration time.Duration `json:"duration"`
	}

	CozeResult struct {
		State   string      `json:"state"`
		Execute CozeExecute `json:"execute"`
		Task    CozeTask    `json:"task"`
		Url     string      `json:"url"`
	}
)
