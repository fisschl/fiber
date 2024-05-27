$image = "registry.cn-shanghai.aliyuncs.com/fisschl/fiber:latest"

docker build -t $image .
docker push $image
