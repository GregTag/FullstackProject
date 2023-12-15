import React from 'react';
import { useLoaderData } from 'react-router-dom';
import { Sheet, Stack, AspectRatio, Typography, Box, Divider } from '@mui/joy';
import UserLists from '../common/UserLists';

const default_avatar = 'https://cdn.icon-icons.com/icons2/2468/PNG/512/user_icon_149329.png';

function ProfilePage() {
    const profile = useLoaderData();
    return (
        <Sheet sx={{ mx: { sm: 0, md: 25 }, p: '20px' }}>
            <Stack spacing={2}>
                <Stack direction={{ xs: 'column', sm: 'row' }} spacing={2}>
                    <Box sx={{ flexBasis: 300 }}>
                        <AspectRatio ratio={1 / 1} >
                            <img src={profile.user.avatar || default_avatar} alt={profile.user.login} />
                        </AspectRatio>
                    </Box>
                    <Stack direction="column" spacing={2}>
                        <Typography level="h1">{profile.user.fullname}</Typography>
                        <Divider />
                    </Stack>
                </Stack>
                {profile.tracks && <UserLists tracks={profile.tracks} />}
            </Stack>
        </Sheet >
    );
}

export default ProfilePage;
