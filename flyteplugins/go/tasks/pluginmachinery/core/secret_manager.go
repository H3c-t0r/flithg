package core

import (
	"context"
	coreIdl "github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/core"
)

type SecretManager interface {
	Get(ctx context.Context, key string) (string, error)
	GetForSecret(ctx context.Context, secret *coreIdl.Secret) (string, error)
}
