import {
  MDXEditor,
  headingsPlugin,
  listsPlugin,
  quotePlugin,
  markdownShortcutPlugin,
  codeBlockPlugin,
  codeMirrorPlugin,
  tablePlugin,
  thematicBreakPlugin,
  linkPlugin,
  toolbarPlugin,
  linkDialogPlugin,
  diffSourcePlugin,
  frontmatterPlugin,
  BlockTypeSelect,
  BoldItalicUnderlineToggles,
  CodeToggle,
  CreateLink,
  ListsToggle,
  UndoRedo,
  InsertTable,
  InsertCodeBlock,
  ConditionalContents,
  ChangeCodeMirrorLanguage,
} from "@mdxeditor/editor";
import "@mdxeditor/editor/style.css";
import {
  Button,
  Dialog,
  DialogActions,
  DialogContent,
  DialogContentText,
  DialogTitle,
  SpeedDial,
  TextField,
} from "@mui/material";
import { useState } from "react";
import { postThread } from "../../../services/threads";
import { Check, Upload } from "lucide-react";
import { Add, Cancel } from "@mui/icons-material";

function App() {
  const [markdown, setMarkdown] = useState("");
  const [title, setTitle] = useState("Untitled Thread");
  const [image, setImage] = useState("https://placehold.co/400");
  const [imageBuffer, setImageBuffer] = useState("");
  const [imageModal, setImageModal] = useState(false);
  const [tags, setTags] = useState(["gossip"]);

  // Define supported languages
  const languages = {
    js: "JavaScript",
    jsx: "React",
    tsx: "React TypeScript",
    python: "Python",
    java: "Java",
    cpp: "C++",
    css: "CSS",
    html: "HTML",
  };

  return (
    <div className="text-left w-full p-4 pt-6">
      <div className="flex flex-row gap-5 items-center ">
        <div
          onClick={() => setImageModal(true)}
          className="relative w-24 h-24 rounded-md"
        >
          <img
            className="w-full h-full object-cover rounded-md"
            src={image}
          ></img>
          <span className="text-white z-1 absolute top-0 left-0 w-full h-full flex items-center justify-center bg-gray-500/50 rounded-md">
            <Upload className="stroke-2" />
          </span>
        </div>
        <Dialog open={imageModal} onClose={() => setImageModal(false)}>
          <DialogTitle>Add an Image</DialogTitle>
          <DialogContent>
            <DialogContentText>
              Find an image on the Internet and add it here!
            </DialogContentText>
            <TextField
              value={imageBuffer}
              onChange={(e) => setImageBuffer(e.currentTarget.value)}
              autoFocus
              required
              margin="dense"
              label="Image URL"
              type="url"
              fullWidth
              variant="standard"
            />
          </DialogContent>
          <DialogActions>
            <Button
              onClick={() => {
                setImageModal(false);
              }}
            >
              Cancel
            </Button>
            <Button
              onClick={() => {
                setImage(imageBuffer);
                setImageBuffer("");
              }}
              type="submit"
            >
              Add
            </Button>
          </DialogActions>
        </Dialog>
        <div>
          <div
            contentEditable
            suppressContentEditableWarning={true}
            className="text-3xl font-bold text-black focus:ring-0 focus:outline-none"
            onBlur={(e) => setTitle(e.currentTarget.textContent || "")}
          >
            {title}
          </div>
          <div className="mt-3 focus:ring-0 focus:outline-none flex flex-row flex-wrap gap-2">
            {tags.map((item, index) => {
              return (
                <span
                  contentEditable
                  key={index}
                  suppressContentEditableWarning={true}
                  onBlur={(e) => {
                    const tmp = tags;
                    tmp[index] = e.currentTarget.textContent ?? "";
                    console.log(tags);
                    setTags(tmp);
                  }}
                  className="flex flex-row items-center w-fit gap-2 px-2 bg-gray-300 rounded-xl text-base font-normal"
                >
                  {item}
                  <span
                    onClick={() => {
                      setTags([...tags.filter((_, i) => i !== index)]);
                      console.log([...tags]);
                      console.log([...tags.filter((_, i) => i !== index)]);
                    }}
                  >
                    <Cancel sx={{ width: 15 }} />
                  </span>
                </span>
              );
            })}
            <span className="flex flex-row items-center w-fit gap-2 px-2 bg-gray-300 rounded-xl text-base font-normal">
              <span onClick={() => setTags([...tags, "new_tag"])}>
                <Add sx={{ width: 15 }} />
                Tags
              </span>
            </span>
          </div>
        </div>
      </div>
      <MDXEditor
        className="mt-5"
        markdown={markdown}
        onChange={setMarkdown}
        plugins={[
          headingsPlugin({ allowedHeadingLevels: [2, 3, 4, 5, 6] }),
          listsPlugin(),
          quotePlugin(),
          markdownShortcutPlugin(),
          codeBlockPlugin({ defaultCodeBlockLanguage: "js" }),
          codeMirrorPlugin({
            codeBlockLanguages: languages,
          }),
          tablePlugin(),
          thematicBreakPlugin(),
          linkPlugin(),
          frontmatterPlugin(),
          toolbarPlugin({
            toolbarContents: () => (
              <div className="flex flex-wrap gap-2">
                <UndoRedo />
                <BlockTypeSelect />
                <BoldItalicUnderlineToggles />
                <CodeToggle />
                <CreateLink />
                <InsertTable />
                <ListsToggle />
                <ConditionalContents
                  options={[
                    {
                      when: (editor) => editor?.editorType === "codeblock",
                      contents: () => <ChangeCodeMirrorLanguage />,
                    },
                    {
                      fallback: () => (
                        <>
                          <InsertCodeBlock />
                        </>
                      ),
                    },
                  ]}
                />
              </div>
            ),
          }),
          linkDialogPlugin(),
          diffSourcePlugin({
            diffMarkdown: markdown,
            viewMode: "rich-text",
          }),
        ]}
        contentEditableClassName="prose max-w-full text-left min-h-[300px] p-4 px-12"
      />
      <SpeedDial
        ariaLabel="SpeedDial basic example"
        onClick={async () => {
          const id = await postThread(
            title,
            `# ${title}\n` + markdown,
            tags,
            image
          );
          window.location.href = `/thread/${id}`;
        }}
        className="p-6"
        sx={{
          position: "absolute",
          bottom: 16,
          right: 16,
          color: "white",
          "& .MuiFab-root": {
            backgroundColor: "black",
            color: "white",
          },
          "& .MuiFab-root:hover": {
            backgroundColor: "#333",
          },
        }}
        icon={<Check />}
      ></SpeedDial>
    </div>
  );
}

export default App;
