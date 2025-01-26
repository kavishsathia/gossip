import { IconButton } from "@mui/material";
import { ThreadComment } from "../../../../services/comments/types";
import { Heart, MessageCircle } from "lucide-react";
import {
  likeThreadComment,
  unlikeThreadComment,
} from "../../../../services/comments";
import { useState } from "react";

function ModerationFlag({ comment }: { comment: ThreadComment }) {
  const [liked, setLiked] = useState(comment.Liked);
  const [likes, setLikes] = useState(comment.Likes);

  return (
    <div className="flex items-center space-x-4">
      <div className="flex items-center space-x-1">
        <IconButton
          onClick={async () => {
            if (liked) {
              setLiked(!liked);
              setLikes(likes - 1);
              await unlikeThreadComment(comment.ID);
            } else {
              setLiked(!liked);
              setLikes(likes + 1);
              await likeThreadComment(comment.ID);
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
        <span className="text-sm text-gray-600">{comment.Comments}</span>
      </div>
    </div>
  );
}

export default ModerationFlag;
