import { Button, TextField } from "@mui/material";
import { useState } from "react";
import { loginAsUser } from "../../../services/threads";

function App() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  return (
    <div className="h-screen w-screen grid grid-cols-2">
      <div className="bg-neutral-900 h-screen"></div>
      <div className="bg-slate-100 h-screen flex justify-center items-center">
        <div className="w-3/5 space-y-5">
          <div className="w-full">
            <h1 className="block text-4xl font-semibold w-full mb-2">Login</h1>
            <h3 className="mb-3">Please enter your credentials to continue</h3>
          </div>
          <TextField
            id="username"
            value={username}
            onChange={(e) => setUsername(e.currentTarget.value)}
            label="Username"
            variant="outlined"
            className="w-full"
            size="small"
          />
          <TextField
            id="password"
            value={password}
            onChange={(e) => setPassword(e.currentTarget.value)}
            label="Password"
            variant="outlined"
            className="w-full"
            size="small"
            type="password"
          />
          <Button
            id="login"
            disabled={username === "" || password === ""}
            onClick={async () => {
              const success = await loginAsUser(username, password);
              if (success) {
                window.location.href = "/";
              }
            }}
            variant="contained"
            fullWidth
            disableElevation
            style={{ textTransform: "none" }}
            sx={{
              bgcolor: "black",
              "&:hover": {
                bgcolor: "#333",
              },
            }}
          >
            Login
          </Button>
          <div className="flex items-center gap-4">
            <div className="h-px bg-gray-300 flex-grow" />
            <span className="text-gray-500">or</span>
            <div className="h-px bg-gray-300 flex-grow" />
          </div>
          <Button
            variant="outlined"
            onClick={() => (window.location.href = "/signup")}
            fullWidth
            disableElevation
            sx={{
              textTransform: "none",
              color: "black",
              bgcolor: "#fff",
              borderColor: "#e0e0e0",
              "&:hover": {
                bgcolor: "#f1f3f4",
                borderColor: "#dadce0",
              },
            }}
          >
            Sign Up
          </Button>
          <p>
            Oops!{" "}
            <a className="cursor-pointer text-neutral-600 hover:text-neutral-900 hover:underline">
              Forgot your password?
            </a>
          </p>
        </div>
      </div>
    </div>
  );
}

export default App;
