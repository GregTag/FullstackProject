import { useLoaderData } from 'react-router-dom';
import Comments from '../common/Comments';
import React from 'react';
import Rating from '../common/Rating';
import { Sheet, Stack, AspectRatio, Typography, Box, Divider } from '@mui/joy';
import StatusMenu from '../common/StatusMenu';

function MediaPage() {
    const media = useLoaderData();

    return (
        <Sheet sx={{ mx: { sm: 0, md: 25 }, p: '20px', flexGrow: 1 }}>
            <Stack spacing={2}>
                <Stack direction={{ xs: 'column', sm: 'row' }} spacing={2}>
                    <Box sx={{ flexBasis: 300 }}>
                        <AspectRatio ratio={1 / 1} >
                            <img src={media.image} alt={media.title} />
                        </AspectRatio>
                    </Box>
                    <Stack direction="column" spacing={2}>
                        <Typography level="h1">{media.title}</Typography>
                        <Divider />
                        <Typography>{media.description}</Typography>
                        <Stack direction="row" spacing={2}>
                            <Rating rating={media.rating} media_id={media.id} />
                            <StatusMenu media_id={media.id} />
                        </Stack>
                    </Stack>
                </Stack>
                <Comments media={media.id} />
            </Stack>
        </Sheet>
    );
}

export default MediaPage;
