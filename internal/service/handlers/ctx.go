package handlers

import (
	"context"
	"net/http"
	"testService/internal/data"

	"gitlab.com/distributed_lab/logan/v3"
)

type ctxKey int

const (
	logCtxKey ctxKey = iota
	dbCxtKey
)

func CtxLog(entry *logan.Entry) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, logCtxKey, entry)
	}
}

func Log(r *http.Request) *logan.Entry {
	return r.Context().Value(logCtxKey).(*logan.Entry)
}

func CtxDB(entry data.Blobs) func(context.Context) context.Context {
	return func(ctx context.Context) context.Context {
		return context.WithValue(ctx, dbCxtKey, entry)
	}
}

func DB(r *http.Request) data.Blobs {
	return r.Context().Value(dbCxtKey).(data.Blobs)
}
