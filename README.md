# Learning Docker From First Principles

This repo documents my hands-on journey learning Docker and containerized systems from the ground up, with a focus on building a correct mental model of how containers, images, networking, storage, and orchestration actually work.

---

## What This Repository Covers

### Dockerfiles
- Image layering & multi-stage builds
- Minimal production images (build vs runtime)

### Images & Containers
- Images = blueprints, containers = running instances
- Isolation: filesystem, processes, user space
- Rebuilding images vs restarting containers

### Container Orchestration
- Docker Compose: multi-container setup, dependencies, health checks, restart policies
- Horizontal scaling of services

### Networking
- Container-to-container communication
- Docker DNS & service name resolution
- Host vs container port bindings

### Storage
- Ephemeral filesystems vs persistent volumes
- Running PostgreSQL in containers safely

### Configuration
- Environment variables & runtime config
- Avoid baking secrets in images

### Reverse Proxy & Load Balancing
- NGINX limitations with dynamic containers
- Traefik: automatic discovery, health-aware routing, zero-reload load balancing


---

                    ##        .            
              ## ## ##       ==            
           ## ## ## ##      ===            
       /""""""""""""""""\___/ ===        
      {~~ ~~~~ ~~~ ~~~~ ~~ ~ /  ===- ~~~   
       \______ o          __/            
         \    \        __/             
          \____\______/        

          |          |
       __ |  __   __ | _  __   _
      /  \| /  \ /   |/  / _\ | 
      \__/| \__/ \__ |\_ \__  |        