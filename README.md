# Zoltraak: Distributed API Gateway & Traffic Control System

Zoltraak is a production-grade, distributed, and configuration-driven API Gateway in Go. 

Inspired by *Frieren: Beyond Journey’s End*, Zoltraak acts as the defensive shield and offensive routing spear for high-performance downstream microservices. It intercepts edge traffic to manage authentication, rate limiting, load balancing, circuit breaking, and retry logic before forwarding requests.

---

## 🏗️ Frieren-Themed Module Architecture

The gateway is composed of dedicated, decoupled modules:

*   **Zoltraak** (Main Gateway Engine): Boots the systems, accepts connections, and orchestrates the request pipeline.
*   **Grimoire** (Routing Engine): Matches paths, rewrites URLs, injects request IDs, and maps routes to clusters.
*   **Spellbook** (Configuration System): Parses and validates YAML rules, supporting thread-safe hot-reloads at runtime.
*   **ManaWall** (Rate Limiter): Distributed IP and API Key rate limiter backed by Redis Lua script transactions.
*   **Flamme** (Circuit Breaker): A finite state machine protecting backends from cascading latency and failure storms.
*   **Himmel** (Load Balancer): Selects backend targets using Round Robin, Weighted Round Robin, or Least Connections.
*   **Fern** (Health Checker): Periodically executes active and passive health probes to evict unhealthy backend instances.
*   **Stark** (Retry Engine): Retries transient failures on healthy hosts using exponential backoff, full jitter, and retry budgets.
*   **Aureole** (Service Discovery): Resolves backend target endpoints dynamically (via DNS/Consul/static definitions).
*   **Chronicle** (Observability & Metrics): Exposes latency histograms, error rates, and connection states to Prometheus.

---

## 🚀 Key Features

1.  **High-Performance Reverse Proxying**: Built on tuned `httputil.ReverseProxy` transport pools.
2.  **Thread-Safe Dynamic Configuration**: Config reloads without process restarts via atomic memory pointer swaps.
3.  **Distributed Rate Limiting**: Centralized atomic counters run inside a Redis Cluster using custom Lua scripts.
4.  **Circuit Breaking & Self-Healing**: Automated downstream isolation with cooldown periods and probe-based recovery.
5.  **Smart Retries**: Prevents thundering herd problems using randomized backoffs and token-based retry budgets.
6.  **Observability & Tracking**: Exposes Prometheus metrics and structures logs into performance-friendly JSON via Zap.

---

## 📁 Repository Layout

```
Zoltraak/
├── cmd/
│   └── gateway/
│       └── main.go                 # Gateway main entrypoint
├── internal/
│   ├── app/                        # Server engine orchestration
│   ├── config/                     # Spellbook configuration system
│   ├── engine/                     # Core gateway components (router, proxy, cb)
│   └── pkg/                        # Private internal utilities
├── pkg/
│   └── logger/                     # Shared structured logging (Zap wrapper)
├── configs/
│   └── default.yaml                # Example configuration schema
├── deployments/                    # Docker, Docker Compose, Prometheus and Grafana configs
└── scripts/                        # Load testing and benchmarking scripts
```

---

## 🛠️ Development Guides & Documents

For detailed implementation instructions, refer to the following project documents:

*   **System Architecture Blueprint**: [zoltraak_architecture_blueprint.md](file:///home/abhijit_1859/Documents/codes/Zoltraak/zoltraak_architecture_blueprint.md)
*   **Layer-by-Layer Learning Guide**: [learning_guide.md](file:///home/abhijit_1859/Documents/codes/Zoltraak/learning_guide.md)
*   **Simplified Layman's Manual**: [layman_guide.md](file:///home/abhijit_1859/Documents/codes/Zoltraak/layman_guide.md)
*   **Interactive Task List**: [todo_list.md](file:///home/abhijit_1859/Documents/codes/Zoltraak/todo_list.md)
*   **Phase 1 Code Guide**: [phase_1_guide.md](file:///home/abhijit_1859/Documents/codes/Zoltraak/phase_1_guide.md)
