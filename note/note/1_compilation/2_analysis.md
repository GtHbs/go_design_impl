### 词法分析

> 将源代码转为token序列

#### lex

> lex是生成词法分析器的工具，用其生成的工具可以将源代码转换为token序列

> lex文件不能直接使用，需要使用lex命令将其转为C语言代码，然后编译为可执行文件

![8_lex_compile.png](../../img/8_lex_compile.png)


#### go

> go的词法分析是通过`src/cmd/compile/internal/syntax/scanner.go`中的scanner结构体实现的

### 语法分析

> 根据某种特定的形式文法对token序列构成的输入文本进行分析并确定起语法结果的过程。
> 
> 语法分析的过程会使用自顶向下或自底向上的方式进行推导。


#### 文法














