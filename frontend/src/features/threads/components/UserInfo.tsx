import { Avatar, Box, Button, Modal, Typography } from "@mui/material";
import { useEffect, useState } from "react";
import { getUser } from "../../../services/auth";
import { Profile } from "../../../services/auth/types";
import { dateTranslate } from "./Thread";
import { useSearchParams } from "react-router";

function App({ username, userId }: { username: string; userId: number }) {
  const [open, setOpen] = useState(false);
  const [user, setUser] = useState<Profile>();
  const handleOpen = () => setOpen(true);
  const handleClose = () => setOpen(false);
  const [, setSearchParams] = useSearchParams();

  useEffect(() => {
    const fetchThreads = async () => {
      const data = await getUser(userId);
      setUser(data);
    };

    fetchThreads();
  }, [username, userId]);

  return (
    <div>
      <p
        onClick={handleOpen}
        className="cursor-pointer hover:underline text-sm h-fit font-semibold"
      >
        {username}
      </p>
      <Modal
        open={open}
        onClose={handleClose}
        aria-labelledby="modal-modal-title"
        aria-describedby="modal-modal-description"
      >
        <Box
          sx={{
            position: "absolute",
            top: "50%",
            left: "50%",
            transform: "translate(-50%, -50%)",
            width: 400,
            bgcolor: "background.paper",
            boxShadow: 24,
            p: 4,
          }}
        >
          <Avatar
            src={user?.ProfileImage}
            sx={{ width: "70px", height: "70px" }}
          ></Avatar>
          <Typography
            id="modal-modal-title"
            variant="h5"
            component="h1"
            sx={{ mt: 2 }}
          >
            {user?.Username}
          </Typography>
          <Typography id="modal-modal-description" sx={{ mt: 0.5 }}>
            Joined on {dateTranslate(user?.CreatedAt ?? "")}
          </Typography>
          <div className="flex flex-row items-center gap-3 my-4 justify-between">
            <div className="flex flex-col min-w-16">
              <p className="text-2xl font-semibold">{user?.Posts}</p>
              <p className="text-base font-bold">Threads</p>
            </div>

            <div className="h-16 w-[1px] bg-black"></div>

            <div className="flex flex-col min-w-16">
              <p className="text-2xl font-semibold">{user?.Comments}</p>
              <p className="text-base font-bold">Comments</p>
            </div>

            <div className="h-16 w-[1px] bg-black"></div>

            <div className="flex flex-col min-w-16">
              <p className="text-2xl font-semibold">{user?.Aura}</p>
              <p className="text-base font-bold">Aura</p>
            </div>
          </div>
          <div className="hidden lg:block">
            <Button
              onClick={() => {
                setSearchParams({ ["search"]: `@${username}` });
                handleClose();
              }}
              variant="contained"
              disableElevation
              fullWidth
              size="medium"
              sx={{ mt: 2, background: "#111" }}
            >
              View Their Threads
            </Button>
          </div>
          <div className="block lg:hidden">
            <Button
              onClick={() => {
                window.location.href = `/?search=@${username}`;
              }}
              variant="contained"
              disableElevation
              fullWidth
              size="medium"
              sx={{ mt: 2, background: "#111" }}
            >
              View Their Threads
            </Button>
          </div>
        </Box>
      </Modal>
    </div>
  );
}

export default App;
