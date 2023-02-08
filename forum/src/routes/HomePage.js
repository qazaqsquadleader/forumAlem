import React from 'react';
import Header from '../components/Header/Header';
import Body from '../components/Body/Body';

import { useLocation } from "react-router-dom";


const HomePage = () => {
// const {username} = props;
const location = useLocation();
const { state } = location;
const username = state.username;


    return (
        <div>
        <Header />
        <Body />
        <div>
        <h2> Hello {username}</h2>

        </div>
        </div>
    );
};

export default HomePage;