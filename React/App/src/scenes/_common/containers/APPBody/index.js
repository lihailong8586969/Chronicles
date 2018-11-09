import React, {Component} from 'react';
import {findDOMNode} from 'react-dom';



class AppBody extends Component{

    constructor(props) {

        super(props);
        this.state = {date:new Date()};
    }

    render(){

        return (

            <div id="AppBody">
                <p>AppBody</p>
            </div>

        );
    }

}


export default AppBody;


