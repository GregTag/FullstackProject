import { Link as RouterLink } from "react-router-dom";
import { selectResults } from "../logic/slices/search";
import { useSelector } from "react-redux";
import { Sheet, Stack, Typography, Grid, Card, CardOverflow, AspectRatio, Link } from "@mui/joy";
import FilterComponent from "../common/Filters";

function SidePanel() {
    return (
        <Sheet variant="soft" sx={{ minWidth: { sm: "100vw", md: 300 }, maxWidth: { sm: "100vw", md: 300 } }} >
            <Typography level="h3" sx={{ textAlign: "center" }}>Filters</Typography>
            <FilterComponent />
        </ Sheet>
    )
}

function Result({ result }) {
    return (
        <Grid xs={6} sm={4} lg={3} xl={2} sx={{ p: 2 }}>
            <Card variant="soft">
                <CardOverflow>
                    <AspectRatio ratio={1 / 1}>
                        <img src={result.image} alt={result.title} />
                    </AspectRatio>
                </CardOverflow>
                <Link component={RouterLink} to={`/media/${result.id}`} overlay underline="none">
                    <Typography level="h4">{result.title}</Typography>
                </Link>
            </Card>
        </Grid >
    )
}

function MainGrid() {
    const results = useSelector(selectResults);
    return (<Grid container sx={{ flexGrow: 1 }}>
        {results.map((result) => (<Result key={result.id} result={result} />))}
    </Grid>)
}

function SearchPage() {

    return (
        <Stack direction={{ sm: "column", md: "row" }} sx={{ flexGrow: 1 }}>
            <SidePanel />
            <MainGrid />
        </Stack>

    );
}

export default SearchPage;
