import { enqueueSnackbar } from "notistack";
import { reportThread } from "../../../../services/threads";
import { Thread } from "../../../../services/threads/types";

function ModerationFlag({ thread }: { thread: Thread }) {
  return (
    thread.ModerationFlag && (
      <div className="w-full p-4">
        <div className="flex flex-row flex-wrap gap-y-5 justify-between items-center bg-red-600/20 w-full rounded-md p-5 border-2 border-red-700">
          <div>
            <p className="font-bold">Content Moderation Warning</p>
            <p>
              This thread was flagged for {thread.ModerationFlag}. Please report
              it if the flagging was accurate.
            </p>
          </div>
          <div className="flex-shrink-0">
            <button
              onClick={() => {
                reportThread(thread.ID);
                enqueueSnackbar("The thread was reported. Thank you!");
              }}
              className="bg-red-600/20 border-red-700"
            >
              Report
            </button>
          </div>
        </div>
      </div>
    )
  );
}

export default ModerationFlag;
