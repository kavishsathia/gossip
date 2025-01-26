import { Avatar } from "@mui/material";
import UserInfo from "../commons/UserInfo";
import { ThreadComment } from "../../../../services/comments/types";

function ModerationFlag({ comment }: { comment: ThreadComment }) {
  return (
    <div className="flex flex-row space-x-2 items-center">
      <Avatar
        src={comment.ProfileImage}
        sx={{ width: "25px", height: "25px" }}
      ></Avatar>
      <UserInfo
        userId={comment.UserID ?? 0}
        username={comment.Username ?? ""}
      />
    </div>
  );
}

export default ModerationFlag;
