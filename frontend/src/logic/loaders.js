import { redirect } from 'react-router-dom';
import { applySelector } from './slices/store';
import { selectCurrentUser } from './slices/user';
import { mediaApi, userApi } from './api';

export async function loadMedia({ params }) {
    const response = await mediaApi.mediaLoadIdGet(params.id);
    return response.data.data;
}

export async function loadProfile({ params }) {
    const user = applySelector(selectCurrentUser);
    if (!params.name) {
        if (user) {
            return redirect(`/profile/${user.login}`);
        } else {
            return redirect('/auth');
        }
    }
    const response = await userApi.userLoadLoginGet(params.name);
    return response.data.data;
}
