package vendors

import (
	"bytes"
)

type Vendor interface {
	GetName() string
	IsConfigured() bool
	Configure() error
	ListModels() ([]string, error)
	SendStream([]*common.Message, *common.ChatOptions, chan string) error
	Send([]*common.Message, *common.ChatOptions) (string, error)
	Setup() error
	SetupFillEnvFileContent(*bytes.Buffer)
}
