package server

import (
	"context"
	"fmt"
	"github.com/flyteorg/flyte/flyteartifacts/pkg/blob"
	"github.com/flyteorg/flyte/flyteartifacts/pkg/configuration"
	"github.com/flyteorg/flyte/flyteartifacts/pkg/db"
	"github.com/flyteorg/flyte/flyteartifacts/pkg/server/processor"
	"github.com/flyteorg/flyte/flyteidl/gen/pb-go/flyteidl/artifact"
	"github.com/flyteorg/flyte/flytestdlib/database"
	"github.com/flyteorg/flyte/flytestdlib/promutils"
	"github.com/go-gormigrate/gormigrate/v2"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/pkg/errors"
	"google.golang.org/grpc"

	_ "net/http/pprof" // Required to serve application.
)

type ArtifactService struct {
	artifact.UnimplementedArtifactRegistryServer
	Metrics       ServiceMetrics
	Service       CoreService
	EventConsumer processor.EventsProcessorInterface
}

func NewArtifactService(ctx context.Context, scope promutils.Scope) *ArtifactService {
	cfg := configuration.GetApplicationConfig()
	fmt.Println(cfg)
	eventsCfg := configuration.GetEventsProcessorConfig()

	storage := db.NewStorage(ctx, scope.NewSubScope("storage:rds"))
	blobStore := blob.NewArtifactBlobStore(ctx, scope.NewSubScope("storage:s3"))
	coreService := NewCoreService(storage, &blobStore, scope.NewSubScope("server"))
	eventsReceiverAndHandler := processor.NewBackgroundProcessor(ctx, *eventsCfg, &coreService, scope.NewSubScope("events"))
	if eventsReceiverAndHandler != nil {
		eventsReceiverAndHandler.StartProcessing(ctx)
	}

	return &ArtifactService{
		Metrics:       InitMetrics(scope),
		Service:       coreService,
		EventConsumer: eventsReceiverAndHandler,
	}
}

func HttpRegistrationHook(ctx context.Context, gwmux *runtime.ServeMux, grpcAddress string, grpcConnectionOpts []grpc.DialOption, _ promutils.Scope) error {
	err := artifact.RegisterArtifactRegistryHandlerFromEndpoint(ctx, gwmux, grpcAddress, grpcConnectionOpts)
	if err != nil {
		return errors.Wrap(err, "error registering execution service")
	}
	return nil
}

func GrpcRegistrationHook(ctx context.Context, server *grpc.Server, scope promutils.Scope) error {
	serviceImpl := NewArtifactService(ctx, scope)
	artifact.RegisterArtifactRegistryServer(server, serviceImpl)

	return nil
}

// GetMigrations should be hidden behind the storage interface in the future.
func GetMigrations(_ context.Context) []*gormigrate.Migration {
	return db.Migrations
}
func GetDbConfig() *database.DbConfig {
	cfg := configuration.ApplicationConfig.GetConfig().(*configuration.ApplicationConfiguration)
	return &cfg.ArtifactDatabaseConfig
}
