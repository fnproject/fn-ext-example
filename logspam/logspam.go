package logspam

import (
	"context"
	"fmt"

	"github.com/fnproject/fn/api/models"
)

// This implements the extender.CallListener interface
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
