import { getCurrentDate } from "./getCurrentDate";

export const dateShortFormat = (): string => {
  const currentDate = getCurrentDate();

  const formattedDate =
    currentDate.getFullYear().toString() +
    (currentDate.getMonth() + 1).toString().padStart(2, "0") +
    currentDate.getDate().toString().padStart(2, "0");

  // Response
  return formattedDate;
};