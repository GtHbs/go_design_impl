### error

> `panic(v interface{})`
> 
> `recover() interface{}`

> golang中error是接口类型，在函数执行过程中可以调用panic抛出异常，
> 抛出异常之前会指定defer语句，如果在defer中使用了recover函数，则函数正常执行，
> 否则对于上层函数而已就等于是在外层调用了panic