import { login } from './slices/user';
import { json, redirect } from 'react-router-dom';
import { dispatch } from './slices/store';
import { fetchSearchResults } from './slices/search';
import { commentApi, userApi } from './api';


async function performLogin(data) {
    return userApi.userLoginPost({ login: data.get('username'), password: data.get('password') });
}

async function performRegister(data) {
    if (data.get('password') !== data.get('repeated')) {
        return json({ error: 'Passwords do not match' }, { status: 400 });
    }
    if (!data.get('email')) {
        return json({ error: 'Email is required' }, { status: 400 });
    }
    const fullname = data.get('fullname') || data.get('username');
    return userApi.userRegisterPost({ login: data.get('username'), password: data.get('password'), fullname: fullname, email: data.get('email') });
}

export async function actionAuth({ request }) {
    const data = await request.formData();
    if (!data.get('username')) {
        return json({ error: 'Username is required' }, { status: 400 });
    }
    if (!data.get('password')) {
        return json({ error: 'Password is required' }, { status: 400 });
    }
    const action = data.get('action');
    let response;
    if (action === 'login') {
        response = await performLogin(data);
    } else if (action === 'register') {
        response = await performRegister(data);
    } else {
        return json({ error: 'Invalid action' }, { status: 400 });
    }
    const user = response.data.data;
    dispatch(login(user));
    return redirect('/profile');
}

export async function actionComment({ request }) {
    const data = await request.formData();
    const content = data.get('content');
    const media_id = parseInt(data.get('media_id'));
    if (!content) {
        return {}; // ignore
    }
    await commentApi.commentAddPost({ media_id: media_id, content: content });
    return {};
}

export async function actionSearch() {
    dispatch(fetchSearchResults());
    return {};
}
