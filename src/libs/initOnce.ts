import { mkdirSync } from "fs";
import { defineUTC } from "./defineUTC";

let utcReady = false;
let dirReady = false;

export const directory = process.env.LOGS_FOLDER ?? "/logs";

export const initOnce = () => {
  if (!utcReady) {
    defineUTC(); // si esto cambia TZ global, mejor hacerlo 1 sola vez
    utcReady = true;
  }
  if (!dirReady) {
    // recursive evita tener que chequear existsSync
    mkdirSync(directory, { recursive: true });
    dirReady = true;
  }
}