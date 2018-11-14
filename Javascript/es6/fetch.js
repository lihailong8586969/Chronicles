// 由于 Fetch API 是基于 Promise 设计, 有必要先学习一下 Promise

let workingAPI = "https://api.wordnik.com/ds/words.json";

function setup(){

	let promise = fetch(workingAPI);
	promise.then(response=>{

		return response.json();
	}).then(json=>{

		return fetch();
	}).then(response=>{

		return response.json();
	}).then(json=>{

		createImg(json.data[0].images['fixed']);
	}).catch(err=>{

		console.log(err);
	});

}
