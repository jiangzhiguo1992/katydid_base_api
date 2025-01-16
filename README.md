# katydid_base_api
katydid base api

#### 构建单个dockerfile (katydid_base_api-client替换为自己的docker镜像名称)
#### dev:
```shell
    docker build -f deployments/docker/client/dev/Dockerfile -t  katydid_base_api-client.
```
#### pro:
```shell
    docker build -f deployments/docker/client/pro/Dockerfile -t katydid_base_api-client .
```
