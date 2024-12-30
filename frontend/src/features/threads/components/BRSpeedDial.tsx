import SpeedDial from "@mui/material/SpeedDial";
import { Link } from "react-router";
import { Plus } from "lucide-react";

export default function BRSpeedDial() {
  return (
    <Link className="text-white" to="/editor">
      <SpeedDial
        ariaLabel="SpeedDial basic example"
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
        icon={<Plus />}
      ></SpeedDial>
    </Link>
  );
}
