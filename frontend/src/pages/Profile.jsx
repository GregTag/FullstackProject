import { useLoaderData } from "react-router-dom";

function ProfilePage() {
    const profile = useLoaderData();

    return (
        <div className="page">
            <div className="block">
                <div className="side">
                    <img className="media-image" src={profile.avatar} alt={profile.name} />
                </div>
                <div className="other">
                    <h1>{profile.name}</h1>
                    Bruh
                </div>
            </div>
        </div>
    );
}

export default ProfilePage;
