package logspam

import (
	"context"
	"fmt"

	"github.com/fnproject/fn/api/models"
	"github.com/fnproject/fn/api/server"
	"github.com/fnproject/fn/fnext"
)

func init() {
	server.RegisterExtension(&spamExt{})
}

type spamExt struct {
}

func (e *spamExt) Name() string {
	return "github.com/treeder/fn-ext-example/logspam"
}

func (e *spamExt) Setup(s fnext.ExtServer) error {
	s.AddCallListener(&LogSpam{})
	return nil
}

type LogSpam struct {
}

func (l *LogSpam) BeforeCall(ctx context.Context, call *models.Call) error {
	fmt.Println("YO! This is an annoying message that will happen every time a function is called.")
	return nil
}

func (l *LogSpam) AfterCall(ctx context.Context, call *models.Call) error {
	fmt.Println("YO! And this is an annoying message that will happen AFTER every time a function is called.")
	return nil
}
