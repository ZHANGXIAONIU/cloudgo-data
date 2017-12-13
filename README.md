# Cloudgo-IO

使用Go语言构建数据库服务。

## 实现

实现有两种方式，分别是使用原生`database/sql`和ORM技术，本仓库旨在通过比较两者性能上的差异，分析两者各自的优缺点。

### 原生`database/sql`实现

这里模仿Java的Jdbc编程风格，采用entity-dao-service的结构模型来编写数据服务。

- `entity`提供实体定义及有关操作，例如新建实体
- `dao`封装直接对数据库进行的操作
- `service`调用`dao`中方法，向上提供各种服务

具体代码可在`entities`目录下查看。

### `xorm`实现

ORM技术也是Java世界中广泛存在的技术；使用ORM库使得我们完全不需要编写任何数据库访问代码，直接执行SQL语句。具体代码可见服务端的路由函数。

举例：查询所有userinfo只需2行

    ulist := make([]entities.UserInfo, 0)
    mySQLEngine.Table("userinfo").Find(&ulist)

## 比较

原生`database/sql`优点是简单高效，性能好可自定义性强，而缺点显而易见，开发成本过高（需要类似于negroni这样的中间件来解决问题）；而ORM的优点就是即插即用，而缺点则是性能的不足。

### orm 是否就是实现了 dao 的自动化？

并不是，orm还实现了service的自动化。

### ab测试比较性能

分别对两种实现方式使用ab测试，测试它们的性能。

- 原生`database/sql`

    $ ab -n 10000 -c 1000 http://localhost:8080/service/userinfo?userid=
    This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
    Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
    Licensed to The Apache Software Foundation, http://www.apache.org/

    Benchmarking localhost (be patient)
    Completed 1000 requests
    Completed 2000 requests
    Completed 3000 requests
    Completed 4000 requests
    Completed 5000 requests
    Completed 6000 requests
    Completed 7000 requests
    Completed 8000 requests
    Completed 9000 requests
    Completed 10000 requests
    Finished 10000 requests


    Server Software:
    Server Hostname:        localhost
    Server Port:            8080

    Document Path:          /service/userinfo?userid=
    Document Length:        219 bytes

    Concurrency Level:      1000
    Time taken for tests:   3.214 seconds
    Complete requests:      10000
    Failed requests:        1661
       (Connect: 0, Receive: 0, Length: 1661, Exceptions: 0)
    Non-2xx responses:      1661
    Total transferred:      10062001 bytes
    HTML transferred:       8835289 bytes
    Requests per second:    3111.84 [#/sec] (mean)
    Time per request:       321.353 [ms] (mean)
    Time per request:       0.321 [ms] (mean, across all concurrent requests)
    Transfer rate:          3057.75 [Kbytes/sec] received

    Connection Times (ms)
                  min  mean[+/-sd] median   max
    Connect:        0   91 290.0      1    1045
    Processing:     0  127 250.2     64    2050
    Waiting:        0  126 250.3     63    2050
    Total:          0  218 466.4     66    2745

    Percentage of the requests served within a certain time (ms)
      50%     66
      66%     98
      75%    120
      80%    131
      90%   1050
      95%   1352
      98%   1970
      99%   2736
     100%   2745 (longest request)

- `xorm`

    $ ab -n 10000 -c 1000 http://localhost:8080/service/userinfo?userid=
    This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
    Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
    Licensed to The Apache Software Foundation, http://www.apache.org/

    Benchmarking localhost (be patient)
    Completed 1000 requests
    Completed 2000 requests
    Completed 3000 requests
    Completed 4000 requests
    Completed 5000 requests
    Completed 6000 requests
    Completed 7000 requests
    Completed 8000 requests
    Completed 9000 requests
    Completed 10000 requests
    Finished 10000 requests


    Server Software:
    Server Hostname:        localhost
    Server Port:            8080

    Document Path:          /service/userinfo?userid=
    Document Length:        229 bytes

    Concurrency Level:      1000
    Time taken for tests:   4.404 seconds
    Complete requests:      10000
    Failed requests:        4581
       (Connect: 0, Receive: 0, Length: 4581, Exceptions: 0)
    Total transferred:      2485532 bytes
    HTML transferred:       1254694 bytes
    Requests per second:    2270.64 [#/sec] (mean)
    Time per request:       440.404 [ms] (mean)
    Time per request:       0.440 [ms] (mean, across all concurrent requests)
    Transfer rate:          551.15 [Kbytes/sec] received

    Connection Times (ms)
                  min  mean[+/-sd] median   max
    Connect:        0  114 319.7      1    1032
    Processing:     2  149 366.7     68    3368
    Waiting:        2  147 366.8     67    3368
    Total:          2  263 611.4     70    4393

    Percentage of the requests served within a certain time (ms)
      50%     70
      66%     93
      75%    111
      80%    122
      90%   1089
      95%   1531
      98%   2749
      99%   2768
     100%   4393 (longest request)

可见，xorm要慢1秒多，性能差异还是较大。
