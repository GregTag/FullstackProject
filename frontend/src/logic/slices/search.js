import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import config from '../../config/config.json';
import mediaAttributes from '../../config/media_attributes.json';

function getDefaultValue(filter) {
    return [filter.min, filter.max];
}

const searchSlice = createSlice({
    name: 'search',
    initialState: {
        searchTerm: '',
        filters: {
            category: mediaAttributes.category.values[0],
            genres: [],
            rating: getDefaultValue(mediaAttributes.rating),
            duration: getDefaultValue(mediaAttributes.duration),
            year: getDefaultValue(mediaAttributes.year),
        },
        results: [],
        state: 'idle',
        error: null,
    },
    reducers: {
        setSearchTerm: (state, action) => {
            state.searchTerm = action.payload;
        },

        setFilter: (state, action) => {
            const { filter, value } = action.payload;
            state.filters[filter] = value;
        },

        setResults: (state, action) => {
            state.results = action.payload;
        },
    },
    extraReducers: (builder) => {
        builder
            .addCase(fetchSearchResults.pending, (state) => {
                state.status = 'loading';
            })
            .addCase(fetchSearchResults.fulfilled, (state, action) => {
                state.status = 'succeeded';
                state.results = action.payload;
            })
            .addCase(fetchSearchResults.rejected, (state, action) => {
                state.status = 'failed';
                state.error = action.error.message;
            });
    },
});

export const fetchSearchResults = createAsyncThunk('search/fetchSearchResults', async () => {
    const response = await fetch(`${config.api_base_url}/api/search` /* options */);
    const results = await response.json();
    return results[0].hits;
});

export const setFilterAndFetch = (params) => async (dispatch) => {
    await dispatch(setFilter(params));
    await dispatch(fetchSearchResults());
};

export const setSearchTermAndFetch = (params) => async (dispatch) => {
    await dispatch(setSearchTerm(params));
    await dispatch(fetchSearchResults());
};

export const selectResults = (state) => state.search.results;
export const selectFilters = (state) => state.search.filters;
export const selectStatus = (state) => state.search.status;
export const selectError = (state) => state.search.error;
export const { setSearchTerm, setFilter, setResults } = searchSlice.actions;
export default searchSlice.reducer;
