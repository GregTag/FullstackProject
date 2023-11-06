import config from '../config/config.json'
import { login } from "./slices/user";
import { json, redirect } from "react-router-dom";
import { dispatch } from "./slices/store";
import { fetchSearchResults } from './slices/search';

export async function actionAuth({ params, request }) {
    const data = await request.formData();
    console.log('Form data', data);
    if (!data.get('password')) {
        return json({ error: 'Password is required' }, { status: 400 });
    }
    console.log('Username', data.get('username'));
    const response = await fetch(`${config.api_base_url}/api/users/${data.get('username')}`);
    console.log('Response', response);
    const user = await response.json();
    console.log('User', user);
    console.log('Dispatch', dispatch(login(user[0])));
    return redirect('/profile');
}

export async function actionComment({ params, request }) {
    const data = await request.formData();
    console.log('Comment form data', data);
    const content = data.get('content');
    if (!content) {
        return {}; // ignore
    }
    const response = await fetch(`${config.api_base_url}/api/comment/add`, /* options */);
    console.log('Comment response', response);
    return response;
}

export async function actionSearch({ params, request }) {
    dispatch(fetchSearchResults());
    return {};
}
