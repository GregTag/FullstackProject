import { Tabs, TabList, TabPanel, Tab } from "@mui/joy"
import mediaAttributes from "../config/media_attributes.json";
import TableWithSort from "./TableWithSort";
import { useNavigate } from "react-router-dom";
import { mapStatus } from "../logic/slices/user";

const categories = mediaAttributes.category.values;

const head = [
    {
        id: 'title',
        label: 'Title',
    },
    {
        id: 'status',
        label: 'Status',
    },
    {
        id: 'rating',
        label: 'Rating',
    },
]

const options = [5, 10, 25, 50, 100];

function MediaTable({ category, profile }) {
    const rows = Object.entries(profile.media).filter(([_, val]) => val.category === category).map(([id, { status, ...other }]) => ({ id, status: mapStatus(status), ...other }));
    console.log(rows);
    const navigate = useNavigate();
    const handler = (event, row) => navigate(`/media/${row.id}`);
    return (
        <TableWithSort head={head} rowsPerPageOptions={options} rows={rows} handleClick={handler} />
    )
}

function UserLists({ profile }) {
    return (
        <Tabs color="primary" variant="outlined" defaultValue={categories[0]} sx={{
            borderRadius: 'lg',
            boxShadow: 'sm',
            overflow: 'auto'
        }}>
            <TabList tabFlex={1}>
                {categories.map((category) => (
                    <Tab key={category} value={category} variant="solid" color="primary" sx={{ flexGrow: 1 }}>{category}</Tab>
                ))}
            </TabList>

            {categories.map((category) => (
                <TabPanel key={category} value={category} sx={{ p: 0 }}>
                    <MediaTable category={category} profile={profile} />
                </TabPanel>
            ))}
        </Tabs>
    )
}

export default UserLists
