import "@mdxeditor/editor/style.css";
import { useState } from "react";
import {
  createThreadComment,
  createThreadCommentComment,
  listThreadCommentComments,
  listThreadComments,
} from "../../../../services/comments";
import React from "react";
import { ThreadComment } from "../../../../services/comments/types";
import Comment from "./Comment";
import { TextField } from "@mui/material";

function App({ id, isFromPost }: { id: number | null; isFromPost: boolean }) {
  const [comments, setComments] = useState<ThreadComment[]>();
  const [comment, setComment] = useState<string>("");

  React.useEffect(() => {
    const fetchThreads = async () => {
      let data;
      if (isFromPost) {
        data = await listThreadComments(id ?? 0);
      } else {
        data = await listThreadCommentComments(id ?? 0);
      }
      setComments(data);
    };

    fetchThreads();
  }, [id, isFromPost]);

  return (
    <div className="w-full">
      <div>
        <TextField
          size="small"
          onKeyDown={async (e) => {
            if (e.key === "Enter") {
              if (isFromPost) {
                await createThreadComment(id ?? 0, comment);
                const data = await listThreadComments(id ?? 0);
                setComments(data);
                setComment("");
              } else {
                await createThreadCommentComment(id ?? 0, comment);
                const data = await listThreadCommentComments(id ?? 0);
                setComments(data);
                setComment("");
              }
            }
          }}
          value={comment}
          onChange={(e) => setComment(e.target.value)}
          variant="filled"
          fullWidth
          label="Leave a comment"
        ></TextField>
      </div>
      <div className="mt-6">
        {comments?.map((comment, index) => (
          <Comment key={index} comment={comment} index={index}></Comment>
        ))}
      </div>
    </div>
  );
}

export default App;
