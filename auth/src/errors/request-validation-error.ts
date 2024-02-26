import { ValidationError } from "express-validator";

import { CustomError } from "./custom-error";

export class RequestValidationError extends CustomError {
    constructor(public errors: ValidationError[]) {
        super("Invalid request parameters");

        Object.setPrototypeOf(this, RequestValidationError.prototype);
    }

    statusCode = 400;

    serializeErrors() {
        return this.errors.map((error) => {
            if (error.type === "field") {
                return { message: error.msg, field: error.path };
            }
            return { message: error.msg };
        });
    }
}
