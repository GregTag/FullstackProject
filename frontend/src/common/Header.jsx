import React from 'react';
import { Form, Link as RouterLink } from 'react-router-dom';
import { useDispatch, useSelector } from 'react-redux';
import { selectIsLoggedIn, logout } from '../logic/slices/user';
import { Sheet, Box, Typography, Input, Button, Link, FormControl, Stack } from '@mui/joy';

function Header() {
    const is_logged = useSelector(selectIsLoggedIn);
    const dispatch = useDispatch();
    return (
        <Sheet component="header" variant='soft' sx={{ p: 2 }}>
            <Stack direction={{ xs: 'column', sm: 'row' }} sx={{ alignItems: 'center' }} spacing={2} >
                <Link component={RouterLink} to="/" underline='none'>
                    <Typography level="h2">
                        Universal Media Organizer
                    </Typography>
                </Link>
                <FormControl sx={{ flexGrow: 1 }} size="lg" color="neutral" component={Form} action="/search" method="post">
                    <Input name='query' placeholder="Search..." variant="outlined" endDecorator={
                        <Button size="lg" type="submit" color="neutral" variant="soft">Search</Button>
                    } />
                </FormControl>
                {is_logged
                    ? (
                        <Box>
                            <Button size="lg" color="neutral" variant="soft" component={RouterLink} to="/profile">My Profile</Button>
                            <Button size="lg" color="neutral" onClick={() => dispatch(logout())} variant="soft">
                                Logout
                            </Button>
                        </Box>
                    )
                    : (
                        <Button size="lg" color="neutral" component={RouterLink} to="/auth" variant="soft">Login</Button>
                    )
                }
            </Stack>
        </Sheet >
    );
}

export default Header;
