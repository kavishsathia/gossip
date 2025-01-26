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

function Editor({
  markdown,
  setMarkdown,
  editable,
}: {
  markdown: string;
  setMarkdown: React.Dispatch<React.SetStateAction<string>>;
  editable: boolean;
}) {
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
    <MDXEditor
      className="mt-5"
      readOnly={!editable}
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
        linkDialogPlugin(),
        diffSourcePlugin({
          diffMarkdown: markdown,
          viewMode: "rich-text",
        }),
      ].concat(
        editable
          ? [
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
            ]
          : []
      )}
      contentEditableClassName="prose max-w-full text-left mb-10 p-4 px-12"
    />
  );
}

export default Editor;
