# Docker容器的数据卷

1. 数据卷概念及作用  
数据卷概念:数据卷是宿主机中的一个目录或文件,当容器目录和数据卷目录绑定后,对方的修改会立即绑定。一个数据卷可以被多个容器同时挂载,一个容器也可以被挂载多个数据卷  
数据卷作用:容器数据持久化  外部机器和容器间通信 容器之间数据交换



2. 配置数据卷  
docker run -it name=c1 -v 宿主机目录(文件):容器内目录(文件)  
- 注意事项  
1. 目录必须是绝对路径
2. 如果目录不存在,会自动创建
3. 可以挂载多个数据卷



3. 配置数据卷容器