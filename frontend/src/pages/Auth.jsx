import './Auth.css'
import { login } from '../logic/user'
import { useDispatch } from 'react-redux';
import { Link, useNavigate } from 'react-router-dom';
import { useState } from 'react';

function Login({ onRegister }) {
    const dispatch = useDispatch();
    const navigate = useNavigate();
    return (
        <form onSubmit={(event) => {
            event.preventDefault();
            // const { username, password } = event.target.elements;
            dispatch(login());
            navigate('/profile');
        }}>
            <label htmlFor='username'>Username:</label>
            <input type='text' name='username' />
            <label htmlFor='password'>Password:</label>
            <input type='password' name='password' />
            <div className='row-container'>
                <button type='submit' className='button'>Login</button>
                <button onClick={onRegister} className='button'>Register</button>
            </div>
        </form >
    )
}

function Register() {
    const dispatch = useDispatch();
    const navigate = useNavigate();
    return (
        <form onSubmit={(event) => {
            event.preventDefault();
            const { password, repeated } = event.target.elements;
            if (password.value !== repeated.value) {
                alert('Passwords do not match!');
                return;
            }

            dispatch(login());
            navigate('/profile');
        }}>
            <label htmlFor='username'>Username:</label>
            <input type='text' name='username' />
            <label htmlFor='password'>Password:</label>
            <input type='password' name='password' />
            <label htmlFor='repeated'>Repeat Password:</label>
            <input type='password' name='repeated' />
            <div className='row-container'>
                <button type='submit' className='button'>Register</button>
            </div>
        </form>
    )
}


function AuthPage() {
    const [registering, setRegistering] = useState(false);
    return (
        <div className='central-panel'>
            <Link to="/"><h1>Universal Media Organizer</h1></Link>
            {registering ? <Register /> : <Login onRegister={() => setRegistering(true)} />}
        </div>
    )
}

export default AuthPage;
