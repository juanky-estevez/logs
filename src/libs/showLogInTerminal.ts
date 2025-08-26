import { dateLongFormat } from "./dateLongFormat";
import { LogType } from "./safeLog";

const colors = {
  error: "\x1b[31m",
  info: "\x1b[34m",
  success: "\x1b[32m",
  warning: "\x1b[33m",
};

export const showLogInTerminal = (type: LogType, message: string): void => {
  const messageLog: string =
    "\x1b[37m" +
    "[" +
    colors[type] +
    type +
    "\x1b[37m" +
    "] " +
    "\x1b[90m" +
    process.env.ENVIRONMENT +
    " " +
    dateLongFormat() +
    " " +
    "\x1b[37m" +
    message +
    "\x1b[0m";

  console.log(messageLog);
};