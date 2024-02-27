import { CustomError } from "./custom-error";

export class InternalServerError extends CustomError {
    constructor(message: string) {
        super(message);
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
