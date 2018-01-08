package log

import (
	"encoding/json"
	"os"
)

type Stdout struct {
	Color bool `json:"color"`
}

func (std *Stdout) Init(config ...string) error {
	var c string
	if len(config) == 0 {
		c = `{"color":true}`
	} else {
		c = config[0]
	}
	return json.Unmarshal([]byte(c), std)
}

func (std Stdout) Write(p []byte) (int, error) {
	if std.Color {

	}
	return os.Stdout.Write(p)
}
