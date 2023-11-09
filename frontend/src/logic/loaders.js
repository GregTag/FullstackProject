import config from '../config/config.json';
import { redirect } from 'react-router-dom';
import { applySelector } from './slices/store';
import { selectCurrentUser } from './slices/user';

export async function loadMedia({ params }) {
    return fetch(`${config.api_base_url}/api/media/${params.name}`, {
        method: 'GET',
    }).then((res) => res.json());
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
    return fetch(`${config.api_base_url}/api/users/${params.name}`, {
        method: 'GET',
    }).then((res) => res.json()).then((res) => res ? res[0] : null);
}

export async function loadComments(media) {
    return fetch(`${config.api_base_url}/api/comments/${media}`, {
        method: 'GET',
    }).then((res) => res.json());
}
