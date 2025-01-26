import { Thread } from "../../../../services/threads/types";
import { Link, useSearchParams } from "react-router";

function ModerationFlag({ thread }: { thread: Thread }) {
  const [, setSearchParams] = useSearchParams();
  return thread?.ThreadTags && thread?.ThreadTags.length > 0 ? (
    <div className="p-4 pt-0 pb-6">
      <p className="text-xs mb-2">TAGS</p>
      <div className="flex flex-row items-center gap-3">
        {(thread?.ThreadTags ?? []).map((tag, index) => {
          return (
            <span>
              <span
                className="hidden lg:block"
                onClick={() => setSearchParams({ ["search"]: `#${tag.Tag}` })}
              >
                <span
                  key={index}
                  className="cursor-pointer w-fit gap-2 px-2 bg-gray-300 hover:bg-gray-400 rounded-xl text-base font-normal"
                >
                  {tag.Tag}
                </span>
              </span>
              <Link className="block lg:hidden" to={`/?search=%23${tag.Tag}`}>
                <span
                  key={index}
                  className="cursor-pointer w-fit gap-2 px-2 bg-gray-300 hover:bg-gray-400 rounded-xl text-base font-normal"
                >
                  {tag.Tag}
                </span>
              </Link>
            </span>
          );
        })}
      </div>
    </div>
  ) : (
    <div></div>
  );
}

export default ModerationFlag;
