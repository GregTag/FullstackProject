import React, { useRef } from 'react';
import { useDispatch, useSelector } from 'react-redux';
import mediaAttributes from '../config/media_attributes.json';
import { Accordion, AccordionDetails, AccordionGroup, AccordionSummary, Typography, Checkbox, Radio, RadioGroup, List, ListItem, Slider } from '@mui/joy';
import { selectFilters, setFilterAndFetch } from '../logic/slices/search';

function FilterAccordion({ title, children }) {
    return (
        <Accordion defaultExpanded>
            <AccordionSummary>
                <Typography level="h4">{title}</Typography>
            </AccordionSummary>
            <AccordionDetails>
                {children}
            </AccordionDetails>
        </Accordion>
    );
}

function CategoryRadio({ value }) {
    const lower = value.toLowerCase();
    return (
        <ListItem>
            <Radio value={lower} label={value} />
        </ListItem>
    );
}

function GenreCheckbox({ value, filter }) {
    const dispatch = useDispatch();
    const lower = value.toLowerCase();
    const handler = (event) => dispatch(setFilterAndFetch({ filter: 'genres', value: event.target.checked ? filter.concat(lower) : filter.filter((val) => val !== lower) }));

    return (
        <ListItem sx={{ mr: '0px' }}>
            <Checkbox value={value} label={value} disableIcon overlay size="sm" checked={filter.includes(lower)} onChange={handler} />
        </ListItem>
    );
}

function DebounceSlider(props) {
    const { handleDebounce, debounceTimeout, ...rest } = props;
    const timerRef = useRef();
    const handleChange = (event) => {
        if (timerRef.current) {
            clearTimeout(timerRef.current);
        }

        timerRef.current = setTimeout(() => {
            handleDebounce(event.target.value);
        }, debounceTimeout);
    };

    return <Slider {...rest} onChange={handleChange} />;
}

function RangeSlider({ min, max, step, name, filter }) {
    const dispatch = useDispatch();
    const handler = (value) => dispatch(setFilterAndFetch({ filter: name, value }));

    return (
        <DebounceSlider min={min} max={max} step={step} defaultValue={filter} valueLabelDisplay="auto" sx={{ mt: '16px' }} debounceTimeout={1000} handleDebounce={handler} />
    );
}

function FilterComponent() {
    const dispatch = useDispatch();
    const filters = useSelector(selectFilters);
    return (
        <AccordionGroup>
            <FilterAccordion title="Category">
                <RadioGroup value={filters.category} onChange={(event) => dispatch(setFilterAndFetch({ filter: 'category', value: event.target.value }))}>
                    <List orientation="horizontal" wrap>
                        {mediaAttributes.category.values.map((value, index) => (<CategoryRadio key={index} value={value} />))}
                    </List>
                </RadioGroup>
            </FilterAccordion>
            <FilterAccordion title="Genres">
                <List orientation="horizontal" wrap sx={{
                    '--List-gap': '8px',
                    '--ListItem-radius': '20px'
                }}>
                    {mediaAttributes.genres.values.map((value, index) => (<GenreCheckbox key={index} value={value} filter={filters.genres} />))}
                </List>
            </FilterAccordion>
            <FilterAccordion title="Rating">
                <RangeSlider {...mediaAttributes.rating} filter={filters.rating} />
            </FilterAccordion>
            <FilterAccordion title="Duration">
                <RangeSlider {...mediaAttributes.duration} filter={filters.duration} />
            </FilterAccordion>
            <FilterAccordion title="Release year">
                <RangeSlider {...mediaAttributes.year} filter={filters.year} />
            </FilterAccordion>
        </AccordionGroup>
    );
}

export default FilterComponent;
