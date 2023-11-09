import React from 'react';
import { Link as RouterLink } from 'react-router-dom';
import { Sheet, Button, Typography, Stack, Divider } from '@mui/joy';
import { useDispatch } from 'react-redux';
import { setFilter } from '../logic/slices/search';
import mediaAttributes from '../config/media_attributes.json';

function SearchButton({ to, category }) {
    const dispatch = useDispatch();
    const lower = category.toLowerCase();
    const handler = () => dispatch(setFilter({ filter: 'category', value: lower }));
    return (
        <Button component={RouterLink} to={to} size='lg' onClick={handler}>{category}</Button>
    );
}

function HomePage() {
    return (
        <Sheet sx={{ mx: { sm: 0, md: 25 }, p: '20px', flexGrow: 1 }}>
            <Stack spacing={2}>
                <Typography level="h2">Home Page</Typography>
                <Divider />
                <Typography>
                    Universal Media Organizer (UMO) is a comprehensive media management platform designed to help you organize and track your favorite content across various forms of media. Whether it&apos;s movies, series, podcasts, audiobooks, or even online lectures, UMO provides a unified platform to manage them all.
                </Typography>
                <Typography>
                    You can find new content using the search page. With flexible filter system you will definitely find something for yourself.
                </Typography>
                <Stack direction="row" spacing={2}>
                    {mediaAttributes.category.values.map((category) => (
                        <SearchButton key={category} to="/search" category={category} />
                    ))}
                </Stack>
                <Typography>
                    With UMO, you can track the progress of viewing content, observe the history and statistics of views, rate, post comments and most importantly share your favorite content with others.
                </Typography>
                <Typography>
                    UMO is built with a user-friendly interface and is designed to work across different devices, providing you with a seamless media management experience no matter where you are.
                </Typography>
                <Typography>
                    In summary, Universal Media Organizer is not just another media tracking tool. It&apos;s a comprehensive solution designed to revolutionize the way you manage and consume your favorite media content.
                </Typography>
            </Stack>
        </Sheet>
    );
}

export default HomePage;
