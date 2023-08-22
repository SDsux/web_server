# Go实现简易Web框架
主要是根据极客兔兔博客学习得到，加入了自己的理解而成
源网站：https://geektutu.com/post/gee.html
## 相关实现
1、利用前缀树实现动态路。
2、设计上下文(Context)，封装 Request 和 Response ，提供对 JSON、HTML 等返回类型的支持
3、实现分组控制
4、预留中间件插入模块
5、提供错误恢复函数
