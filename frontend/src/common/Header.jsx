import { Form, Link } from 'react-router-dom';
import { useSelector } from 'react-redux';
import { selectIsLoggedIn } from '../logic/user';
import './Header.css';

function Header() {
    const isLoggedIn = useSelector(selectIsLoggedIn);
    return (
        <header>
            <Link to="/"><h2 className='title'>Universal Media Organizer</h2></Link>
            <Form className='search-bar'>
                <input type="text" placeholder="Search..." />
                <button type="submit" className='button'>Search</button>
            </Form>
            {isLoggedIn ? (
                <Link to="/profile" className='button'>My Profile</Link>
            ) : (
                <Link to="/auth" className='button'>Login</Link>
            )}

        </header >
    );
};

export default Header;
