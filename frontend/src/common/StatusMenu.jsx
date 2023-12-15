import React from 'react';
import { Dropdown, MenuButton, Menu, MenuItem } from '@mui/joy';
import { useDispatch, useSelector } from 'react-redux';
import { selectIsLoggedIn, selectTrack, mapStatus, unsetTrack, setTrack } from '../logic/slices/user';
import mediaAttributes from '../config/media_attributes.json';

export default function StatusMenu({ media_id }) {
    const dispatch = useDispatch();
    const is_logged = useSelector(selectIsLoggedIn);
    const track = useSelector(selectTrack(media_id));
    const track_status = track?.track_status;
    const clickHandler = (value) => () => {
        dispatch(setTrack({ media_id: media_id, track_status: value, rating: track?.rating }));
    };
    return (
        <Dropdown>
            {is_logged && <MenuButton>{mapStatus(track_status)}</MenuButton>}
            <Menu>
                {mediaAttributes.status.values.map((status) => (<MenuItem key={status} onClick={clickHandler(status)} selected={status === track_status}>{mapStatus(status)}</MenuItem>))}
                <MenuItem color="danger" selected={!track_status} onClick={() => dispatch(unsetTrack(media_id))}>Not tracked</MenuItem>
            </Menu>
        </Dropdown>
    );
}
