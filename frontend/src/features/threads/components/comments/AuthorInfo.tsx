import UserInfo from "../commons/UserInfo";
import { ThreadComment } from "../../../../services/comments/types";

function ModerationFlag({ comment }: { comment: ThreadComment }) {
  return (
    <div className="flex flex-row space-x-2 items-center">
      <UserInfo
        width="25px"
        picture={true}
        userId={comment.UserID ?? 0}
        username={comment.Username ?? ""}
      />
      <UserInfo
        width=""
        picture={false}
        userId={comment.UserID ?? 0}
        username={comment.Username ?? ""}
      />
    </div>
  );
}

export default ModerationFlag;
