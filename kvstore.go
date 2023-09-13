package kvstore

import (
	"context"
	"errors"
)

// Store represents the key-value storage.
type Store interface {
	Put(ctx context.Context, key string, value []byte) error
	Get(ctx context.Context, key string) (*Pair, error)
	Delete(ctx context.Context, key string) error
	Exists(ctx context.Context, key string) (bool, error)
}

// AtomicStore represents the key-value storage that supports atomic operations.
type AtomicStore interface {
	AtomicPut(ctx context.Context, key string, value []byte, prev *Pair) (bool, *Pair, error)
	AtomicGet(ctx context.Context, key string) (*Pair, error)
	AtomicDelete(ctx context.Context, key string, prev *Pair) (bool, error)
}

// WatchStore represents the key-value storage that supports watch operations.
type WatchStore interface {
	Watch(ctx context.Context, key string) (<-chan *Pair, error)
	WatchTree(ctx context.Context, prefix string) (<-chan []*Pair, error)
}

// Pair stored in key-value storage.
type Pair struct {
	Key   string
	Value []byte
}

// Key-value storage errors.
var (
	ErrKeyNotFound = errors.New("key not found")
)
