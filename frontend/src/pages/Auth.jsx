import React, { useState } from 'react';
import { Form, Link as RouterLink, useRouteError } from 'react-router-dom';
import { Stack, Link, Card, FormControl, Input, Button, Typography, Alert, FormLabel } from '@mui/joy';

function LoginForm({ onRegister }) {
    return (
        <Form action='/auth' method='POST'>
            <FormControl>
                <FormLabel>Username:</FormLabel>
                <Input type='text' name='username' />
            </FormControl>
            <FormControl>
                <FormLabel>Password:</FormLabel>
                <Input type='password' name='password' />
            </FormControl>
            <Stack direction="row" sx={{ justifyContent: 'center', pt: 2 }} spacing={2}>
                <Button size='lg' type='submit'>Login</Button>
                <Button size='lg' onClick={onRegister}>Register</Button>
            </Stack>
        </Form>
    );
}

function RegisterForm({ onGoBack }) {
    return (
        <Form action='/auth' method='POST'>
            <FormControl>
                <FormLabel htmlFor='username'>Username:</FormLabel>
                <Input type='text' name='username' />
            </FormControl>
            <FormControl>
                <FormLabel htmlFor='password'>Password:</FormLabel>
                <Input type='password' name='password' />
            </FormControl>
            <FormControl>
                <FormLabel htmlFor='repeated'>Repeat Password:</FormLabel>
                <Input type='password' name='repeated' />
            </FormControl>
            <Stack direction="row" sx={{ justifyContent: 'center', pt: 2 }} spacing={2}>
                <Button size='lg' type='submit'>Register</Button>
                <Button size='lg' onClick={onGoBack}>Go Back</Button>
            </Stack>
        </Form>
    );
}

function AuthPage() {
    const error = useRouteError();
    if (error) {
        console.error(error);
    }
    const [registering, setRegistering] = useState(false);
    return (
        <Stack direction="column" sx={{ alignItems: 'center', justifyContent: 'center', height: '100vh' }}>
            <Card sx={{ p: 4 }}>
                <Link component={RouterLink} to="/" underline='none' sx={{ justifyContent: 'center' }}>
                    <Typography level="h2">Universal Media Organizer</Typography>
                </Link>
                {error && <Alert color="danger">Error!</Alert>}
                {registering ? <RegisterForm onGoBack={() => setRegistering(false)} /> : <LoginForm onRegister={() => setRegistering(true)} />}
            </Card>
        </Stack>
    );
}

export default AuthPage;
