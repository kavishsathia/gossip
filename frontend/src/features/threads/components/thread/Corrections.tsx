import { Thread } from "../../../../services/threads/types";

function Corrections({ thread }: { thread: Thread }) {
  return thread.ThreadCorrections.length > 0 ? (
    <div className="w-full p-4">
      <div className="flex flex-row flex-wrap gap-y-5 justify-between items-center bg-teal-200/10 w-full rounded-md p-5 border-2 border-teal-700">
        <div>
          <p className="font-bold">Corrections</p>
          <ul className="list-decimal ml-5">
            {thread.ThreadCorrections.map((item) => (
              <li>{item.Correction}</li>
            ))}
          </ul>
        </div>
      </div>
    </div>
  ) : (
    <div />
  );
}

export default Corrections;
