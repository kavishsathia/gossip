import { useEffect, useState } from "react";
import Sidebar from "../components/Sidebar";
import { Outlet } from "react-router";
import { getMe } from "../../../services/threads";
import { Profile } from "../../../services/threads/types";
import { useSnackbar } from "notistack";
import { LogOut } from "lucide-react";
import { Avatar } from "@mui/material";

function App() {
  const [me, setMe] = useState<Profile>();
  const { enqueueSnackbar } = useSnackbar();

  useEffect(() => {
    const fetchMe = async () => {
      const data = await getMe();
      if (data === null) {
        window.location.href = "/login";
        return;
      }
      setMe(data);
    };

    fetchMe();

    const ws = new WebSocket("ws://localhost:8080/notifications");
    ws.onmessage = function (event) {
      console.log(event.data);
      enqueueSnackbar(event.data, {
        preventDuplicate: true,
      });
    };
  }, []);

  return (
    <div className="w-screen h-screen bg-neutral-50 flex flex-col">
      <div className="h-12 w-full flex justify-between items-center border-b-2 border-neutral-200 flex-shrink-0 px-5 bg-neutral-900">
        <LogOut className="rotate-180 text-white" />
        <img
          className="h-5 w-auto filter invert"
          src="/public/logo.png"
          alt="Logo"
        />
        <Avatar
          src={me?.ProfileImage}
          sx={{ width: "35px", height: "35px" }}
        ></Avatar>
      </div>

      <div className="w-full grid grid-cols-6 flex-1 overflow-hidden">
        <div className="border-r col-span-2 border-gray-50 hidden lg:block">
          <Sidebar />
        </div>

        <div className="lg:col-span-4 col-span-6 overflow-y-auto">
          <Outlet />
        </div>
      </div>
    </div>
  );
}

export default App;
