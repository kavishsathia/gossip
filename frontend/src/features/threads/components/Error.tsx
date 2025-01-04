import { useSearchParams } from "react-router";
import BRSpeedDial from "./BRSpeedDial";
import { Ghost } from "lucide-react";

function Error() {
  const [search] = useSearchParams();
  return (
    <div className="flex flex-col justify-center items-center h-full ">
      <Ghost className="size-16" />
      <div className="w-2/5 text-center mt-5">
        <h1 className="text-3xl font-semibold">{search.get("status")}</h1>
        <h2 className="text-xl mt-3">{search.get("status-text")}</h2>
      </div>
      <BRSpeedDial />
    </div>
  );
}

export default Error;
