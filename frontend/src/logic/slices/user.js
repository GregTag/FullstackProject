import { createSlice } from '@reduxjs/toolkit';
import { trackApi, userApi } from '../api';

const emptyUser = {
    login: null,
    fullname: null,
    avatar: null,
};

const userSlice = createSlice({
    name: 'user',
    initialState: {
        is_logged_in: false,
        current_user: emptyUser,
        tracks: {},
    },
    reducers: {
        login: (state, action) => {
            state.is_logged_in = true;
            state.current_user = action.payload.user;

            if (action.payload.tracks) {
                for (const track of action.payload.tracks) {
                    state.tracks[track.media_id] = track;
                }
            }
        },
        logout: (state) => {
            state.is_logged_in = false;
            state.current_user = emptyUser;
            state.tracks = {};
            userApi.userLogoutPost().catch((err) => console.log(err)).finally(() => {
                localStorage.removeItem('auth');
            });
        },
        track: (state, action) => {
            const track = action.payload;
            if (!state.is_logged_in) {
                return;
            }
            state.tracks[track.media_id] = track;
        },
        untrack: (state, action) => {
            const { media_id } = action.payload;
            if (!state.is_logged_in) {
                return;
            }
            delete state.tracks[media_id];
        },
    },
});

export const setTrack = (track) => async (dispatch, getState) => {
    const user = getState().user;
    if (!user) {
        throw new Error('User not logged in');
    }

    let response;
    const trackData = { media_id: track.media_id, track_status: track.track_status, rating: track.rating };
    if (track.media_id in user.tracks) {
        response = await trackApi.trackEditPut(trackData);
    } else {
        response = await trackApi.trackAddPost(trackData);
    }
    track = response.data.data;

    dispatch(userSlice.actions.track(track));
    return {};
};

export const unsetTrack = (media_id) => async (dispatch, getState) => {
    const user = getState().user.current_user;
    if (!user) {
        throw new Error('User not logged in');
    }
    await trackApi.trackDeleteIdDelete(media_id);
    dispatch(userSlice.actions.untrack({ media_id }));
    return {};
};

export const selectIsLoggedIn = (state) => state.user.is_logged_in;
export const selectCurrentUser = (state) => state.user.current_user;
export const selectTrack = (media_id) => (state) => state.user.tracks[media_id];
export const { login, logout } = userSlice.actions;
export default userSlice.reducer;

export function mapStatus(status) {
    switch (status) {
    case 'planned':
        return 'Planned';
    case 'in_progress':
        return 'In Progress';
    case 'completed':
        return 'Completed';
    case 'dropped':
        return 'Dropped';
    default:
        return 'Not tracked';
    }
}
