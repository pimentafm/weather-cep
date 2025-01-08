#!/bin/bash

echo "Starting Docker cleanup..."

# Stop all running containers
echo "Stopping containers..."
docker stop $(docker ps -aq) 2>/dev/null

# Remove all containers
echo "Removing containers..."
docker rm $(docker ps -aq) 2>/dev/null

# Remove all networks (except default ones)
echo "Removing networks..."
docker network rm $(docker network ls -q) 2>/dev/null

# Remove all images
echo "Removing images..."
docker rmi $(docker images -q) -f 2>/dev/null

# Remove all volumes
echo "Removing volumes..."
docker volume rm $(docker volume ls -q) 2>/dev/null

# Final cleanup
echo "Performing final cleanup..."
docker system prune -af --volumes 2>/dev/null

echo "Cleanup complete!"