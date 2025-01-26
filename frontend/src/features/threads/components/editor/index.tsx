import "@mdxeditor/editor/style.css";
import { Box, CircularProgress } from "@mui/material";
import { useEffect, useState } from "react";
import { getThread } from "../../../../services/threads";

import { useParams } from "react-router";
import PostSpeedDial from "./PostSpeedDial";
import ImageSelectionModal from "./ImageSelectionModal";
import Editor from "../commons/Editor";
import MetadataEditor from "./MetadataEditor";

function App() {
  const id = Number(useParams().id) || undefined;
  const [markdown, setMarkdown] = useState("");
  const [title, setTitle] = useState("Untitled Thread");
  const [image, setImage] = useState("https://placehold.co/400");
  const [tags, setTags] = useState(["gossip"]);
  const [loading, setLoading] = useState(false);

  useEffect(() => {
    if (id) {
      const fetchThreads = async () => {
        const data = await getThread(id);
        setTags(data.ThreadTags.map((item) => item.Tag));
        setMarkdown(data.Body);
        setImage(data.Image ?? "https://placehold.co/400");
        setTitle(data.Title);
        setLoading(false);
      };

      setLoading(true);
      fetchThreads();
    }
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

  return (
    <div className="text-left w-full p-4 pt-6">
      <div className="flex flex-row gap-5 items-center ">
        <ImageSelectionModal image={image} setImage={setImage} />
        <MetadataEditor
          title={title}
          tags={tags}
          setTags={setTags}
          setTitle={setTitle}
        />
      </div>
      <Editor markdown={markdown} setMarkdown={setMarkdown} editable={true} />
      <PostSpeedDial
        id={id}
        title={title}
        markdown={markdown}
        tags={tags}
        image={image}
      />
    </div>
  );
}

export default App;
