import { createBrowserRouter, RouterProvider } from 'react-router-dom';
import './App.css';

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
    path: '/', element: <Root />, errorElement: <ErrorPage />, children: [
      { path: '/', element: <HomePage />, errorElement: <ErrorPage /> },
      { path: '/media/:name?', loader: loadMedia, action: actionComment, element: <MediaPage />, errorElement: <ErrorPage /> },
      { path: '/profile/:name?', loader: loadProfile, element: <ProfilePage />, errorElement: <ErrorPage /> },
      { path: '/search', loader: actionSearch, action: actionSearch, element: <SearchPage />, errorElement: <ErrorPage /> }
    ]
  },
  { path: '/auth', action: actionAuth, element: <AuthPage />, errorElement: <AuthPage /> },

]);

function App() {
  return (
    <RouterProvider router={router} />
  );
}

export default App;
