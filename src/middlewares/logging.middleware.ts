import { NextFunction, Request, Response } from 'express';
import { NestMiddleware } from '@nestjs/common';
import { nanoid } from 'nanoid';
import { logInfo } from '..';

export class LoggingMiddleware implements NestMiddleware {
  use(req: Request, res: Response, next: NextFunction) {
    const start = Date.now();
    const { body, method, originalUrl, params } = req;

    const reqID = nanoid(8);
    req.requestId = reqID;

    const reqLog =
      `[${reqID}] ${method} ${originalUrl}` +
      (body ? ' - ' + JSON.stringify(sanitizeBody(body)) : '') +
      (Object.keys(params).length !== 0 ? ' - ' + JSON.stringify(params) : '');

    logInfo(reqLog);

    const originalSend = res.send;
    res.send = function (body?: any): any {
      const duration = Date.now() - start;
      const statusCode = res.statusCode;
      const parsedBody = safeParseResponse(body);

      if (statusCode < 400) {
        const debugMessage =
          process.env.DEBUG === 'ON'
            ? ` - response: ${JSON.stringify(parsedBody)}}`
            : '';

        const infoLog = `[${reqID}] duration: ${duration}ms ${debugMessage}`;
        logInfo(infoLog);
      }

      return originalSend.call(this, body);
    };

    next();
  }
}

const safeParseResponse = (body: any): any => {
  try {
    return typeof body === 'string' ? JSON.parse(body) : body;
  } catch {
    return body;
  }
};

const sanitizeBody = (body: any): any => {
  const SENSITIVE_KEYS = ['password', 'oldPassword', 'newPassword'];

  if (Array.isArray(body)) {
    return body.map(sanitizeBody);
  }

  if (typeof body === 'object' && body !== null) {
    const sanitized: any = {};
    for (const key of Object.keys(body)) {
      if (SENSITIVE_KEYS.includes(key)) {
        sanitized[key] = '****';
      } else {
        sanitized[key] = sanitizeBody(body[key]);
      }
    }
    return sanitized;
  }

  return body;
};