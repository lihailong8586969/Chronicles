> * [PHP中被忽略的性能优化利器：生成器](https://segmentfault.com/a/1190000012334856)
> * [PHP性能优化利器：生成器 yield理解](https://www.cnblogs.com/zuochuang/p/8176868.html)

首先明确一个概念：生成器yield关键字不是返回值，他的专业术语叫产出值，只是生成一个值

那么代码中 foreach 循环的是什么？其实是PHP在使用生成器的时候，会返回一个 Generator 类的对象。 foreach 可以对该对象进行迭代，每一次迭代，PHP会通过 Generator 实例计算出下一次需要迭代的值。这样 foreach 就知道下一次需要迭代的值了。

而且，在运行中 for 循环执行后，会立即停止。等待 foreach 下次循环时候再次和  for  索要下次的值的时候，循环才会再执行一次，然后立即再次停止。直到不满足条件不执行结束。


未使用生成器时：
createRange 函数内的 for 循环结果被很快放到 $data 中，并且立即返回。所以， foreach 循环的是一个固定的数组。

使用生成器时： 
createRange 的值不是一次性快速生成，而是依赖于 foreach 循环。 foreach 循环一次， for 执行一次。