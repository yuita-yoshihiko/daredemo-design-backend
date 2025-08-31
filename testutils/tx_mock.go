package testutils

import "context"

type MockTxManager struct{}

func (m *MockTxManager) DoInTx(ctx context.Context, f func(context.Context) (any, error)) (any, error) {
	return f(ctx)
}
