// 很有创意：http://blog.youlunshidai.com/resume/js/index.js
// 原项目 ： https://github.com/ccfromstar/resume


var Coder = function(dom) {

	this._dom = dom;
};

Coder.prototype = {

	constructor: Coder,
	
	// 写文字 ...
	write: function(code, second) {
		var dom = this._dom;
		var o = this;
		var _second = second ? second : 1000;
		_second = m + _second;
		setTimeout(function() {
			$('#' + dom).scrollTop( $('#' + dom)[0].scrollHeight );
			$("#" + dom).append(code);
		}, _second);
		m = _second;
	},
	
	setClass: function(dom, styles,second) {
		var _second = second ? second : 0;
		_second = m + _second;
		for (var i in styles) {
			setTimeout(function() {
				$(dom).css(styles[i].name, styles[i].val);
			}, _second);
		}
		m = _second;
	},
	
	addClass:function(dom,className){
		var second = m;
		setTimeout(function() {
			$(dom).addClass(className);
		}, second);
	}
}