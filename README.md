# katydid_base_api
katydid base api

#### 构建单个镜像(Dockerfile) (katydid_base_api-client替换为自己的docker镜像名称)
#### dev:
```shell
    docker build -f deployments/docker/dev/client/Dockerfile -t  katydid_base_api-client_dev.
```
#### pro:
```shell
    docker build -f deployments/docker/prod/client/Dockerfile -t katydid_base_api-client_prod .
```
#### 构建组合镜像(docker-compose.yml)
#### dev:
```shell
    docker-compose -f deployments/docker/dev/docker-compose.yml build
```
#### prod:
```shell
    docker-compose -f deployments/docker/prod/docker-compose.yml build
```