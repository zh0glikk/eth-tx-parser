package interfaces

import (
	"context"
	"sync"
)

type Runner interface {
	Run(ctx context.Context, wg *sync.WaitGroup)
}
