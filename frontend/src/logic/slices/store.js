import { configureStore } from '@reduxjs/toolkit';
import { persistStore, persistReducer } from 'redux-persist'
import storage from 'redux-persist/lib/storage'
import userReducer from './user';
import searchReducer from './search';

const config = {
    key: 'root',
    storage,
}

const store = configureStore({
    reducer: {
        user: persistReducer(config, userReducer),
        search: persistReducer(config, searchReducer),
    }
});

export default store;
export const persistor = persistStore(store);

export const applySelector = (selector) => selector(store.getState())
export const dispatch = (action) => store.dispatch(action)
