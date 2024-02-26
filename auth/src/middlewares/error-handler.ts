import { NextFunction, Request, Response } from "express";
import { CustomError } from "../errors/custom-error";

export const errorHandler = (
    err: CustomError,
    req: Request,
    res: Response,
    next: NextFunction
) => {
    if (err instanceof CustomError) {
        return res.status(400).send({ errors: err.serializeErrors() });
    }

    res.status(400).send({
        message: "Something went wrong",
    });
};
