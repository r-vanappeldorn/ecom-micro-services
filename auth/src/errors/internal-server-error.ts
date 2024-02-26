import { CustomError } from "./custom-error";

export class InternalServerError extends CustomError {
    constructor() {
        super("Internal server error");
    }

    statusCode = 500;

    serializeErrors() {
        return [
            {
                message: "Something went wrong on the server",
            },
        ];
    }
}
