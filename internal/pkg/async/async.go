package async

import (
	"context"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
)

// Future interface has the method signature for await
type Future interface {
	Await() (*types.Transaction, error)
}

type future struct {
	await func(ctx context.Context) (*types.Transaction, error)
}

func (f future) Await() (*types.Transaction, error) {
	return f.await(context.Background())
}

// Exec executes the async function
func Exec(f func() (*types.Transaction, error)) Future {
	var result *types.Transaction
	var err error
	c := make(chan struct{})
	go func() {
		defer close(c)
		result, err = f()
	}()
	return future{
		await: func(ctx context.Context) (*types.Transaction, error) {
			select {
			case <-ctx.Done():
				log.Println("aaaaaaaaaaaaaaaaaaaaa")
				return nil, ctx.Err()
			case <-c:
				return result, err
			}
		},
	}
}
