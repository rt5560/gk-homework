### 1.使用 redis benchmark 工具, 测试 10 20 50 100 200 1k 5k 字节 value 大小，redis get set 性能。

* docker启动redis
>docker run -p 6379:6379 --name redis -d redis

* 使用redis-benchmark进行测试
>redis-benchmark -t set,get -n 100000 -d 10

测试结果如下
![测试结果](http://7xthla.com1.z0.glb.clouddn.com/redis-benchmark_result.jpg)

---
###2.写入一定量的 kv 数据, 根据数据大小 1w-50w 自己评估, 结合写入前后的 info memory 信息 , 分析上述不同 value 大小下，平均每个 key 的占用内存空间。

![测试结果](http://7xthla.com1.z0.glb.clouddn.com/1642951285631.jpg)