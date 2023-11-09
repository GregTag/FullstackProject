import React, { useState } from 'react';
import { Typography, Card, Stack, Link } from '@mui/joy';
import { Star, StarBorder } from '@mui/icons-material';
import { useDispatch, useSelector } from 'react-redux';
import { selectRating, setMedia } from '../logic/slices/user';

function StarWithAction({ index, setSelect, select, rating, onClick }) {
    return (<Link underline="none" color={index <= select ? 'warning' : 'neutral'} component="button" onClick={onClick} onMouseEnter={() => setSelect(index)}>
        {index <= rating ? <Star /> : <StarBorder />}
    </Link>);
}

function UserRate({ rating, media_id }) {
    const [select, setSelect] = useState(rating);
    const dispatch = useDispatch();
    const clickHandler = (index) => () => dispatch(setMedia({ id: media_id, rating: index }));
    return (<Stack direction="row">
        {Array(10).fill().map((_, index) => <StarWithAction key={index} index={index + 1} select={select} setSelect={setSelect} rating={rating} onClick={clickHandler(index + 1)} />)}
    </Stack>);
}

const colors = ['danger', 'warning', 'success'];
function ratingToColor(rating) {
    if (rating < 5) return colors[0];
    if (rating < 8) return colors[1];
    return colors[2];
}

function Rating({ rating, media_id }) {
    const user_rating = useSelector(selectRating(media_id));
    return (<Card>
        <Stack direction="row" spacing={2}>
            <Typography color={ratingToColor(rating)}>Total rating: {rating}</Typography>
            {user_rating != null && <Typography color={ratingToColor(user_rating)}>Your rating: {user_rating}</Typography>}
        </Stack>
        {user_rating != null && <UserRate rating={user_rating} media_id={media_id} />}
    </Card>);
}

export default Rating;
