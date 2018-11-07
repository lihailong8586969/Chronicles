import React, {Component} from 'react';
import {findDOMNode} from 'react-dom';
import GuideItem from './components/GuideItem/index'

import { Row, Col } from 'antd';


require('./style.scss');


class AppLeft extends Component {

    constructor(props) {

        super(props);
    }

    render() {

        return (

            <div id="AppLeft">
                <GuideItem role="history" />
                <GuideItem role="history" />
                <GuideItem role="history" />
                <div className="hr"></div>
                <GuideItem role="history" />
                <GuideItem role="history" />
                <GuideItem role="history" />
            </div>

        );
    }

}


export default AppLeft;


