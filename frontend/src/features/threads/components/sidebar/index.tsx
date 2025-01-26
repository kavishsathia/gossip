import {
  Button,
  CircularProgress,
  InputAdornment,
  TextField,
} from "@mui/material";
import * as React from "react";
import { listThreads } from "../../../../services/threads";
import { Thread } from "../../../../services/threads/types";
import ThreadCard from "./ThreadCard";
import { useSearchParams } from "react-router";
import { Ghost } from "lucide-react";

function App() {
  const [searchParams, setSearchParams] = useSearchParams();
  const search = searchParams.get("search") ?? "";
  const [threads, setThreads] = React.useState<Thread[]>([]);
  const [, setSearch] = React.useState<string>(search);
  const [loading, setLoading] = React.useState<boolean>(false);
  const [page, setPage] = React.useState(1);

  document.addEventListener("keydown", (e) => {
    if ((e.ctrlKey || e.metaKey) && e.key === "k") {
      e.preventDefault();
      document.getElementById("search")?.focus();
    }
  });

  React.useEffect(() => {
    const fetchThreads = async () => {
      const data = await listThreads(search, 1);
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
          label="Search for keywords, #tags or @people"
          variant="outlined"
          className="w-full"
          fullWidth
          size="small"
          onChange={async (e) => {
            setPage(1);
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
      ) : threads.length === 0 ? (
        <div className="text-xl text-center flex flex-col items-center h-full justify-center">
          <Ghost className="size-12" />
          <span className="mt-3 w-2/3">
            Hmm, nothing here yet! Why not start the conversation?
          </span>
        </div>
      ) : (
        <div className="grid grid-cols-1 gap-3 overflow-y-auto overscroll-contain">
          {threads.map((item, index) => (
            <div key={index}>
              <ThreadCard item={item} />
            </div>
          ))}
          {page === 0 ? (
            <div />
          ) : (
            <Button
              onClick={async () => {
                const newThreads = await listThreads(search, page + 1);
                setThreads(threads.concat(newThreads));

                if (newThreads.length < 10) {
                  setPage(0);
                  return;
                }

                setPage(page + 1);
              }}
              fullWidth
              variant="outlined"
            >
              Load more
            </Button>
          )}
        </div>
      )}
    </div>
  );
}

export default App;
