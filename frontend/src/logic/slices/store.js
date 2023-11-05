import { configureStore } from '@reduxjs/toolkit';
import userReducer from './user';

const store = configureStore({
    reducer: {
        user: userReducer
    }
});

export default store;

export const applySelector = (selector) => selector(store.getState())
export const dispatch = (action) => store.dispatch(action)
