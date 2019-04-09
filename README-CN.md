# 使用

如果您想要使用 `Debug` 功能，您需要在项目中增加如下代码：
```go
   var debug utils.Debug           //  声明一个 Debug 类型的变量
   debug = utils.Init("flag")      //  初始化变量，你可以使用任意字符串替换flag
   debug("Print what you want!")   //  打印你需要的信息
```
除此之外， 您还需要设置 `Debug` 环境变量，该变量的值与你初始化 debug 变量时传入的字符串一致。