import { createSlice } from '@reduxjs/toolkit';

const emptyUser = {
    login: null,
    fullname: null,
    avatar: null,
    media: {},
};

const userSlice = createSlice({
    name: 'user',
    initialState: {
        is_logged_in: false,
        current_user: emptyUser,
    },
    reducers: {
        login: (state, action) => {
            state.is_logged_in = true;
            state.current_user = action.payload;
        },
        logout: (state) => {
            state.is_logged_in = false;
            state.current_user = emptyUser;
        },
        setMedia: (state, action) => {
            const media = action.payload;
            if (!state.is_logged_in) {
                return;
            }
            state.current_user.media[media.id] = media;
        },
        deleteMedia: (state, action) => {
            const { media_id } = action.payload;
            if (!state.is_logged_in) {
                return;
            }
            if (media_id in state.current_user.media) {
                delete state.current_user.media[media_id];
            }
        },
    },
});

export const setMedia = (media) => async (dispatch, getState) => {
    const user = getState().user.current_user;
    if (!user) {
        throw new Error('User not logged in');
    }

    dispatch(userSlice.actions.setMedia(media));
    return {};
};

export const deleteMedia = (media_id) => async (dispatch, getState) => {
    const user = getState().user.current_user;
    if (!user) {
        throw new Error('User not logged in');
    }

    dispatch(userSlice.actions.deleteMedia({ media_id }));
    return {};
};

export const selectIsLoggedIn = (state) => state.user.is_logged_in;
export const selectCurrentUser = (state) => state.user.current_user;
export const selectStatus = (media_id) => (state) => state.user.current_user.media[media_id]?.status;
export const selectRating = (media_id) => (state) => state.user.current_user.media[media_id]?.rating;
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
