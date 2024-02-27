import { Request, Response, NextFunction } from "express";
import { InternalServerError } from "../errors/internal-server-error";
import jwt from "jsonwebtoken";

interface UserPayload {
    id: string;
    email: string;
}

declare global {
    namespace Express {
        interface Request {
            currentUser?: UserPayload;
        }
    }
}

export const currentUser = (
    req: Request,
    res: Response,
    next: NextFunction
) => {
    if (!process.env.JWT_KEY) {
        throw new InternalServerError("JWT_KEY env var is undefined");
    }

    if (!req.session?.jwt) {
        return next();
    }

    try {
        const payload = jwt.verify(
            req.session.jwt,
            process.env.JWT_KEY
        ) as UserPayload;
        req.currentUser = payload;
    } catch (err) {}
    next();
};
