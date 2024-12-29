DOCKER_DEFAULT_PLATFORM="linux/amd64" docker build -t gossip .
aws lightsail push-container-image --region ap-southeast-1 --service-name gossip --label backend --image gossip:latest --profile personal