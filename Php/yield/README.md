这里的间隔一秒其实就是 sleep(1) 造成的后果。但是为什么第一次没有间隔？那是因为：

未使用生成器时： createRange 函数内的 for 循环结果被很快放到 $data 中，并且立即返回。所以， foreach 循环的是一个固定的数组。
使用生成器时： createRange 的值不是一次性快速生成，而是依赖于 foreach 循环。 foreach 循环一次， for 执行一次。