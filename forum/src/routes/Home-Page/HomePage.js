import React, {useEffect, useState} from 'react';
import Header from '../../components/Header/Header';
import Body from '../../components/Body/Body';
import './HomePage.css'
// import { json } from 'react-router-dom';
const HomePage = (props) => {
// const [isAuth, setIsAuth] = useState(false);
// const [user, setUser] = useState();
const [posts, setPosts] = useState([{}]);
console.log(props);
useEffect (() => {
  // setIsAuth(props.isAuth)
  // setUser(props.username)  
  const fetchData = async () => {
      
      // await fetch(`http://localhost:8080/api/checkUser`, {
      //   headers: {
      //     'Accept': 'application/json ',
      //     'Credentials': 'include'
      //   },
      //   method: "GET",
      //   credentials: 'include',
      // }).then((r) => {
      //     if(r.ok){
      //       setIsAuth (true);
      //       return r.json();
      //     } else if (r.status === 401){
      //       setIsAuth(false)
      //       return null
      //     } else {
      //       throw new Error("Server error")
      //     }
      //   }
      //   )
      
      await fetch(`http://localhost:8080/api/home`, {
          headers: {
            'Accept': 'application/json ',
            'Credentials': 'include'
          },
          method: "GET",
          credentials: 'include',
        }).then((r) => r.json())
        .then((data) => {
            if (data !== null){
              setPosts(data)
            } else {
              console.log("No posts in Homepage");
            }
        });
      };
      fetchData();

    },[props.isAuth, props.username]);
    return (
        <div>
        <Header status = {props.isAuth}/>
        <Body createPost = {false} posts={posts} isAuth={props.isAuth}/>
        <div>
        </div>
        </div>
    );
};

export default HomePage;