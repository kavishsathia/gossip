import { Add, Cancel } from "@mui/icons-material";

function MetadataEditor({
  title,
  setTitle,
  tags,
  setTags,
}: {
  title: string;
  tags: string[];
  setTitle: React.Dispatch<React.SetStateAction<string>>;
  setTags: React.Dispatch<React.SetStateAction<string[]>>;
}) {
  return (
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
  );
}

export default MetadataEditor;
