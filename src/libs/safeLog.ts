import { directory, initOnce } from "./initOnce";
import { saveLogInFile } from "./saveLogInFile";
import { showLogInTerminal } from "./showLogInTerminal";

export type LogType = "error" | "info" | "success" | "warning";

export const safeLog = (type: LogType, message: string) => {
  try {
    initOnce();
    saveLogInFile(type, message, directory);
    showLogInTerminal(type, message);
  } catch (err: any) {
    console.error("\x1b[31m!!!Error!!!\x1b[0m Saving log in file:", err?.message ?? err);
    try { showLogInTerminal(type, message); } catch {}
  }
}