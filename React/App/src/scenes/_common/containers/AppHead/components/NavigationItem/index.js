import React, {Component} from 'react';
import {findDOMNode} from 'react-dom';


// 输入框
import { Input } from 'antd';


// fontawesome
import {library} from '@fortawesome/fontawesome-svg-core'
import {FontAwesomeIcon} from '@fortawesome/react-fontawesome'
import {faBars} from '@fortawesome/free-solid-svg-icons'
library.add(faBars);


// 需要放到 import 的后面
require("./style.scss");

class NavigationItem extends Component {

    state = {
    
        menuButtonStatus: "",
    };


    constructor(props) {

        super(props);
    }

    handleMenuClick(e){

        if(this.state.menuButtonStatus===""){

            this.setState({menuButtonStatus:"active"});
        }else{

            this.setState({menuButtonStatus:""});
        }
    }

    generateContent() {

        let icon;
        if (this.props.role === "menu") {
           
            icon = <button onClick={(e)=>this.handleMenuClick(e)} className={ this.state.menuButtonStatus + " btn-menu"} ><FontAwesomeIcon className="icon" icon="bars" size="lg"/></button>;
        } else if (this.props.role === "logo") {

            // icon = <div className="logo"><span className="text">Youtube</span><i className="fab fa-youtube"></i></div>;
            icon = <div className="logo"><img src={require("./images/youtube.png")} /></div>;
        } else if (this.props.role === "search") {

            const Search = Input.Search;
            icon = <Search  placeholder="input search text" enterButton size="large" onSearch={value => console.log(value)} />;
        } else{


        }

        return icon;
    }

    render() {

        return (

            <div className="navigationItem">
                {this.generateContent()}
            </div>

        );
    }

}


export default NavigationItem;


