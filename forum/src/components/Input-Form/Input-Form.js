import React, { useState } from 'react';
import './Input.css'
// import { document } from 'global/document';
// function getCookie(cname) {
//     const name = cname + "=";
//     const decodedCookie = decodeURIComponent(document.cookie);
//     const ca = decodedCookie.split(';');
//     for(let i = 0; i <ca.length; i++) {
//       let c = ca[i];
//       while (c.charAt(0) === ' ') {
//         c = c.substring(1);
//       }
//       if (c.indexOf(name) === 0) {
//         return c.substring(name.length, c.length);
//       }
//     }
//     return "";
//   }
  
async function handleSubmit(event, text) {
    event.preventDefault();
    // const cookie = document.cookie.split(';').find(c => c.trim().startsWith('token='));

    // if (!cookie) {
    //   console.error('No token found in cookie');
    //   return;
    // }
    // const token = cookie.split('=')[1];
    const token = document.cookie
    console.log(token);
    try {
      const response = await fetch('http://localhost:8080/api/createPost', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        //   'Accept': 'text/plain',
          'Cookie': `${token}`,
        },
        body: JSON.stringify({ text: text }),
      });
      const data = await response.json();
      console.log(data);
    } catch (error) {
      console.error(error);
    }
  };


  function InputForm() {
    const [text, setText] = useState('');
  
    return (
      <form className='input-form' onSubmit={(e) => handleSubmit(e, text)}>
        <textarea value={text} onChange={(e) => setText(e.target.value)} />
        <button type="submit">Submit</button>
      </form>
    );
  }
export default InputForm;