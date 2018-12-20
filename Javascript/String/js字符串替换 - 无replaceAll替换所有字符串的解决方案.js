// js字符串替换 - 无replaceAll替换所有字符串的解决方案

// https://blog.csdn.net/chwshuang/article/details/52583496


// replace() 只能替换一次，所以，用 split + join 即可解决

localName = localName.split('.').join('');


