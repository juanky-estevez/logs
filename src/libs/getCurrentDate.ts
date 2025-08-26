import { utcOffset } from "./defineUTC";

export const getCurrentDate = (): Date => {
  const dateUtc_0 = new Date();
  const currentDate = new Date(
    dateUtc_0.getTime() + utcOffset * 60 * 60 * 1000
  );
  return currentDate;
};