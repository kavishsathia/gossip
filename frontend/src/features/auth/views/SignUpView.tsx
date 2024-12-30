import { Button, TextField } from "@mui/material";
import { useState } from "react";
import { createUser } from "../../../services/threads";

function isStrongPassword(password: string) {
  const minLength = 8;

  const hasUppercase = /[A-Z]/.test(password);
  const hasLowercase = /[a-z]/.test(password);
  const hasDigit = /\d/.test(password);
  const hasSpecialChar = /[!@#$%^&*(),.?":{}|<>]/.test(password);
  const hasMinLength = password.length >= minLength;

  return (
    hasUppercase && hasLowercase && hasDigit && hasSpecialChar && hasMinLength
  );
}

function App() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");
  const [confirmPassword, setConfirmPassword] = useState("");

  return (
    <div className="h-screen w-screen grid lg:grid-cols-2">
      <div className="bg-neutral-900 h-screen hidden lg:block"></div>
      <div className="bg-slate-100 h-screen flex justify-center items-center">
        <div className="w-4/5 md:w-2/5 lg:w-3/5 space-y-5">
          <div className="w-full">
            <h1 className="block text-4xl font-semibold w-full mb-2">
              Sign Up
            </h1>
            <h3 className="mb-3">Please enter your credentials to continue</h3>
          </div>
          <TextField
            value={username}
            onChange={(e) => setUsername(e.currentTarget.value)}
            label="Username"
            variant="outlined"
            className="w-full"
            size="small"
          />
          <TextField
            error={password !== "" && !isStrongPassword(password)}
            value={password}
            onChange={(e) => setPassword(e.currentTarget.value)}
            type="password"
            label="Password"
            variant="outlined"
            className="w-full"
            size="small"
            helperText={
              password !== "" && !isStrongPassword(password)
                ? "The password needs to have one uppercase letter, lowercase letter, digit and special character and at least 8 characters long"
                : ""
            }
          />
          <TextField
            error={confirmPassword !== "" && confirmPassword !== password}
            value={confirmPassword}
            onChange={(e) => setConfirmPassword(e.currentTarget.value)}
            type="password"
            label="Confirm Password"
            variant="outlined"
            className="w-full"
            size="small"
            helperText={
              confirmPassword !== "" && confirmPassword !== password
                ? "The passwords do not match"
                : ""
            }
          />
          <Button
            onClick={async () => {
              if (confirmPassword === password) {
                const res = await createUser(username, password);
                if (res < 0) {
                  const newUsername =
                    username + Math.floor(Math.random() * (99 + 1));
                  alert(
                    `Oops! The username is taken, try another one, maybe ${newUsername}?`
                  );
                } else {
                  window.location.href = "/login";
                }
              }
            }}
            variant="contained"
            disabled={
              password === "" ||
              confirmPassword === "" ||
              !isStrongPassword(password) ||
              confirmPassword !== password
            }
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
            Sign Up
          </Button>
          <div className="flex items-center gap-4">
            <div className="h-px bg-gray-300 flex-grow" />
            <span className="text-gray-500">or</span>
            <div className="h-px bg-gray-300 flex-grow" />
          </div>
          <Button
            variant="outlined"
            onClick={() => (window.location.href = "/login")}
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
            Login
          </Button>
        </div>
      </div>
    </div>
  );
}

export default App;
