import { LogOut } from "lucide-react";
import { signOut } from "../../../../services/auth";
import { Link } from "react-router";
import { Profile } from "../../../../services/auth/types";
import UserInfo from "../commons/UserInfo";

function TopBar({ me }: { me: Profile }) {
  return (
    <div
      className="h-12 w-full flex justify-between items-center 
        border-b-2 border-neutral-200 flex-shrink-0 px-5 bg-teal-700"
    >
      <LogOut
        onClick={() => {
          if (confirm("Are you sure you want to sign out?")) {
            signOut();
          }
        }}
        className="rotate-180 text-white cursor-pointer"
      />
      <Link to="/">
        <img className="h-6 w-auto" src="/logo.png" alt="Logo" />
      </Link>
      <UserInfo
        width="35px"
        picture={true}
        userId={me.ID ?? 0}
        username={me.Username ?? ""}
      />
    </div>
  );
}

export default TopBar;
