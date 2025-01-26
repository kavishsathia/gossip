const padNumber = (num: number) => num.toString().padStart(2, "0");

const getOrdinalSuffix = (day: number) => {
  if (day > 3 && day < 21) return "th";
  switch (day % 10) {
    case 1:
      return "st";
    case 2:
      return "nd";
    case 3:
      return "rd";
    default:
      return "th";
  }
};

export const dateTranslate = (dateString: string) => {
  const date = new Date(dateString);
  const day = date.getDate();
  const month = date.toLocaleString("en-US", { month: "short" });
  const year = date.getFullYear();
  const time = `${padNumber(date.getHours())}:${padNumber(date.getMinutes())}`;
  return `${day}${getOrdinalSuffix(day)} ${month} ${year} at ${time}`;
};
