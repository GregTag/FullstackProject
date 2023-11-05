import { Link } from "react-router-dom";
import { selectResults } from "../logic/slices/search";
import { useSelector } from "react-redux";

function SeachResult({ result }) {
    return (
        <div key={result.id} className="search-result">
            <Link to={`/media/${result.id}`}>
                <div className="search-result-image">
                    <img src={result.image} alt={result.title} />
                </div>
                <div className="search-result-info">
                    <div className="search-result-title">
                        {result.title}
                    </div>
                </div>
            </Link>
        </div>
    );
}

function SearchPage() {
    const results = useSelector(selectResults);
    return (
        <div className="block">
            <div className="side">

            </div>
            <div className="other">
                {results.map((result) => <SeachResult result={result} />)}
            </div>
        </div>
    );
}

export default SearchPage;
