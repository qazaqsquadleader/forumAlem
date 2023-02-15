import React, { useEffect, useState } from 'react';
import './Input.css'
import { TextField, MenuItem } from '@mui/material';
async function handleSubmit(event, text,title, name) {
    event.preventDefault();
    try {
      const response = await fetch('http://localhost:8080/api/createPost', {
        method: 'POST',
        credentials: 'include',
        headers: {
          'Content-Type': 'application/json',
        //   'Accept': 'text/plain',
        },
        body: JSON.stringify({ 
          Content: text,
          Title: title,
          Author: name,
        }),
      });
      const data = await response.json();
      console.log(data);
    } catch (error) {
      console.error(error);
    }
  };

  const currencies = [
    {
      value: 'USD',
      label: '$',
    },
    {
      value: 'EUR',
      label: 'SOME CATEGORIES FOR THE FUTURE',
    },
    {
      value: 'BTC',
      label: '฿',
    },
    {
      value: 'JPY',
      label: '¥',
    },
  ];
const InputForm = (props) => {
    const [text, setText] = useState('');
    const [title, setTitle] = useState('')
    // const [user, setUser] = useState();    
    // console.log(props.username);

// const [isAuth, setIsAuth] = useState();
    // useEffect(() => {
    //   if (props.isAuth){
    //     setIsAuth(props.isAuth)
    //       setUser(props.username)
    //       console.log(user);
    //     }   
    // }, [props.isAuth, props.username]) 
    
    return (
      <form className='input-form' onSubmit={(e) => handleSubmit(e, text,title, props.username)}>
        <div className='input-group'>
          <h1>Create Post</h1>  
        </div>
        <div className='text-field-group'>
          <TextField id="title-form" className='mu-textfield' label="Title" variant="outlined" value={title} onChange={(e) => setTitle(e.target.value)} />
          <TextField
          id="outlined-select-currency"
          select
          label=""
          defaultValue="EUR"
          sx={{backgroundColor: 'aliceblue'}}
        >
          {currencies.map((option) => (
            <MenuItem key={option.value} value={option.value}>
              {option.label}
            </MenuItem>
          ))}
        </TextField>
        </div> 
        <div className='textarea-group'>
          <textarea value={text} onChange={(e) => setText(e.target.value)} />
          <button type="submit">Submit</button>
        </div>
      </form>
    );
  }
export default InputForm;