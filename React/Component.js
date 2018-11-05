import React, {Component} from 'react';
import {findDOMNode} from 'react-dom';
import logo from './logo.svg';
import './App.css';

class App extends Component {

    constructor(props) {

        super(props);
        this.state = {date:new Date()};
    }

    // 组件的生命周期可分成三个状态：
    // Mounting：已插入真实 DOM
    // Updating：正在被重新渲染
    // Unmounting：已移出真实 DOM


    // componentWillMount 在渲染前调用,在客户端也在服务端。
    componentWillMount(){

        console.log("componentWillMount");

        // 使用同步函数alert()阻塞一下, 即可发现此时组件还未挂载
        // alert("componentWillMount : 组件还没挂载呢！不过下0.000001秒就挂上了");
    }

    // 在第一次渲染后调用，只在客户端。
    // 之后组件已经生成了对应的DOM结构，可以通过this.getDOMNode()来进行访问。
    // 如果你想和其他JavaScript框架一起使用，
    // 可以在这个方法中调用setTimeout, setInterval或者发送AJAX请求等操作(防止异步操作阻塞UI)。
    componentDidMount() {

        console.log("componentDidMount ： this.getDOMNode() = ");
        console.log(findDOMNode(this));
        // alert("componentDidMount : 挂上了！");

        this.timerID = setInterval(
            () => this.tick(),
            1000
        );
    }

    // componentWillReceiveProps 在组件接收到一个新的 prop (更新后)时被调用。这个方法在初始化render时不会被调用。
    componentWillReceiveProps(){

        console.log("componentWillReceiveProps");
    }

    // shouldComponentUpdate 返回一个布尔值。
    // 在组件接收到新的props或者state时被调用。
    // 在初始化时或者使用forceUpdate时不被调用。
    // shouldComponentUpdate(){
    //
    //     console.log("shouldComponentUpdate");
    //     return true;
    // }

    // componentWillUpdate 在组件接收到新的props或者state但还没有render时被调用。在初始化时不会被调用。
    // componentWillUpdate(){
    //
    //     console.log("componentWillUpdate");
    // }

    // componentDidUpdate 在组件完成更新后立即调用。在初始化时不会被调用。
    // componentDidUpdate(){
    //
    //     console.log("componentDidUpdate");
    // }

    // 在组件从 DOM 中移除之前立刻被调用。
    componentWillUnmount() {

        console.log("componentWillUnmount");
        clearInterval(this.timerID);
    }

    render() {
        return (
            <div className="App">
                <header className="App-header">
                    <img src={logo} className="App-logo" alt="logo"/>
                    <p>
                        Edit <code>src/App.js</code> and save to reload.
                    </p>

                    <p>现在是 {this.state.date.toLocaleTimeString()} </p>

                    <a
                        className="App-link"
                        href="https://reactjs.org"
                        target="_blank"
                        rel="noopener noreferrer"
                    >
                        Learn React
                    </a>
                </header>
            </div>
        );
    }


    tick(){

        this.setState({

            date: new Date()
        });
    }

}

export default App;
