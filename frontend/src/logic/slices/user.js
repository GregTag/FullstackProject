import { createSlice } from "@reduxjs/toolkit";

const userSlice = createSlice({
    name: "user",
    initialState: {
        current_user: null,
    },
    reducers: {
        login: (state, action) => {
            state.current_user = {
                username: action.payload.username,
                avatar_url: action.payload.avatar_url,
            };
        },
        logout: state => {
            state.current_user = null;
        }
    }
});

export const selectIsLoggedIn = state => state.user.current_user !== null;
export const selectCurrentUser = state => state.user.current_user;
export const { login, logout } = userSlice.actions;
export default userSlice.reducer;
