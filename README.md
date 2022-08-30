## go-zero 学习开发微服务项目

#### 各服务暴露的端口：
| 服务          | API 端口  | RPC 端口 |
|-------------|---------|--------|
| RBAC        | 18001   | 28001  |
| Account     | 18002   | 28002  |
| Application | 18003   | 28003  |
| Marketing   | 18004   | 28004  |

### 项目手动部署配置
#### 生成 Dockerfile 文件
进入 app 对应服务的 api|rpc 目录执行：`goctl docker -go svc.go[服务的go文件]`

#### 生成 K8s deployment 文件
```shell
# 生成模板，具体参数视服务而定
# api
goctl kube --nodePort 30020 deploy -replicas 1 -requestCpu 100 -requestMem 50 -limitCpu 200 -limitMem 100 -name account-api -namespace business -image silentcxl/account-api:v1.0.0 -o account-api-v1.0.0.yaml -port 18002 --serviceAccount find-endpoints

# rpc
goctl kube deploy -replicas 3 -requestCpu 100 -requestMem 50 -limitCpu 200 -limitMem 100 -name rbac-rpc -namespace business -image rbac-rpc:v1.0.0 -o rbac-rpc.yaml -port 28001 --serviceAccount find-endpoints
```

#### API 调用 RPC 配置文件服务发现参数
```yaml
RbacRpcClient:
  Target: "k8s://business/rbac-client-svc:28001"
```