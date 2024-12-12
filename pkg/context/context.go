package appcontext

import "context"

type AppContext struct {
	ctx context.Context
}

func NewAppContext() *AppContext {
	return &AppContext{
		ctx: context.Background(),
	}
}

func (a *AppContext) Context() context.Context {
	return a.ctx
}

func (a *AppContext) SetContext(ctx context.Context) {
	a.ctx = ctx
}

func (a *AppContext) Set(key, value interface{}) {
	a.ctx = context.WithValue(a.ctx, key, value)
}

func (a *AppContext) Get(key interface{}) interface{} {
	return a.ctx.Value(key)
}
