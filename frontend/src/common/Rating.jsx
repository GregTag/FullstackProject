import React, { useState } from 'react';
import { Typography, Card, Stack, Link } from '@mui/joy';
import { Star, StarBorder } from '@mui/icons-material';
import { useDispatch, useSelector } from 'react-redux';
import { selectTrack, setTrack } from '../logic/slices/user';

function StarWithAction({ index, setSelect, select, rating, onClick }) {
    return (<Link underline="none" color={index <= select ? 'warning' : 'neutral'} component="button" onClick={onClick} onMouseEnter={() => setSelect(index)}>
        {index <= rating ? <Star /> : <StarBorder />}
    </Link>);
}

function UserRate({ track }) {
    const [select, setSelect] = useState(track.rating);
    const dispatch = useDispatch();
    const clickHandler = (index) => () => dispatch(setTrack({ media_id: track.media_id, rating: index, track_status: track?.track_status }));
    return (<Stack direction="row">
        {Array(10).fill().map((_, index) => <StarWithAction key={index} index={index + 1} select={select} setSelect={setSelect} rating={track.rating} onClick={clickHandler(index + 1)} />)}
    </Stack>);
}

const colors = ['danger', 'warning', 'success'];
function ratingToColor(rating) {
    if (rating < 5) return colors[0];
    if (rating < 8) return colors[1];
    return colors[2];
}

function Rating({ media }) {
    const track = useSelector(selectTrack(media.id));
    const user_rating = track?.rating;
    const rating = media.number_of_ratings == 0 ? 0 : media.cumulative_rating / media.number_of_ratings;
    return (<Card>
        <Stack direction="row" spacing={2}>
            <Typography color={ratingToColor(rating)}>Total rating: {rating}</Typography>
            {user_rating != null && <Typography color={ratingToColor(user_rating)}>Your rating: {user_rating}</Typography>}
        </Stack>
        {user_rating != null && <UserRate track={track} />}
    </Card>);
}

export default Rating;
