import React, { useState, useEffect } from "react";
import { Alert } from "@mui/material";
import { Link } from "react-router-dom";

const SignUp = () => {
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

    
    const [mail, setMail] = useState('');
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [status, setStatus] = useState(null);
    const handleMail = (e) => {
        setMail(e.target.value);
    }
    const handleUsername = (e) => {
        setUsername(e.target.value);
    }
    const handlePassword = (e) => {
        setPassword(e.target.value);
    }
    
    const handleStatus = () => {
        if (status === "Sign-up success"){
            return (
                <Alert severity="success">
                    Success!
                </Alert>)
        } else if (status === "Incorrect input. Try again"){
            return (  
                <Alert severity="error">
                    Username taken. Try again
                </Alert>)
        }
    }


    const sendForm = async (e) => {
        e.preventDefault();
        
// Check for Mail and password validity
        // (async() => {
            await fetch(`http://localhost:8080/api/signup`, 
            {
                headers: {
                    'Accept': 'text/plain',
                    'Content-type': 'text/plain'
                },
                method: "POST",
                body: JSON.stringify({
                    Email: mail,
                    Username: username,
                    password: password
                })
            }).then((r) => {
                setStatus("Sign-up success")
            
                if (!r.ok){
                    setStatus("Incorrect input. Try again")
                }
                
                
            })
            
        // })();
    }

    return (
        <div className="modal__background">
        <div className="modal__window">
    
            <form className="auth-form" name="form-auth" onSubmit={sendForm}>
    
                <label className="auth-form__label">
                    <span className="auth-form__placeholder">email</span>
                    <input className="auth-form__input input-email" type="email" name="email" value={mail} onChange={handleMail} autoComplete="off" required/>
                </label>
                <label className="auth-form__label">
                    <span className="auth-form__placeholder">username</span>
                    <input className="auth-form__input input-email" type="username" name="username" value={username} onChange={handleUsername} autoComplete="off" required/>
                </label>
                <label className="auth-form__label">
                    <span className="auth-form__placeholder">password</span>
                    <input className="auth-form__input input-password" type="password" name="password" value={password} onChange={handlePassword} autocomlete="off" required/>
                    <div className="auth-form__toggler">
                        <i className="la la-eye auth-form__icon"></i>
                        <input type="checkbox" className="auth-form__checkbox password-toggler"/>
                    </div>
                </label>
                
                <div className="auth-form__answer">{handleStatus()}</div>
              
{/*     
                <div className="auth-form__answer">{status}</div>
     */}
                <input className="auth-form__submit" type="submit" value="Sign Up"/>
                
                <div className="auth-form__bottom">
                    <span>Already have an account? </span>
                    <Link to = "/signin">
                    Sign in
                    </Link>
                </div>
            </form>
    
        </div>
    </div>
    
    )
}   

export default SignUp;