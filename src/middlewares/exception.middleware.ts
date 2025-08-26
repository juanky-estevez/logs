import {
  ExceptionFilter,
  Catch,
  ArgumentsHost,
  HttpException,
  BadRequestException,
} from '@nestjs/common';
import { Request, Response } from 'express';
import { logError } from '..';

@Catch()
export class ExceptionMiddleware implements ExceptionFilter {
  catch(exception: any, host: ArgumentsHost) {
    const ctx = host.switchToHttp();
    const response = ctx.getResponse<Response>();
    const request = ctx.getRequest<Request>();
    const status =
      exception instanceof HttpException ? exception.getStatus() : 500;

    // Default code y message
    let code = exception.code || 'unknownError';
    let message: any =
      exception.message || exception.code || 'Internal server error';

    if (exception instanceof BadRequestException) {
      const res = exception.getResponse();

      if (typeof res === 'object' && res !== null) {
        if ('message' in res) {
          message = res.message;
        }
        if ('error' in res) {
          code = res.error;
        }
      }
    }

    const errorLog = `[${(request as any).ID || 'NO_ID'}] - status: ${status} - code: ${code} - error: ${
      Array.isArray(message) ? message.join('; ') : message
    }`;
    logError(errorLog);

    const respMessage =
      code === 'Bad Request' && Array.isArray(message) ? message[0] : code;

    response.status(status).json({
      statusCode: status,
      message: respMessage,
    });
  }
}