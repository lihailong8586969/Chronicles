import React, {Component} from 'react';
import {findDOMNode} from 'react-dom';

import { Row, Col } from 'antd';

class Home extends Component{

    constructor(props) {

        super(props);
        this.state = {date:new Date()};
    }

    render(){

        return (

             <div id="Route-to-Home">
                <p>Home</p>
            </div>

        );
    }

}


export default Home;
