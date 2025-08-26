import { appendFileSync, PathOrFileDescriptor } from "fs";
import { dateLongFormat } from "./dateLongFormat";
import { dateShortFormat } from "./dateShortFormat";

export const saveLogInFile = (
  type: string,
  message: string,
  directory: string
): void => {
  // Log file name
  const filename =
    process.env.PROJECT_NAME +
    "_" +
    process.env.ENVIRONMENT +
    "_" +
    dateShortFormat() +
    ".log";

  // Format log
  const messageLog: string =
    "[" +
    type +
    "] " +
    process.env.ENVIRONMENT +
    " " +
    dateLongFormat() +
    " " +
    message +
    "\n";

  // Write on file
  let filePath: PathOrFileDescriptor;
  if (directory[directory.length - 1] === "/") {
    filePath = directory + filename;
  } else {
    filePath = directory + "/" + filename;
  }

  appendFileSync(filePath, messageLog);
};