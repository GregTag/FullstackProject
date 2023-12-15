import axios from 'axios';
import { Configuration, UserApi, MediaApi, CommentApi, TrackApi } from '../api/index';

axios.interceptors.request.use(
    (config) => {
        const auth = localStorage.getItem('auth');
        if (auth) {
            config.headers.setAuthorization(auth);
        }
        console.log('Making request', config);
        return config;
    }
);

axios.interceptors.response.use((response) => {
    if (response.status === 401) {
        localStorage.removeItem('auth');
    }
    const auth_header = response.headers.getAuthorization();
    if (auth_header) {
        localStorage.setItem('auth', auth_header);
    }
    console.log('Got response', response);
    return response;
});

const config = new Configuration({
    // basePath: window.location.origin,
    basePath: 'http://localhost:8000',
});

export const userApi = new UserApi(config, config.basePath, axios);
export const mediaApi = new MediaApi(config, config.basePath, axios);
export const commentApi = new CommentApi(config, config.basePath, axios);
export const trackApi = new TrackApi(config, config.basePath, axios);
