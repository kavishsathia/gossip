import { enqueueSnackbar } from "notistack";
import { likeThread, unlikeThread } from "../../../../services/threads";
import { Thread } from "../../../../services/threads/types";
import { Heart, MessageCircle, Share2 } from "lucide-react";

function ModerationFlag({
  thread,
  setThread,
}: {
  thread: Thread;
  setThread: React.Dispatch<React.SetStateAction<Thread | undefined>>;
}) {
  return (
    <div className="grid grid-cols-3 border-neutral-200 border-b-2 border-t-2 py-3">
      <div
        onClick={async () => {
          if (thread?.Liked) {
            setThread({
              ...thread,
              Liked: false,
              Likes: thread.Likes - 1,
            });
            await unlikeThread(thread.ID);
          } else {
            setThread({
              ...thread!,
              Liked: true,
              Likes: thread!.Likes + 1,
            });
            await likeThread(thread!.ID);
          }
        }}
        className="space-x-2 w-full flex justify-center cursor-pointer"
      >
        <Heart
          color={thread?.Liked ? "red" : "black"}
          className={`inline ${
            thread?.Liked ? "fill-red-500" : "hover:fill-red-300"
          } hover:scale-110 `}
        />
        <span>
          <p className="hidden md:inline-block">Like •</p> {thread?.Likes}
        </span>
      </div>
      <div className="space-x-2 w-full flex justify-center">
        <MessageCircle className="inline" />
        <span>
          <p className="hidden md:inline-block">Comment •</p> {thread?.Comments}
        </span>
      </div>
      <div
        onClick={() => {
          navigator.clipboard.writeText(window.location.href);
          enqueueSnackbar("The link has been copied to your clipboard!");
        }}
        className="space-x-2 w-full flex justify-center"
      >
        <Share2 className="inline" />
        <span>
          <p className="hidden md:inline-block">Share •</p> {thread?.Shares}
        </span>
      </div>
    </div>
  );
}

export default ModerationFlag;
