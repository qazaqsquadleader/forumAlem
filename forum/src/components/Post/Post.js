import React, { useState } from 'react';
import './Post.css'


  function Post(props) {
    // const [text, setText] = useState('');
    return (
      <div className='post-container'>
        <div className='post-header'>
        <p>{props.date}</p>
        <p>{props.title}</p>
        <p>{props.author}</p>
        </div>
        
        <div className='post-body'>
            <p>{props.content}</p>
        </div>

      </div>
    );
  }
export default Post;