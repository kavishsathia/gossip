import { CircularProgress, InputAdornment, TextField } from "@mui/material";
import * as React from "react";
import { listThreads } from "../../../services/threads";
import { Thread } from "../../../services/threads/types";
import ThreadCard from "../components/ThreadCard";
import { useSearchParams } from "react-router";

function App() {
  const [searchParams, setSearchParams] = useSearchParams();
  const search = searchParams.get("search") ?? "";
  const [threads, setThreads] = React.useState<Thread[]>([]);
  const [, setSearch] = React.useState<string>(search);
  const [loading, setLoading] = React.useState<boolean>(false);

  document.addEventListener("keydown", (e) => {
    if ((e.ctrlKey || e.metaKey) && e.key === "k") {
      e.preventDefault();
      document.getElementById("search")?.focus();
    }
  });

  React.useEffect(() => {
    const fetchThreads = async () => {
      const data = await listThreads(encodeURIComponent(search));
      setThreads(data);
      setLoading(false);
    };

    setLoading(true);
    fetchThreads();
  }, [search]);

  return (
    <div className="bg-teal-50/75 h-[calc(100vh-3rem)] w-full border-r-2 border-neutral-200 p-5 flex flex-col">
      <div className="mb-5">
        <TextField
          value={search}
          id="search"
          label="Search"
          variant="outlined"
          className="w-full"
          fullWidth
          size="small"
          onChange={async (e) => {
            setLoading(true);
            setSearch(e.target.value);
            setSearchParams({ ["search"]: e.target.value });
            setLoading(false);
          }}
          slotProps={{
            input: {
              endAdornment: (
                <InputAdornment position="end">
                  <span className="px-2 py-0.5 flex items-center justify-center rounded-md bg-gray-200 text-xs">
                    ⌘K
                  </span>
                </InputAdornment>
              ),
            },
          }}
        />
      </div>
      {loading ? (
        <div className="w-full flex justify-center h-full items-center">
          <CircularProgress />
        </div>
      ) : (
        <div className="grid grid-cols-1 gap-3 overflow-y-auto overscroll-contain">
          {threads.map((item, index) => (
            <div key={index}>
              <ThreadCard item={item} />
            </div>
          ))}
        </div>
      )}
    </div>
  );
}

export default App;
