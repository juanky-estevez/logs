import { safeLog } from "./libs/safeLog";

export const logError = (message: string): void => safeLog("error", message);
export const logInfo  = (message: string): void => safeLog("info", message);
export const logSuccess = (message: string): void => safeLog("success", message);
export const logWarning = (message: string): void => safeLog("warning", message);

export * from './middlewares';