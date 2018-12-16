// Deleting array elements in JavaScript - delete vs splice
// https://stackoverflow.com/questions/500606/deleting-array-elements-in-javascript-delete-vs-splice/9815010#9815010


// Array Remove - By John Resig (MIT Licensed)
Array.prototype.remove = function(from, to) {
    var rest = this.slice((to || from) + 1 || this.length);
    this.length = from < 0 ? this.length + from : from;
    return this.push.apply(this, rest);
};

// Remove the second item from the array
array.remove(1);
// Remove the second-to-last item from the array
array.remove(-2);
// Remove the second and third items from the array
array.remove(1,2);
// Remove the last and second-to-last items from the array
array.remove(-2,-1);


// 方案错误：

// data 是数组
let data = YAML.safeLoad(fs.readFileSync( file ).toString());
data.words.forEach((word, index)=>{
    if(word.spell === spell){

        // 如果已存在则删掉
        console.log(index);

        // 部分元素未删除
        // 错误的原因是数组已经改变，下一个循环的时候 index 增加，导致在数组中无法找到
        data.words.remove(index);
    }
});


// 解决方案 ：

// 获取库内容
let data=[], words = YAML.safeLoad(fs.readFileSync( file ).toString()).words;
words.forEach((word, index)=>{
	if(word.spell !== spell){

	    data.push({ "spell":word.spell, "explain":word.explain });
	}
});
data.push({ "spell":spell, "explain":explain });