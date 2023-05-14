# go-slim-micro-service
## 以 grpc + etcd 为核心的微服务脚手架
#### · 基础服务（如db、redis等）通过 NewServiceContext 注入至 domain.ServiceContext 容器中。
#### · delivery.Register() 实例化注册服务绑定到对应的服务实现。
## 开源组件：grpc、etcd、gorm、yaml.v3
