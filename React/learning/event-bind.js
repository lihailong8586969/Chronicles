/*--------------------第1种-------------------------*/
class Counter extends React.Component{
  constructor(props) {
      super(props);
      this.state = {clickCount: 0};
  }
  
  handleClick(e) {
    this.setState(function(state) {
      return {clickCount: state.clickCount + 1};
    });
  }
  render () {
    return (<h2 onClick={(e)=>this.handleClick(e)}>点我！点击次数为: {this.state.clickCount}</h2>);
  }
}
ReactDOM.render(
  <Counter />,
  document.getElementById('example')
);



/*--------------------第2种-------------------------*/
class Counter extends React.Component{
  constructor(props) {
      super(props);
      this.state = {clickCount: 0};
  }
  
  handleClick() {
    this.setState(function(state) {
      return {clickCount: state.clickCount + 1};
    });
  }
  render () {
    return (<h2 onClick={this.handleClick.bind(this)}>点我！点击次数为: {this.state.clickCount}</h2>);
  }
}
ReactDOM.render(
  <Counter />,
  document.getElementById('example')
);


/*--------------------第3种-------------------------*/
class Counter extends React.Component{
  constructor(props) {
      super(props);
      this.state = {clickCount: 0};

      this.handleClick = this.handleClick.bind(this); // 重要 ！！！
  }
  
  handleClick() {
    this.setState(function(state) {
      return {clickCount: state.clickCount + 1};
    });
  }
  render () {
    return (<h2 onClick={this.handleClick}>点我！点击次数为: {this.state.clickCount}</h2>);
  }
}
ReactDOM.render(
  <Counter />,
  document.getElementById('example')
);
