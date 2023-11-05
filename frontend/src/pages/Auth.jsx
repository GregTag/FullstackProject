import './Auth.css'
import { Form, Link, useRouteError } from 'react-router-dom';
import { useState } from 'react';

function LoginForm({ onRegister }) {
    return (
        <Form action='/auth' method='post'>
            <label htmlFor='username'>Username:</label>
            <input type='text' name='username' />
            <label htmlFor='password'>Password:</label>
            <input type='password' name='password' />
            <div className='row-container'>
                <button type='submit' className='button'>Login</button>
                <button onClick={onRegister} className='button'>Register</button>
            </div>
        </Form >
    )
}

function RegisterForm({ onGoBack }) {
    return (
        <Form action='/auth' method='post'>
            <label htmlFor='username'>Username:</label>
            <input type='text' name='username' />
            <label htmlFor='password'>Password:</label>
            <input type='password' name='password' />
            <label htmlFor='repeated'>Repeat Password:</label>
            <input type='password' name='repeated' />
            <div className='row-container'>
                <button type='submit' className='button'>Register</button>
                <button onClick={onGoBack} className='button'>Go Back</button>
            </div>
        </Form>
    )
}


function AuthPage() {
    const error = useRouteError();
    if (error) {
        console.error(error);
    }
    const [registering, setRegistering] = useState(false);
    return (
        <div className='central-panel'>
            <Link to="/"><h1>Universal Media Organizer</h1></Link>
            {error && <div className='error'>Error!</div>}
            {registering ? <RegisterForm onGoBack={() => setRegistering(false)} /> : <LoginForm onRegister={() => setRegistering(true)} />}
        </div>
    )
}

export default AuthPage;
