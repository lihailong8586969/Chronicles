import React from 'react';
import ReactDOM from 'react-dom';
import * as serviceWorker from './services/serviceWorker';
import { BrowserRouter, Link } from 'react-router-dom';
import { Router, Route } from 'react-router';

import AppLeft from './scenes/_common/containers/AppLeft/index'
import AppBody from './scenes/_common/containers/APPBody/index'
import AppHead from "./scenes/_common/containers/AppHead/index";

import Home from "./scenes/Home/index";
import News from "./scenes/News/index";

import { Row, Col } from 'antd';

require("./scenes/_common/css/global.scss");

// package.json 默认使用的是 node_modules 下的 react-scripts/config/webpack.config.dev.js



const App = (

    <BrowserRouter >
      <div id="App">
            <AppHead />
            <div id="AppMain">
                <Row>
                    <Col span={4}><AppLeft /></Col>
                    <Col span={20}>
                        <AppBody>
                            <Route path="/" component={Home}>
                                <Route path="home" component={Home} />
                             	<Route path="news" component={News} />
                            </Route>
                        </AppBody>
                    </Col>
                </Row>
            </div>
        </div>
    </BrowserRouter>

);



// 项目的入口文件

ReactDOM.render(App, document.getElementById('Root'));

// If you want your app to work offline and load faster, you can change
// unregister() to register() below. Note this comes with some pitfalls.
// Learn more about service workers: http://bit.ly/CRA-PWA
serviceWorker.unregister();

