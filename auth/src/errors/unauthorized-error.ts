import { CustomError } from "./custom-error";

export class UnauthorizedError extends CustomError {
    constructor() {
        super("User is unauthorized");

        Object.setPrototypeOf(this, UnauthorizedError.prototype);
    }

    statusCode = 401;

    serializeErrors() {
        return [{ message: "Not authorized" }];
    }
}
