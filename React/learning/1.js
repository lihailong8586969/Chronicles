/*------------------ 可以直接把函数当成组件调用 ------------------*/
function NumberList(props) {

  const numbers = props.numbers;
  const listItems = numbers.map((number) =>
    <li key={number.toString()}>
      {number}
    </li>
  );

  return (
    <ul>{listItems}</ul>
  );
}

const numbers = [1, 2, 3, 4, 5];

ReactDOM.render(
  
  // 可以直接把函数当成组件调用
  <NumberList numbers={numbers} />,
  document.getElementById('example')
);/*------------------ 可以直接把函数当成组件调用 ------------------*/


/*------------------ key 的正确使用 ------------------*/
function ListItem(props) {
  // 对啦！这里不需要指定key:
  return <li>{props.value}</li>;
}

function NumberList(props) {
  const numbers = props.numbers;
  const listItems = numbers.map((number) =>
    // 又对啦！key应该在数组的上下文中被指定
    <ListItem key={number.toString()}
              value={number} />
  );
  return (
    <ul>
      {listItems}
    </ul>
  );
}

const numbers = [1, 2, 3, 4, 5];
ReactDOM.render(
  <NumberList numbers={numbers} />,
  document.getElementById('example')
);/*------------------ key 的正确使用 ------------------*/


