import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import './App.css';

import Root from './common/Root';
import HomePage from './pages/Home';
import MediaPage from './pages/Media';
import ProfilePage from './pages/Profile';
import SearchPage from './pages/Search';
import ErrorPage from './pages/Error';
import AuthPage from './pages/Auth';

const router = createBrowserRouter([
  {
    path: '/', element: <Root />, errorElement: <ErrorPage />, children: [
      { path: '/', element: <HomePage />, errorElement: <ErrorPage /> },
      { path: '/media', element: <MediaPage />, errorElement: <ErrorPage /> },
      { path: '/profile', element: <ProfilePage />, errorElement: <ErrorPage /> },
      { path: '/search', element: <SearchPage />, errorElement: <ErrorPage /> }
    ]
  },
  { path: '/auth', element: <AuthPage />, errorElement: <ErrorPage /> },

]);

function App() {
  return (
    <RouterProvider router={router} />
  );
}

export default App;
