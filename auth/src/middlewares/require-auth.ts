import { Request, NextFunction, Response } from "express";
import { UnauthorizedError } from "../errors/unauthorized-error";

/**
 * !Impotent! only use after currentUser middleware
 *
 * @param req
 * @param res
 * @param next
 */
export const requireAuth = (
    req: Request,
    res: Response,
    next: NextFunction
) => {
    if (!req.currentUser) {
        throw new UnauthorizedError();
    }

    next();
};
