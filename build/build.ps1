$image = "registry.cn-shanghai.aliyuncs.com/fisschl/golang:latest"

docker pull golang:1.22
docker tag golang:1.22 $image
docker push $image
 