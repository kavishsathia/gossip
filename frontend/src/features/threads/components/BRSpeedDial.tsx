import SpeedDial from "@mui/material/SpeedDial";
import { Link } from "react-router";
import { Plus } from "lucide-react";

export default function BRSpeedDial() {
  return (
    <SpeedDial
      ariaLabel="SpeedDial basic example"
      className="p-6"
      sx={{
        position: "absolute",
        bottom: 16,
        right: 16,
        color: "white",
        "& .MuiFab-root": {
          backgroundColor: "black",
          color: "white",
        },
        "& .MuiFab-root:hover": {
          backgroundColor: "#333",
        },
      }}
      icon={
        <Link className="text-white" to="/editor">
          <Plus />
        </Link>
      }
    ></SpeedDial>
  );
}
