import "@mdxeditor/editor/style.css";
import { useState } from "react";
import { getThread } from "../../../../services/threads";
import React from "react";
import { Thread } from "../../../../services/threads/types";
import Comments from "../comments";
import { Box, CircularProgress } from "@mui/material";
import { useParams } from "react-router";
import BRSpeedDial from "../commons/CreateSpeedDial";
import Welcome from "../commons/Welcome";
import Editor from "../commons/Editor";
import ModerationFlag from "./ModerationFlag";
import Corrections from "./Corrections";
import EditorToolbox from "./EditorToolbox";
import InteractionBar from "./InteractionBar";
import TagList from "./TagList";
import AuthorInfo from "./AuthorInfo";

function App() {
  const id = Number(useParams().id);
  const [markdown, setMarkdown] = useState(``);
  const [thread, setThread] = useState<Thread>();
  const [loading, setLoading] = React.useState(true);

  React.useEffect(() => {
    const fetchThreads = async () => {
      const data = await getThread(id);
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
      <ModerationFlag thread={thread} />

      <Corrections thread={thread} />

      <div className="flex flex-row justify-between items-start">
        <div className="w-32 h-32 rounded-md m-4 mb-0 mt-6">
          <img
            className="w-full h-full object-cover rounded-md"
            src={thread?.Image || "https://placehold.co/400"}
          ></img>
        </div>
        <EditorToolbox thread={thread} setMarkdown={setMarkdown} />
      </div>

      {thread && markdown && (
        <Editor
          markdown={markdown}
          setMarkdown={setMarkdown}
          editable={false}
        />
      )}

      <AuthorInfo thread={thread} />

      <TagList thread={thread} />

      <InteractionBar thread={thread} setThread={setThread} />

      <div className="w-full pt-8 pb-12 lg:px-8">
        <Comments isFromPost={true} id={id} depth={0} />
      </div>

      <BRSpeedDial />
    </div>
  );
}

export default App;
