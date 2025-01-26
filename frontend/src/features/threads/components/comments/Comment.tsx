import "@mdxeditor/editor/style.css";
import { useState } from "react";
import { ThreadComment } from "../../../../services/comments/types";
import { TextField } from "@mui/material";
import Comments from ".";
import AuthorInfo from "./AuthorInfo";
import EditorToolbox from "../comments/EditorToolbox";
import InteractionBar from "./InteractionBar";

function App({
  comment,
  index,
  depth,
}: {
  comment: ThreadComment;
  index: number;
  depth: number;
}) {
  const [repliesOpen, setRepliesOpen] = useState<boolean>(false);
  const [commentBuffer, setCommentBuffer] = useState<string>(comment.Comment);
  const [editing, setEditing] = useState<boolean>(false);

  return (
    <div
      className={`mb-4 rounded-lg p-5 ${
        index % 2 === 0 ? "bg-teal-600/5" : "bg-neutral-50/0"
      }`}
    >
      <div className="flex flex-row justify-between">
        <AuthorInfo comment={comment} />
        <EditorToolbox
          comment={comment}
          setEditing={setEditing}
          editing={editing}
          commentBuffer={commentBuffer}
          setCommentBuffer={setCommentBuffer}
        />
      </div>

      {editing ? (
        <TextField
          multiline
          value={commentBuffer}
          onChange={(e) => setCommentBuffer(e.currentTarget.value)}
          label="Comment"
          variant="outlined"
          className="w-full"
          size="small"
          sx={{
            marginTop: "10px",
            ".MuiInputBase-input": {
              fontFamily: "Inter",
            },
          }}
        />
      ) : (
        <div className="mt-2">{commentBuffer}</div>
      )}

      <div className="flex items-center justify-between mt-2">
        {depth < 5 ? (
          <span
            onClick={() => setRepliesOpen(!repliesOpen)}
            className="text-sm hover:underline cursor-pointer"
          >
            {repliesOpen ? "Hide" : "Show"} replies
          </span>
        ) : (
          <div />
        )}

        <InteractionBar comment={comment} />
      </div>

      {repliesOpen && depth < 5 ? (
        <div className="grid grid-cols-12 mt-5">
          <div className="flex justify-center mb-5">
            <div
              onClick={() => setRepliesOpen(false)}
              className="w-[1px] bg-gray-300 hover:w-[2px] hover:bg-teal-700 cursor-pointer"
            ></div>
          </div>
          <div className="col-span-11">
            <Comments
              id={comment.ID}
              isFromPost={false}
              depth={depth}
            ></Comments>
          </div>
        </div>
      ) : (
        <div></div>
      )}
    </div>
  );
}

export default App;
