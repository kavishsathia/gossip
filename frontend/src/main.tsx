import { createRoot } from "react-dom/client";
import "./index.css";
import { BrowserRouter, Route, Routes } from "react-router";
import Login from "./features/auth/views/LoginView.tsx";
import SignUp from "./features/auth/views/SignUpView.tsx";
import Layout from "./features/threads/views/LayoutView.tsx";
import Thread from "./features/threads/components/Thread.tsx";
import Editor from "./features/threads/components/Editor.tsx";
import Error from "./features/threads/components/Error.tsx";
import Sidebar from "./features/threads/components/Sidebar.tsx";
import { SnackbarProvider } from "notistack";

import "@fontsource/inter/300.css";
import "@fontsource/inter/400.css";
import "@fontsource/inter/500.css";
import "@fontsource/inter/600.css";
import Welcome from "./features/threads/components/Welcome.tsx";

createRoot(document.getElementById("root")!).render(
  <SnackbarProvider maxSnack={3} preventDuplicate>
    <BrowserRouter>
      <Routes>
        <Route path="/login" element={<Login />} />
        <Route path="/signup" element={<SignUp />} />

        <Route path="/" element={<Layout />}>
          <Route path="/thread/:id" element={<Thread />} />
          <Route path="/editor" element={<Editor />} />
          <Route path="/error" element={<Error />} />
          <Route
            path="/"
            element={
              <div className="h-full">
                <div className="lg:hidden h-full">
                  <Sidebar />{" "}
                </div>
                <div className="hidden lg:block h-full">
                  <Welcome />
                </div>
              </div>
            }
          />
        </Route>
      </Routes>
    </BrowserRouter>
  </SnackbarProvider>
);
