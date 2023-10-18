import { createSlice } from "@reduxjs/toolkit";

const userSlice = createSlice({
    name: "user",
    initialState: {
        is_logged_in: false
    },
    reducers: {
        login: state => {
            state.is_logged_in = true;
        },
        logout: state => {
            state.is_logged_in = false;
        }
    }
});

export const selectIsLoggedIn = state => state.user.is_logged_in;
export const { login, logout } = userSlice.actions;
export default userSlice.reducer;
