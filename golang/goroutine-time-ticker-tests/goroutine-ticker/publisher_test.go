package goroutine_ticker_test

import (
	"testing"
	"time"

	"github.com/nycdavid/scratch/golang/goroutine-time-ticker-tests/goroutine-ticker"
	_assert "github.com/stretchr/testify/assert"
)

type (
	mockClock struct{}
)

func (m *mockClock) C() <-chan time.Time {
}

func (m *mockClock) Stop() {

}

func Test_PeriodicPublish(t *testing.T) {
	assert := _assert.New(t)

	printer := &goroutine_ticker.Printer{}
	printer.Start()

	assert.Equal(printer.Iteration, 5)
}
