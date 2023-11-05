import { createSlice } from "@reduxjs/toolkit";

const searchSlice = createSlice({
    name: "search",
    initialState: {
        searchTerm: "",
        filters: {},
        results: []
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
        }
    }
});

export const selectResults = state => state.search.results;
export const { setSearchTerm, setFilter, setResults } = searchSlice.actions;
export default searchSlice.reducer;
