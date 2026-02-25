## URL Shortener Architecture

1. **Collision**

- if use random:
  - db unique
  - retry loop

- if use base62 (recommended):
  - never collision
  - more clean

2. **Hot key problem**

- If 1 url varial:
  - redis cache/read too much
  - db not reached cause cache hitted
- Advance solution:
  - local in memory cache (tiny LRU)
  - CDN layer

3. **Scalability**

- Partition table base on hash (short code)
- Or shard by id range
- Read replica
- Redis cluster

4. **Architecture Diagram**

```
Client
    |
    v
Load Balancer
    |
    v
Web Servers (API)
    |
    v
Redis Cache
    |
    v
Database (MySQL/PostgreSQL)
```

5. **Additional Features**

```
Rate limiting (IP-based)

Custom alias support

Expiration support

Async analytics via queue

Prometheus metrics

Clean architecture

Docker compose (postgres + redis)

```
