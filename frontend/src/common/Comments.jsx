import { Await, Form } from "react-router-dom";
import './Comments.css'
import { loadComments } from "../logic/loaders";
import React from "react";
import { useSelector } from "react-redux";
import { selectIsLoggedIn } from "../logic/user";

function CommentForm() {
    return (
        <Form method='post' navigate={false}>
            <input type="text" name="content" />
            <input type="submit" value="Send" />
        </Form>
    )
}

function Comment({ comment }) {
    return (<div key={comment.id} className="comment">
        <p>{comment.send_at}</p>
        <p>{comment.content}</p>
    </div>)
}

function Comments({ media }) {
    const is_logged = useSelector(selectIsLoggedIn);
    return (<div className="round-panel">
        <React.Suspense fallback={<div>Loading...</div>}>
            <Await
                resolve={loadComments(media)}
                errorElement={<div>Failed to load comments</div>}
            >
                {comments => comments.map(comment => <Comment comment={comment} />)}
            </Await>
        </React.Suspense>
        {is_logged && <CommentForm />}
    </div>)
}

export default Comments;
