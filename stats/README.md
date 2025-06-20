# stats

stats provides usecase for scoring your app metrics, commonly used for monitoring or capturing your business metric.

## How to Use

`stats.Stats` defines standard stats interface.

```go
type service struct {
  stats stats.Stats
}

func (s *service) DoSomething() {
  // something happened
  if err != nil {
    // Increment your failed metric count, used for monitoring or alerting
    s.stats.Increment("something_failed", stats.Option{
      Tags: stats.Tags{"method": "DoSomething"}
    })
  }
}
```

You can inject stats implementor using provided adapter.
```go

import (
  gogox_prom "github.com/hadan/gogox/stats/prometheus"
)

func main() {
  // e.g: prometheus stats, you can add base tags for your stats.
  prom := gogox_prom.New("mynamespace", stats.Tags{"service":"gogox_service"})

  service := service.NewService(prom)
}
```

Currently supported adapter:
1. Prometheus
2. Datadog (Will be in 1.x release)
3. Nop
