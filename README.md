# Using

If you want to use the `Debug` function, you need to add the following code to your project:
```go
   var debug utils.Debug           //  Declare a variable of type Debug
   debug = utils.Init("flag")      //  Initialize variables, you can replace flag with any string
   debug("Print what you want!")   //  Print the information you need
```
In addition to this, you also need to set the `Debug` environment variable, which has the same value as the string you passed in when you initialized the debug variable.
