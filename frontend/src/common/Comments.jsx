import React from 'react';
import { Form, Link as RouterLink } from 'react-router-dom';
import { useSelector } from 'react-redux';
import { selectIsLoggedIn } from '../logic/slices/user';
import { Typography, Button, Stack, Textarea, Card, CardOverflow, Avatar } from '@mui/joy';

function CommentForm({ media_id }) {
    return (
        <Form method='post'>
            <Stack direction="row" spacing={2}>
                <input type="hidden" name="media_id" value={media_id} />
                <Textarea color="primary" variant="soft" name="content" placeholder="Your comment.." maxRows={2} sx={{ flexGrow: 1 }} />
                <Button size="lg" variant="soft" type="submit">Send</Button>
            </Stack>
        </Form>
    );
}

function Comment({ comment }) {
    return (
        <Card orientation="horizontal">
            <Avatar component={RouterLink} to={`/profile/${comment.sender.login}`} src={comment.sender.avatar} />
            <Stack direction="column">
                <Typography level="title-lg">{comment.sender.fullname} <Typography level="body-xs">{new Date(comment.created_at).toLocaleString()}</Typography></Typography>
                <Typography level="body-md">{comment.content}</Typography>
            </Stack>
        </Card>
    );
}

function Comments({ comments, media_id }) {
    const is_logged = useSelector(selectIsLoggedIn);
    return (<Card color="primary">
        <CardOverflow variant="solid" color="primary">
            <Typography level="title-md">Comments</Typography>
        </CardOverflow>
        {comments.length != 0 ? comments.map(comment => <Comment key={comment.id} comment={comment} />) : <Typography level="body-md">No comments yet</Typography>}
        {is_logged && <CardOverflow color="primary" variant="solid"><CommentForm media_id={media_id} /></CardOverflow>}
    </Card>);
}

export default Comments;
