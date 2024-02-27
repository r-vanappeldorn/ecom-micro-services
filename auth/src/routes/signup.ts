import express, { Request, Response } from "express";
import { body } from "express-validator";
import jwt from "jsonwebtoken";

import { User } from "../models/user";
import { BadRequestError } from "../errors/bad-request-error";
import { InternalServerError } from "../errors/internal-server-error";
import { validateRequest } from "../middlewares/validate-request";

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
            throw new InternalServerError();
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
