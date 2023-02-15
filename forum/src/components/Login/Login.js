import React, { useEffect, useState } from "react";
import { Link, useNavigate } from "react-router-dom";
import { Alert } from "@mui/material";
import './Login.css'
const Login = () => {
    
    useEffect(()=> {   
            const inputText = document.querySelectorAll('.auth-form__input');
            
            inputText.forEach( function(input) {
                    input.addEventListener('focus', function() {
                        this.classList.add('focus');                            this.parentElement.querySelector('.auth-form__placeholder').classList.add('focus');
                    });
                    input.addEventListener('blur', function() {
                        this.classList.remove('focus');
                        if (! this.value) {
                            this.parentElement.querySelector('.auth-form__placeholder').classList.remove('focus');
                        }
                    });
            });
    
            const togglers = document.querySelectorAll('.password-toggler');

            togglers.forEach( function(checkbox) {
                checkbox.addEventListener('change', function() {
        
                    const toggler = this.parentElement,
                            input   = toggler.parentElement.querySelector('.input-password'),
                            icon    = toggler.querySelector('.auth-form__icon');
        
                    if (checkbox.checked) {
                        input.type = 'text';
                        icon.classList.remove('la-eye')
                        icon.classList.add('la-eye-slash');
                    }
        
                    else
                    {
                        input.type = 'password';
                        icon.classList.remove('la-eye-slash')
                        icon.classList.add('la-eye');
                    }
                });
            });
    },[]);
    const navigate = useNavigate();

    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [status, setStatus] = useState(null);
    // useEffect(() => {
    //     console.log(temp.value);
    // }, [temp.value])
    const handleUsername = (e) => {
        setUsername(e.target.value);
    }
    const handlePassword = (e) => {
        setPassword(e.target.value);
    }
    const handleStatus = () => {
        if (status === "Server-error"){
            return (  
                <Alert severity="error">
                    User doesn't exist
                </Alert>)
        }
    }

    const sendForm = (e) => {
        e.preventDefault();
        // setUsername(e.target.username.value);
        // setPassword(e.target.password.value);
        // console.log(username);
        // console.log(password);
// Check for Mail and password validity
        (async() => {
            await fetch(`http://localhost:8080/api/signin`, 
            {
                headers: {
                    'Accept': 'text/plain',
                    'Content-type': 'text/plain',
                    'Credentials': 'include'
                },
                method: "POST",
                credentials: 'include',
                body: JSON.stringify({
                    Username: username,
                    password: password
                }),
            }).then((r) =>{
                if (!r.ok){
                    navigate("/")
                }
                return (r.json())
            })
            .then((data) => {
                const responseJSON = data;
                const {Username} = responseJSON;
                sessionStorage.setItem("Username", Username);
                navigate("/", 
                // {
                //     state: {
                //         username: username
                //     }
                // }
                );
            
            })
        })();
    }
    
    return (
        <div className="modal__background">
        <div className="modal__window">
    
            <form className="auth-form" name="form-auth" onSubmit={sendForm}>
    
                <label className="auth-form__label">
                    <span className="auth-form__placeholder">username</span>
                    <input className="auth-form__input input-email" type="username" value={username} onChange={handleUsername} name="username" autoComplete="off" required/>
                </label>
    
                <label className="auth-form__label">
                    <span className="auth-form__placeholder">password</span>
                    <input className="auth-form__input input-password" type="password" value={password} onChange={handlePassword} name="password" autocomlete="off" required/>
                    <div className="auth-form__toggler">
                        <i className="la la-eye auth-form__icon"></i>
                        <input type="checkbox" className="auth-form__checkbox password-toggler"/>
                    </div>
                </label>
    
                <div className="auth-form__answer">{handleStatus()}</div>
    
                <input className="auth-form__submit" type="submit" value="Login"/>
                
                <div className="auth-form__bottom">
                    <span>Have no account? </span>
                    <Link to = "/signup">
                    Create new
                    </Link>
                </div>
            </form>
    
        </div>
    </div>
    
    )
}

export default Login;