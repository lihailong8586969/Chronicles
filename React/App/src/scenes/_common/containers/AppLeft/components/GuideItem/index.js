import React, {Component} from 'react';
import {findDOMNode} from 'react-dom';

require("./style.scss");

class GuideItem extends Component{

    constructor(props) {

        super(props);
        this.state = {date:new Date()};
    }

    generateContent() {

        let icon;
        if (this.props.role === "history") {
           
            icon = <div><i className="fas fa-history icon"></i><span className="text">历史记录</span></div>
        } else if (this.props.role === "logo") {

        } else if (this.props.role === "search") {

        } else{

        }

        return icon;
    }

    render(){

        return (

            <div className="guideItem">
                {this.generateContent()}
            </div>

        );
    }

}


export default GuideItem;


