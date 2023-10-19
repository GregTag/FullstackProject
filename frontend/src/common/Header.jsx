import { Form, Link } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { selectIsLoggedIn, logout } from '../logic/user';
import './Header.css';

function Header() {
    const isLoggedIn = useSelector(selectIsLoggedIn);
    const dispatch = useDispatch();
    return (
        <header>
            <Link to="/"><h2 className='title'>Universal Media Organizer</h2></Link>
            <Form action='/search' method='GET' className='search-bar'>
                <input type="text" name='query' placeholder="Search..." />
                <button type="submit" className='button'>Search</button>
            </Form>
            {isLoggedIn ? (
                <div className='adaptive-flex'>
                    <Link to="/profile" className='button'>My Profile</Link>
                    <button onClick={() => dispatch(logout())} className='button'>
                        <span class="material-symbols-outlined">logout</span>
                    </button>
                </div>
            ) : (
                <Link to="/auth" className='button'>Login</Link>
            )}

        </header >
    );
};

export default Header;
