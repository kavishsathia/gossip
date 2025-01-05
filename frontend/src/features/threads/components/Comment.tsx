import "@mdxeditor/editor/style.css";
import { useContext, useState } from "react";
import { ThreadComment } from "../../../services/threads/types";
import { Avatar, Button, IconButton, TextField } from "@mui/material";
import { Heart, MessageCircle, PencilIcon, Trash } from "lucide-react";
import Comments from "../components/Comments";
import {
  deleteThreadComment,
  editThreadComment,
  likeThreadComment,
  unlikeThreadComment,
} from "../../../services/threads";
import UserInfo from "./UserInfo";
import { User } from "../context";

function App({ item, index }: { item: ThreadComment; index: number }) {
  const [repliesOpen, setRepliesOpen] = useState<boolean>(false);
  const [liked, setLiked] = useState<boolean>(item.Liked ?? false);
  const [likes, setLikes] = useState<number>(item.Likes ?? 0);
  const [comment, setComment] = useState<string>(item.Comment);
  const [editing, setEditing] = useState<boolean>(false);
  const user = useContext(User);

  console.log(user);

  return (
    <div
      className={`mb-4 rounded-lg p-5 ${
        index % 2 === 0 ? "bg-teal-600/5" : "bg-neutral-50/0"
      }`}
    >
      <div className="flex flex-row justify-between">
        <div className="flex flex-row space-x-2 items-center">
          <Avatar
            src={item.ProfileImage}
            sx={{ width: "25px", height: "25px" }}
          ></Avatar>
          <UserInfo userId={item.UserID ?? 0} username={item.Username ?? ""} />
        </div>
        {item.UserID === user?.ID && !item.Deleted && !editing ? (
          <div className="flex flex-row space-x-2 items-center">
            <PencilIcon
              onClick={() => setEditing(true)}
              className="size-5 hover:text-teal-500 cursor-pointer"
            />
            <Trash
              onClick={() => {
                if (confirm("Are you sure?")) {
                  deleteThreadComment(item.ID);
                  setComment("[deleted]");
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
              editThreadComment(item.ID, comment);
              setEditing(false);
            }}
            disableElevation
          >
            Edit
          </Button>
        ) : (
          <div />
        )}
      </div>
      {editing ? (
        <TextField
          multiline
          value={comment}
          onChange={(e) => setComment(e.currentTarget.value)}
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
        <div className="mt-2">{comment}</div>
      )}
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
              className="w-[1px] bg-gray-300 hover:w-[2px] hover:bg-teal-700 cursor-pointer"
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
