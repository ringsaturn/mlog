
```go
package main

import (
	"context"

	"github.com/ringsaturn/mlog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func DefaultComponentLevels() map[options.LogComponent]options.LogLevel {
	return map[options.LogComponent]options.LogLevel{
		options.LogComponentAll:             options.LogLevelDebug,
		options.LogComponentCommand:         options.LogLevelDebug,
		options.LogComponentTopology:        options.LogLevelDebug,
		options.LogComponentServerSelection: options.LogLevelDebug,
		options.LogComponentConnection:      options.LogLevelDebug,
	}
}

func main() {
	ctx := context.Background()

	var l, _ = zap.NewDevelopment()

	loggerOpts := options.Logger()

	loggerOpts.Sink = mlog.New(l)
	loggerOpts.ComponentLevels = DefaultComponentLevels()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017").SetLoggerOptions(loggerOpts))
	if err != nil {
		panic(err)
	}
	defer client.Disconnect(ctx)

	err = client.Ping(ctx, nil)

	if err != nil {
		panic(err)
	}

	_ = client.Database("wp").Collection("pm25").FindOne(ctx, nil)

}
```
