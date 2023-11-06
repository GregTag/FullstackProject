import { useLoaderData } from "react-router-dom";
import { Sheet, Stack, AspectRatio, Typography, Box, Divider } from "@mui/joy";
import UserLists from "../common/UserLists";

function ProfilePage() {
    const profile = useLoaderData();
    console.log("Profile", profile);
    return (
        <Sheet sx={{ mx: { sm: 0, md: 25 }, p: "20px" }}>
            <Stack spacing={2}>
                <Stack direction={{ xs: "column", sm: "row" }} spacing={2}>
                    <Box sx={{ flexBasis: 300 }}>
                        <AspectRatio ratio={1 / 1} >
                            <img src={profile.avatar} alt={profile.login} />
                        </AspectRatio>
                    </Box>
                    <Stack direction="column" spacing={2}>
                        <Typography level="h1">{profile.fullname}</Typography>
                        <Divider />
                    </Stack>
                </Stack>
                <UserLists profile={profile} />
            </Stack>
        </Sheet >
    );
}

export default ProfilePage;
