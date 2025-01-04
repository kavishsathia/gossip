import BRSpeedDial from "./BRSpeedDial";

function Welcome() {
  return (
    <div className="flex justify-center items-center h-full ">
      <div className="w-2/5 text-center">
        <h1 className="text-3xl font-semibold">Yo, welcome!</h1>
        <h2 className="text-xl mt-3">
          Explore the threads on the left or click the plus button at the
          bottom-right corner to get started.
        </h2>
      </div>
      <BRSpeedDial />
    </div>
  );
}

export default Welcome;
