### 词法分析

> 将源代码转为token序列

#### lex

> lex是生成词法分析器的工具，用其生成的工具可以将源代码转换为token序列

> lex文件不能直接使用，需要使用lex命令将其转为C语言代码，然后编译为可执行文件

![8_lex_compile.png](../../img/8_lex_compile.png)


#### go

> go的词法分析是通过`src/cmd/compile/internal/syntax/scanner.go`
