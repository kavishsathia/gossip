import {
  MDXEditor,
  headingsPlugin,
  listsPlugin,
  quotePlugin,
  markdownShortcutPlugin,
  codeBlockPlugin,
  codeMirrorPlugin,
  tablePlugin,
  thematicBreakPlugin,
  linkPlugin,
  linkDialogPlugin,
  diffSourcePlugin,
  frontmatterPlugin,
} from "@mdxeditor/editor";
import "@mdxeditor/editor/style.css";
import { useContext, useState } from "react";
import {
  deleteThread,
  getThread,
  likeThread,
  reportThread,
  unlikeThread,
} from "../../../services/threads";
import React from "react";
import { Heart, MessageCircle, PencilIcon, Share2, Trash } from "lucide-react";
import { Thread } from "../../../services/threads/types";
import Comments from "./Comments";
import { Avatar, Box, CircularProgress } from "@mui/material";
import { Link, useParams, useSearchParams } from "react-router";
import UserInfo from "./UserInfo";
import BRSpeedDial from "../components/BRSpeedDial";
import Welcome from "./Welcome";
import { User } from "../context";
import { enqueueSnackbar } from "notistack";

const padNumber = (num: number) => num.toString().padStart(2, "0");

const getOrdinalSuffix = (day: number) => {
  if (day > 3 && day < 21) return "th";
  switch (day % 10) {
    case 1:
      return "st";
    case 2:
      return "nd";
    case 3:
      return "rd";
    default:
      return "th";
  }
};

// eslint-disable-next-line react-refresh/only-export-components
export const dateTranslate = (dateString: string) => {
  const date = new Date(dateString);
  const day = date.getDate();
  const month = date.toLocaleString("en-US", { month: "short" });
  const year = date.getFullYear();
  const time = `${padNumber(date.getHours())}:${padNumber(date.getMinutes())}`;
  return `${day}${getOrdinalSuffix(day)} ${month} ${year} at ${time}`;
};

function App() {
  const id = Number(useParams().id);
  const [, setSearchParams] = useSearchParams();
  const [markdown, setMarkdown] = useState(``);
  const [thread, setThread] = useState<Thread>();
  const [loading, setLoading] = React.useState(true);
  const user = useContext(User);

  const languages = {
    js: "JavaScript",
    jsx: "React",
    tsx: "React TypeScript",
    python: "Python",
    java: "Java",
    cpp: "C++",
    css: "CSS",
    html: "HTML",
  };

  React.useEffect(() => {
    const fetchThreads = async () => {
      const data = await getThread(id);
      console.log(222);
      setMarkdown("# " + data.Title + "\n" + data.Body);
      setThread(data);
      setLoading(false);
    };

    setLoading(true);
    fetchThreads();
  }, [id]);

  if (loading) {
    return (
      <Box
        sx={{
          display: "flex",
          flexDirection: "column",
          alignItems: "center",
          justifyContent: "center",
          height: "100vh",
          backgroundColor: "#f9fafb",
        }}
      >
        <CircularProgress size={60} thickness={5} sx={{ color: "#1976d2" }} />
      </Box>
    );
  }

  if (!thread?.ID) {
    return <Welcome />;
  }

  return (
    <div className="text-left w-full p-3 lg:p-6 py-2">
      {thread.ModerationFlag && (
        <div className="w-full p-4">
          <div className="flex flex-row flex-wrap gap-y-5 justify-between items-center bg-red-600/20 w-full rounded-md p-5 border-2 border-red-700">
            <div>
              <p className="font-bold">Content Moderation Warning</p>
              <p>
                This thread was flagged for {thread.ModerationFlag}. Please
                report it if the flagging was accurate.
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
      )}

      <div className="flex flex-row justify-between items-start">
        <div className="w-32 h-32 rounded-md m-4 mb-0 mt-6">
          <img
            className="w-full h-full object-cover rounded-md"
            src={thread?.Image || "https://placehold.co/400"}
          ></img>
        </div>
        {thread.UserID === user?.ID && !thread.Deleted ? (
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
        )}
      </div>

      {thread && markdown && (
        <MDXEditor
          key={id}
          readOnly={true}
          markdown={markdown}
          onChange={setMarkdown}
          plugins={[
            headingsPlugin(),
            listsPlugin(),
            quotePlugin(),
            markdownShortcutPlugin(),
            codeBlockPlugin({ defaultCodeBlockLanguage: "js" }),
            codeMirrorPlugin({
              codeBlockLanguages: languages,
            }),
            tablePlugin(),
            thematicBreakPlugin(),
            linkPlugin(),
            frontmatterPlugin(),
            linkDialogPlugin(),
            diffSourcePlugin({
              diffMarkdown: markdown,
              viewMode: "rich-text",
            }),
          ]}
          contentEditableClassName="prose max-w-full text-left h-fit mb-5 p-4 px-5"
        />
      )}
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

      {thread?.ThreadTags && thread?.ThreadTags.length > 0 ? (
        <div className="p-4 pt-0 pb-6">
          <p className="text-xs mb-2">TAGS</p>
          <div className="flex flex-row items-center gap-3">
            {(thread?.ThreadTags ?? []).map((tag, index) => {
              return (
                <span>
                  <span
                    className="hidden lg:block"
                    onClick={() =>
                      setSearchParams({ ["search"]: `#${tag.Tag}` })
                    }
                  >
                    <span
                      key={index}
                      className="cursor-pointer w-fit gap-2 px-2 bg-gray-300 hover:bg-gray-400 rounded-xl text-base font-normal"
                    >
                      {tag.Tag}
                    </span>
                  </span>
                  <Link
                    className="block lg:hidden"
                    to={`/?search=%23${tag.Tag}`}
                  >
                    <span
                      key={index}
                      className="cursor-pointer w-fit gap-2 px-2 bg-gray-300 hover:bg-gray-400 rounded-xl text-base font-normal"
                    >
                      {tag.Tag}
                    </span>
                  </Link>
                </span>
              );
            })}
          </div>
        </div>
      ) : (
        <div></div>
      )}
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
            <p className="hidden md:inline-block">Comment •</p>{" "}
            {thread?.Comments}
          </span>
        </div>
        <div className="space-x-2 w-full flex justify-center">
          <Share2 className="inline" />
          <span>
            <p className="hidden md:inline-block">Share •</p> {thread?.Shares}
          </span>
        </div>
      </div>
      <div className="w-full pt-8 pb-12 lg:px-8">
        <Comments isFromPost={true} id={id} />
      </div>
      <BRSpeedDial />
    </div>
  );
}

export default App;
