import express, { Request, Response } from "express";
import { body } from "express-validator";
import jwt from "jsonwebtoken";

import { User } from "../models/user";
import { BadRequestError } from "@demo.io/lib/build";
import { InternalServerError } from "@demo.io/lib/build";
import { validateRequest } from "@demo.io/lib/build";

const router = express.Router();

router.post(
    "/api/users/signup",
    [
        body("email").isEmail().withMessage("Email must be valid"),
        body("password")
            .trim()
            .isLength({ min: 8, max: 20 })
            .withMessage("Password must be between 4 and 20 characters"),
    ],
    validateRequest,
    async (req: Request, res: Response) => {
        const { email, password } = req.body;

        const existingUser = await User.findOne({ email });

        if (existingUser) {
            throw new BadRequestError("Email already in use");
        }

        if (!process.env.JWT_KEY) {
            throw new InternalServerError("JWT_KEY env var is undefined");
        }

        const user = User.build({ email, password });
        await user.save();

        const userJWT = jwt.sign(
            {
                id: user.id,
                email: user.email,
            },
            process.env.JWT_KEY
        );

        req.session = { jwt: userJWT };

        res.status(201).send(user);
    }
);

export { router as signupRouter };
