import { Link } from 'react-router-dom';
import './Home.css';

function HomePage() {
    return (
        <div className="page">
            <h2>Home Page</h2>
            <p>
                Universal Media Organizer (UMO) is a comprehensive media management platform designed to help you organize and track your favorite content across various forms of media. Whether it's movies, series, podcasts, audiobooks, or even online lectures, UMO provides a unified platform to manage them all.
            </p>
            <p>
                You can find new content using the search page. With flexible filter system you will definitely find something for yourself.
            </p>
            <div className='row-container'>
                <Link to='/search' className='item button'>Serials</Link>
                <Link to='/search' className='item button'>Movies</Link>
                <Link to='/search' className='item button'>Books</Link>
                <Link to='/search' className='item button'>Lectures</Link>
            </div>
            <p>
                With UMO, you can track the progress of viewing content, observe the history and statistics of views, rate, post comments and most importantly share your favorite content with others.
            </p>
            <p>
                UMO is built with a user-friendly interface and is designed to work across different devices, providing you with a seamless media management experience no matter where you are.
            </p>
            <p>
                In summary, Universal Media Organizer is not just another media tracking tool. It's a comprehensive solution designed to revolutionize the way you manage and consume your favorite media content.
            </p>
        </div>
    )
}

export default HomePage;
