import SpeedDial from "@mui/material/SpeedDial";
import { editThread, postThread } from "../../../../services/threads";
import { Check } from "lucide-react";
import { useState } from "react";
import { CircularProgress } from "@mui/material";

export default function BRSpeedDial({
  id,
  title,
  markdown,
  tags,
  image,
}: {
  id?: number;
  title: string;
  markdown: string;
  tags: string[];
  image: string;
}) {
  const [loading, setLoading] = useState(false);

  return (
    <SpeedDial
      ariaLabel="SpeedDial basic example"
      onClick={async () => {
        setLoading(true);
        if (id) {
          await editThread(id, title, markdown, tags, image);
          window.location.href = `/thread/${id}`;
        } else {
          const id = await postThread(title, markdown, tags, image);
          window.location.href = `/thread/${id}`;
        }
      }}
      className="p-6"
      sx={{
        position: "absolute",
        bottom: 16,
        right: 16,
        color: "white",
        "& .MuiFab-root": {
          backgroundColor: "rgb(15 118 110 / var(--tw-bg-opacity, 1))",
          color: "white",
        },
        "& .MuiFab-root:hover": {
          backgroundColor: "rgb(13 148 136 / var(--tw-bg-opacity, 1))",
        },
      }}
      icon={loading ? <CircularProgress sx={{ color: "white" }} /> : <Check />}
    ></SpeedDial>
  );
}
