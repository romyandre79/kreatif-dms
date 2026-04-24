/**
 * Simple frontend logger with timestamps and service tags.
 */

const getTimestamp = () => new Date().toISOString();

export const logger = {
  info: (message: string, ...args: any[]) => {
    console.log(`%c[${getTimestamp()}] [INFO] [Frontend]: ${message}`, "color: #3b82f6", ...args);
  },
  warn: (message: string, ...args: any[]) => {
    console.warn(`%c[${getTimestamp()}] [WARN] [Frontend]: ${message}`, "color: #f59e0b", ...args);
  },
  error: (message: string, ...args: any[]) => {
    console.error(`%c[${getTimestamp()}] [ERROR] [Frontend]: ${message}`, "color: #ef4444", ...args);
  },
  debug: (message: string, ...args: any[]) => {
    if (process.env.NODE_ENV !== 'production') {
      console.debug(`%c[${getTimestamp()}] [DEBUG] [Frontend]: ${message}`, "color: #64748b", ...args);
    }
  }
};
