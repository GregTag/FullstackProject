import React from 'react';
import { Await, Form, Link as RouterLink } from 'react-router-dom';
import { loadComments } from '../logic/loaders';
import { useSelector } from 'react-redux';
import { selectIsLoggedIn } from '../logic/slices/user';
import { Typography, Button, Stack, Textarea, Card, CardOverflow, Avatar } from '@mui/joy';

function CommentForm() {
    return (
        <Form method='post' navigate="false">
            <Stack direction="row" spacing={2}>
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
                <Typography level="title-lg">{comment.sender.fullname} <Typography level="body-xs">{new Date(comment.send_at).toLocaleString()}</Typography></Typography>
                <Typography level="body-md">{comment.content}</Typography>
            </Stack>
        </Card>
    );
}

function Comments({ media }) {
    const is_logged = useSelector(selectIsLoggedIn);
    return (<Card color="primary">
        <CardOverflow variant="solid" color="primary">
            <Typography level="title-md">Comments</Typography>
        </CardOverflow>
        <React.Suspense fallback={<div>Loading...</div>}>
            <Await
                resolve={loadComments(media)}
                errorElement={<div>Failed to load comments</div>}
            >
                {comments => comments.map(comment => <Comment key={comment.id} comment={comment} />)}
            </Await>
        </React.Suspense>
        {is_logged && <CardOverflow color="primary" variant="solid"><CommentForm /></CardOverflow>}
    </Card>);
}

export default Comments;
