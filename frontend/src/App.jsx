import React from 'react';
import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import { CssVarsProvider, extendTheme } from '@mui/joy/styles';
import { CssBaseline } from '@mui/joy';
import { PersistGate } from 'redux-persist/integration/react';
import store, { persistor } from './logic/slices/store';
import { Provider } from 'react-redux';

import Root from './common/Root';
import HomePage from './pages/Home';
import MediaPage from './pages/Media';
import ProfilePage from './pages/Profile';
import SearchPage from './pages/Search';
import ErrorPage from './pages/Error';
import AuthPage from './pages/Auth';
import { loadMedia, loadProfile } from './logic/loaders';
import { actionAuth, actionComment, actionSearch } from './logic/actions';

const router = createBrowserRouter([
    {
        path: '/',
        element: <Root />,
        errorElement: <ErrorPage />,
        children: [
            { path: '/', element: <HomePage />, errorElement: <ErrorPage /> },
            { path: '/media/:id', loader: loadMedia, action: actionComment, element: <MediaPage />, errorElement: <ErrorPage /> },
            { path: '/profile/:name?', loader: loadProfile, element: <ProfilePage />, errorElement: <ErrorPage /> },
            { path: '/search', loader: actionSearch, action: actionSearch, element: <SearchPage />, errorElement: <ErrorPage /> }
        ]
    },
    { path: '/auth', action: actionAuth, element: <AuthPage />, errorElement: <AuthPage /> }

]);

const theme = extendTheme({
    fontFamily: {
        body: '\'Oswald\', sans-serif',
        display: '\'Oswald\', sans-serif'
    }
});

function App() {
    return (
        <Provider store={store}>
            <PersistGate loading={null} persistor={persistor}>
                <CssVarsProvider defaultMode="dark" theme={theme} disableNestedContext>
                    <CssBaseline />
                    <RouterProvider router={router} />
                </CssVarsProvider>
            </PersistGate>
        </Provider>
    );
}

export default App;
