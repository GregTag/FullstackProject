import config from '../config/config.json';
import { login } from './slices/user';
import { json, redirect } from 'react-router-dom';
import { dispatch } from './slices/store';
import { fetchSearchResults } from './slices/search';

export async function actionAuth({ request }) {
    const data = await request.formData();
    if (!data.get('password')) {
        return json({ error: 'Password is required' }, { status: 400 });
    }
    const response = await fetch(`${config.api_base_url}/api/users/${data.get('username')}`);
    const user = await response.json();
    dispatch(login(user[0]));
    return redirect('/profile');
}

export async function actionComment({ request }) {
    const data = await request.formData();
    const content = data.get('content');
    if (!content) {
        return {}; // ignore
    }
    const response = await fetch(`${config.api_base_url}/api/comment/add` /* options */);
    return response;
}

export async function actionSearch() {
    dispatch(fetchSearchResults());
    return {};
}
