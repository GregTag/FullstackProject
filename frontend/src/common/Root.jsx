import { Outlet } from "react-router-dom";
import Header from "./Header";
import { Stack } from "@mui/joy";

function Root() {
    return (
        <div>
            <Stack direction="column" sx={{ minHeight: "100dvh" }}>
                <Header />
                <Outlet />
            </Stack>
        </div>
    )
}

export default Root;
