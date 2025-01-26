import { Thread } from "../../../../services/threads/types";
import { Avatar } from "@mui/material";
import UserInfo from "../commons/UserInfo";
import { dateTranslate } from "../commons/DatetimeHelper";

function ModerationFlag({ thread }: { thread: Thread }) {
  return (
    <div className="p-4 pt-0 pb-6">
      <p className="text-xs mb-2">WRITTEN BY</p>
      <div className="flex flex-row items-center gap-3">
        <Avatar
          src={thread?.ProfileImage || ""}
          sx={{ width: "30px", height: "30px" }}
        ></Avatar>
        <div className="flex flex-col">
          <UserInfo
            userId={thread?.UserID ?? 0}
            username={thread?.Username ?? ""}
          />
          <p className="text-sm h-fit">
            on {dateTranslate(thread?.CreatedAt ?? "")}
          </p>
        </div>
      </div>
    </div>
  );
}

export default ModerationFlag;
