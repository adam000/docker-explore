set -e

docker build -t explore-docker .
## to start detached, use -d/--detach
docker run -p 8080:8080 --restart on-failure explore-docker
