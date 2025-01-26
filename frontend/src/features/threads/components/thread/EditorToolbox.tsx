import { deleteThread } from "../../../../services/threads";
import { Thread } from "../../../../services/threads/types";
import { Link } from "react-router";
import { PencilIcon, Trash } from "lucide-react";
import { useContext } from "react";
import { User } from "../../context";

function EditorToolbox({
  thread,
  setMarkdown,
}: {
  thread: Thread;
  setMarkdown: React.Dispatch<React.SetStateAction<string>>;
}) {
  const user = useContext(User);

  return thread.UserID === user?.ID && !thread.Deleted ? (
    <div className="flex flex-row space-x-2 items-center m-4 mt-6 pr-5">
      <Link to={`/editor/${thread.ID}`}>
        <PencilIcon className="size-5 hover:text-teal-500 cursor-pointer" />
      </Link>

      <Trash
        onClick={() => {
          if (confirm("Are you sure?")) {
            deleteThread(thread.ID);
            setMarkdown(thread.Title + "\n" + "[deleted]");
            window.location.reload();
          }
        }}
        className="size-5 hover:text-red-500 cursor-pointer"
      />
    </div>
  ) : (
    <div />
  );
}

export default EditorToolbox;
