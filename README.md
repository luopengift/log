1. 日志输出到多个io.Writer中
2. 告警级别和控制台颜色
3. 日志文件按时间和大小轮转
4. 同步/异步flush
5. 可以使用Register(name string)来注册loghandler,调用SetLogger()来配置loghandler,使用loghandler来输出日志的话，会将loghandler的name写入日志内容中
如果没有的话，则使用默认的loghandler, 默认的是"__DEFAULT__"
默认的loghandler不输出
可以所有指定日志的格式
日志格式分成两部分: Head/content
Head同一类日志都是一样的，可以预先配置好,也可以自己重新实现
Content为日志的内容，格式比较多样化，可以临时指定


日志配置
各个output的配置

