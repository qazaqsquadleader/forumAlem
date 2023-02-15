import React, {useState, useEffect} from 'react';
import Header from '../../components/Header/Header';
import Body from '../../components/Body/Body';


const CreatePost = (props) => {

// const [user, setUser] = useState();    

// const [isAuth, setIsAuth] = useState();
// const [isLoading, setIsLoading] = useState(true);

// useEffect(() => {
//   if (props.isAuth){

//     setIsAuth(props.setIsAuth)
//       setUser(props.username)
//       setIsLoading(false)
//       console.log("CreatePost:"+props.username);

//     }   
// }, [props.isAuth, props.username])  

    // return isLoading ? (

    //     <div>Loading...</div>
    //   ) : (
    //     <div>
    //     <Header />
    //     <Body createPost={true} username={user} isAuth={isAuth}/>
    //     </div>
    //   );
    console.log("CreatePOST:", props.username);
    return(
      <div>
        <Header status = {props.isAuth}/>
         <Body createPost={true} username={props.username} isAuth={props.isAuth}/>
         </div>
    )
    
};

export default CreatePost;