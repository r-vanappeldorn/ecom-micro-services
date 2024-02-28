import express, { Request, Response } from "express";
import { body } from "express-validator";
import jwt from "jsonwebtoken";

import { validateRequest } from "@demo.io/lib/build";
import { User } from "../models/user";
import { BadRequestError } from "@demo.io/lib/build";
import { PasswordManager } from "../services/password-manager";
import { InternalServerError } from "@demo.io/lib/build";

const router = express.Router();

router.post(
    "/api/users/signin",
    [
        body("email").isEmail().withMessage("Email must be valid"),
        body("password").trim().notEmpty().withMessage("Password is required"),
    ],
    validateRequest,
    async (req: Request, res: Response) => {
        const { email, password } = req.body;

        const existingUser = await User.findOne({ email });
        if (!existingUser) {
            throw new BadRequestError("Invalid credentials");
        }

        const passwordMatch = await PasswordManager.compare(
            existingUser.password,
            password
        );
        if (!passwordMatch) {
            throw new BadRequestError("Invalid credentials");
        }

        if (!process.env.JWT_KEY) {
            throw new InternalServerError("JWT_KEY env var is undefined");
        }

        const userJWT = jwt.sign(
            {
                id: existingUser.id,
                email: existingUser.email,
            },
            process.env.JWT_KEY
        );

        req.session = { jwt: userJWT };
        res.status(200).send(existingUser);
    }
);

export { router as signinRouter };
