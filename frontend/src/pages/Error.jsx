import { Link as RouterLink, useRouteError } from "react-router-dom"
import { Card, Typography, Stack, Link } from '@mui/joy';

function ErrorPage() {
    const error = useRouteError();
    console.log(error);

    return (
        <Stack direction="column" sx={{ alignItems: "center", justifyContent: "center", height: "100vh" }}>
            <Card sx={{ p: 4, textAlign: 'center' }}>
                <Typography level="h1">{error.status} {error.statusText}</Typography>
                <Link component={RouterLink} to="/" overlay underline="none">
                    <Typography level="h2">Go to Home page</Typography>
                </Link>
            </Card>
        </Stack >
    )
}

export default ErrorPage;
