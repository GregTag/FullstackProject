import { Dropdown, MenuButton, Menu, MenuItem } from "@mui/joy";
import { useDispatch, useSelector } from "react-redux";
import { selectIsLoggedIn, selectStatus, setMedia, deleteMedia, mapStatus } from "../logic/slices/user";
import mediaAttributes from "../config/media_attributes.json";

export default function StatusMenu({ media_id }) {
    const dispatch = useDispatch();
    const is_logged = useSelector(selectIsLoggedIn);
    const current_status = useSelector(selectStatus(media_id));
    const clickHandler = (value) => () => {
        dispatch(setMedia({ id: media_id, status: value }));
    }
    return (
        <Dropdown>
            {is_logged && <MenuButton>{mapStatus(current_status)}</MenuButton>}
            <Menu>
                {mediaAttributes.status.values.map((status) => (<MenuItem key={status} onClick={clickHandler(status)} selected={status === current_status}>{mapStatus(status)}</MenuItem>))}
                <MenuItem color="danger" selected={!current_status} onClick={() => dispatch(deleteMedia(media_id))}>Not tracked</MenuItem>
            </Menu>
        </Dropdown>
    )
}
