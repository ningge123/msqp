package discovery

// Register grpc 注册到etcd
// 原理：创建一个租约，grpc服务注册到etcd，绑定租约
// 过了租约时间，etcd就会删除grpc服务信息
// 实现心跳，完成续租，如果etcd没有就新注册

type Register struct {
}
