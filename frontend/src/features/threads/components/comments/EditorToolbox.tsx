import { Button } from "@mui/material";
import { ThreadComment } from "../../../../services/comments/types";
import { PencilIcon, Trash } from "lucide-react";
import {
  deleteThreadComment,
  editThreadComment,
} from "../../../../services/comments";
import { useContext } from "react";
import { User } from "../../context";

function ModerationFlag({
  comment,
  editing,
  setEditing,
  commentBuffer,
  setCommentBuffer,
}: {
  comment: ThreadComment;
  editing: boolean;
  setEditing: React.Dispatch<React.SetStateAction<boolean>>;
  commentBuffer: string;
  setCommentBuffer: React.Dispatch<React.SetStateAction<string>>;
}) {
  const user = useContext(User);

  return comment.UserID === user?.ID && !comment.Deleted && !editing ? (
    <div className="flex flex-row space-x-2 items-center">
      <PencilIcon
        onClick={() => setEditing(true)}
        className="size-5 hover:text-teal-500 cursor-pointer"
      />
      <Trash
        onClick={() => {
          if (confirm("Are you sure?")) {
            deleteThreadComment(comment.ID);
            setCommentBuffer("[deleted]");
          }
        }}
        className="size-5 hover:text-red-500 cursor-pointer"
      />
    </div>
  ) : editing ? (
    <Button
      variant="outlined"
      size="small"
      onClick={() => {
        editThreadComment(comment.ID, commentBuffer);
        setEditing(false);
      }}
      disableElevation
    >
      Edit
    </Button>
  ) : (
    <div />
  );
}

export default ModerationFlag;
