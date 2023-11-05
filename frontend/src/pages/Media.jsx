import { useLoaderData } from "react-router-dom";
import './Media.css';
import Comments from "../common/Comments";
import Rating from "../common/Rating";

function MediaPage() {
    const media = useLoaderData();

    return (
        <div className="page">
            <div className="block">
                <div className="side">
                    <img className="media-image" src={media.image} alt={media.title} />
                    <Rating rating={media.rating} />
                </div>
                <div className="other">
                    <h1>{media.title}</h1>
                    <p>{media.description}</p>
                </div>
            </div>
            <Comments media={media.id} />
        </div>
    );
}

export default MediaPage;
