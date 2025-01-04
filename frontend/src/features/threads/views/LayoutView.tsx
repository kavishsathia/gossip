import { useEffect, useState } from "react";
import Sidebar from "../components/Sidebar";
import { Link, Outlet } from "react-router";
import { signOut, getMe, getNotifications } from "../../../services/threads";
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

    const ws = getNotifications();
    ws.onmessage = function (event) {
      console.log(event.data);
      enqueueSnackbar(event.data, {
        preventDuplicate: true,
      });
    };
  }, []);

  return (
    <div className="w-screen h-screen bg-neutral-50 flex flex-col">
      <div className="h-12 w-full flex justify-between items-center border-b-2 border-neutral-200 flex-shrink-0 px-5 bg-teal-700">
        <LogOut
          onClick={signOut}
          className="rotate-180 text-white cursor-pointer"
        />
        <Link to="/">
          <img className="h-6 w-auto" src="/logo.png" alt="Logo" />
        </Link>
        <Avatar
          src={me?.ProfileImage}
          sx={{ width: "35px", height: "35px" }}
        ></Avatar>
      </div>

      <div className="w-full grid grid-cols-6 flex-1 overflow-hidden overscroll-contain">
        <div className="border-r col-span-2 border-gray-50 hidden lg:block">
          <Sidebar />
        </div>

        <div className="lg:col-span-4 col-span-6 overflow-y-auto overscroll-contain">
          <Outlet />
        </div>
      </div>
    </div>
  );
}

export default App;
