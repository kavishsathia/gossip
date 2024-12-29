import "@mdxeditor/editor/style.css";
import { useState } from "react";
import { ThreadComment } from "../../../services/threads/types";
import { Avatar, IconButton } from "@mui/material";
import { Heart, MessageCircle } from "lucide-react";
import Comments from "../components/Comments";
import {
  likeThreadComment,
  unlikeThreadComment,
} from "../../../services/threads";
import UserInfo from "./UserInfo";

function App({ item, index }: { item: ThreadComment; index: number }) {
  const [repliesOpen, setRepliesOpen] = useState<boolean>(false);
  const [liked, setLiked] = useState<boolean>(item.Liked ?? false);
  const [likes, setLikes] = useState<number>(item.Likes ?? 0);

  return (
    <div
      className={`mb-4 rounded-lg p-5 ${
        index % 2 === 0 ? "bg-gray-100" : "bg-neutral-50"
      }`}
    >
      <div className="flex flex-row space-x-2 items-center">
        <Avatar
          src={item.ProfileImage}
          sx={{ width: "25px", height: "25px" }}
        ></Avatar>
        <UserInfo userId={item.UserID ?? 0} username={item.Username ?? ""} />
      </div>
      <div className="mt-2">{item.Comment}</div>
      <div className="flex items-center justify-between mt-2">
        <span
          onClick={() => setRepliesOpen(!repliesOpen)}
          className="text-sm hover:underline cursor-pointer"
        >
          {repliesOpen ? "Hide" : "Show"} replies
        </span>

        <div className="flex items-center space-x-4">
          <div className="flex items-center space-x-1">
            <IconButton
              onClick={async () => {
                if (liked) {
                  setLiked(!liked);
                  setLikes(likes - 1);
                  await unlikeThreadComment(item.ID);
                } else {
                  setLiked(!liked);
                  setLikes(likes + 1);
                  await likeThreadComment(item.ID);
                }
              }}
            >
              <Heart
                color={liked ? "red" : "black"}
                className={`w-5 h-5 ${
                  liked ? "fill-red-500" : "hover:fill-red-300"
                } hover:scale-110`}
              />
            </IconButton>
            <span className="text-sm text-gray-600">{likes}</span>
          </div>

          <div className="flex items-center space-x-1">
            <IconButton>
              <MessageCircle className="w-5 h-5 text-gray-600" />
            </IconButton>
            <span className="text-sm text-gray-600">{item.Comments}</span>
          </div>
        </div>
      </div>
      {repliesOpen ? (
        <div className="grid grid-cols-12 mt-5">
          <div className="flex justify-center mb-5">
            <div
              onClick={() => setRepliesOpen(false)}
              className="w-[1px] bg-gray-300
             hover:bg-gray-600 hover:w-[2px] cursor-pointer"
            ></div>
          </div>
          <div className="col-span-11">
            <Comments id={item.ID} isFromPost={false}></Comments>
          </div>
        </div>
      ) : (
        <div></div>
      )}
    </div>
  );
}

export default App;
