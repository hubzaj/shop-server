package runner

import (
	"context"
	"sync"
)

type ShopServerContext struct {
	wg  *sync.WaitGroup
	ctx context.Context
}

func NewShopServiceContext(ctx context.Context) *ShopServerContext {
	return &ShopServerContext{
		wg:  &sync.WaitGroup{},
		ctx: ctx,
	}
}

func (cwg *ShopServerContext) AddTask() {
	cwg.wg.Add(1)
}

func (cwg *ShopServerContext) TaskDone() {
	cwg.wg.Done()
}

func (cwg *ShopServerContext) Wait() {
	cwg.wg.Wait()
}
