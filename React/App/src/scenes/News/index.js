import React, {Component} from 'react';
import {findDOMNode} from 'react-dom';
import AppLeft from '../_common/containers/AppLeft/index'
import AppBody from '../_common/containers/APPBody/index'
import AppHead from "../_common/containers/AppHead/index";

import { Row, Col } from 'antd';

class App extends Component{

    constructor(props) {

        super(props);
        this.state = {date:new Date()};
    }

    render(){

        return (

            <div id="Route-to-News">
                <p>News</p>
            </div>

        );
    }

}


export default App;
