package scheduler

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap"

	"github.com/cortezaproject/corteza-server/pkg/eventbus"
)

type (
	mockEvent struct {
		rType string
		eType string
		match func(matcher eventbus.ConstraintMatcher) bool
	}
)

func (e mockEvent) ResourceType() string {
	return e.rType
}

func (e mockEvent) EventType() string {
	return e.eType
}

func (e mockEvent) Match(matcher eventbus.ConstraintMatcher) bool {
	if e.match == nil {
		return true
	}

	return e.match(matcher)
}

func TestMainServiceFunctions(t *testing.T) {
	r := require.New(t)

	const (
		loopInterval = time.Millisecond
		actionWait   = loopInterval * 10
	)

	r.Nil(gScheduler)
	Setup(zap.NewNop(), eventbus.New(), time.Second)
	r.NotNil(gScheduler)
	r.False(gScheduler.Started())
	r.Equal(gScheduler, Service())
	Service().Start(context.Background())
	Service().OnTick(&mockEvent{}, &mockEvent{}, &mockEvent{}, &mockEvent{})
	time.Sleep(actionWait)
	r.True(gScheduler.Started())
	gScheduler.Stop()
	time.Sleep(actionWait)
	r.False(gScheduler.Started())
}
