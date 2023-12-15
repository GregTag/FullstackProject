import { createSlice, createAsyncThunk } from '@reduxjs/toolkit';
import mediaAttributes from '../../config/media_attributes.json';
import { mediaApi } from '../api';

function getDefaultValue(filter) {
    return [filter.min, filter.max];
}

const searchSlice = createSlice({
    name: 'search',
    initialState: {
        filters: {
            searchTerm: '',
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
            state.filters.searchTerm = action.payload;
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

function getFilterParams(filters) {
    const params = {};
    params.title = filters.searchTerm;
    params.category = filters.category;
    params.genres = filters.genres;
    if (filters.rating[0] !== mediaAttributes.rating.min) { params.rating_from = filters.rating[0]; }
    if (filters.rating[1] !== mediaAttributes.rating.max) { params.rating_to = filters.rating[1]; }
    if (filters.duration[0] !== mediaAttributes.duration.min) { params.duration_from = filters.duration[0]; }
    if (filters.duration[1] !== mediaAttributes.duration.max) { params.duration_to = filters.duration[1]; }
    if (filters.year[0] !== mediaAttributes.year.min) { params.year_from = filters.year[0]; }
    if (filters.year[1] !== mediaAttributes.year.max) { params.year_to = filters.year[1]; }
    return params;
}

export const fetchSearchResults = createAsyncThunk('search/fetchSearchResults', async (_, thunkAPI) => {
    const filters = selectFilters(thunkAPI.getState());
    const params = getFilterParams(filters);
    const response = await mediaApi.searchPost(params);
    return response.data.data || [];
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
