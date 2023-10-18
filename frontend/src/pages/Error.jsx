import { Link, useRouteError } from "react-router-dom"

function ErrorPage() {
    const error = useRouteError();
    console.log(error);

    return (
        <Link to="/">
            <div className="just-text">
                <h1>{error.status} {error.statusText}</h1>
                <h2>Go to Home page</h2>
            </div>
        </Link>
    )
}

export default ErrorPage;
