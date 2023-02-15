import React,{useEffect, useState} from "react";
import InputForm from "../Input-Form/Input-Form";
import './Body.css'
import Post from "../Post/Post";
const Body = (props) => {  

// props.posts ? console.log("exists"): console.log("no");
// const handlePosts =()=>{
//   if (props.posts){
//     props.posts.map(({PostId, Title, Content, CreationDate, Author})=>{
//       return (<Post key={PostId} title={Title} content={Content} date={CreationDate} author={Author}/>) 
//     })
//   }
//   return (<p>asd</p>)
// }
//    CONSIDER UNCOMMENTING
// const [user, setUser] = useState();    
// const [isAuth, setIsAuth] = useState();
// useEffect(() => {
//   if (props.isAuth){
//     setIsAuth(props.isAuth)
//       setUser(props.username)
//     }   
// }, [props.isAuth, props.username, props.posts]) 
const handlePosts = () => {
  
  if (!props.createPost && (props.posts.length!== 1 && Object.keys(props.posts[0]).length !== 0)){
    console.log(props.posts);
  return(
    props.posts.map(({PostId, Title, Content, CreationDate, Author})=>{
      return <Post key={PostId} title={Title} content={Content} date={CreationDate} author={Author}/>
    })
    )
  }
  
}

return (
    <div className="body-container">
      <div className="body-posts-container">
      {props.createPost && <InputForm username= {props.username}/>}    
      {/* DEBUG */}
      {/* {!props.createPost && (props.posts.length !== 1 && Object.keys(props.posts[0].length)!==0 ) && (props.posts.map(({PostId, Title, Content, CreationDate, Author})=>{return (<Post key={PostId} title={Title} content={Content} date={CreationDate} author={Author}/>)}))} */}
      {handlePosts()}
      </div>
    </div>
  )
};

export default Body;