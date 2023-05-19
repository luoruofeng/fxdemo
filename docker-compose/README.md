# consul
**集群信息如下**    
* 创建了一个网络。  
* 将数据和日志信息存储到当前文件夹下的卷积。  
* 启动了3个server节点，其中1个是leader，另外两个是follower。启动了ui功能可以通过*localhost:8500*访问。   
* 启动了一个client节点，通过exec操作集群。

----

* **查询consul集群节点信息**
```shell
docker exec consul-client consul members
```