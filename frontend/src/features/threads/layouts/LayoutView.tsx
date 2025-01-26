import { useEffect, useState } from "react";
import Sidebar from "../components/sidebar";
import { Outlet } from "react-router";
import { getMe } from "../../../services/auth";
import { Profile } from "../../../services/auth/types";
import { useSnackbar } from "notistack";
import { User } from "../context";
import { getNotifications } from "../../../services/notifications";
import TopBar from "../components/topbar";

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
  }, [enqueueSnackbar]);

  if (!me) {
    return <div />;
  }

  return (
    <User.Provider value={me}>
      <div className="w-screen h-screen bg-neutral-50 flex flex-col">
        <TopBar me={me} />

        <div className="w-full grid grid-cols-6 flex-1 overflow-hidden overscroll-contain">
          <div className="border-r col-span-2 border-gray-50 hidden lg:block">
            <Sidebar />
          </div>

          <div className="lg:col-span-4 col-span-6 overflow-y-auto overscroll-contain">
            <Outlet />
          </div>
        </div>
      </div>
    </User.Provider>
  );
}

export default App;
