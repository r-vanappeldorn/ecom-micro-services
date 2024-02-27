import express from "express";
import jwt from "jsonwebtoken";
import { InternalServerError } from "../errors/internal-server-error";

const router = express.Router();

router.get("api/users/currentuser", (req, res) => {
    if (!process.env.JWT_KEY) {
        throw new InternalServerError("JWT_KEY env var is undefined");
    }

    if (!req.session?.jwt) {
        return res.send({ currentUser: null });
    }

    try {
        const payload = jwt.verify(req.session.jwt, process.env.JWT_KEY);
        res.status(200).send({ currentUser: payload });
    } catch (err) {
        res.send({ currentUser: null });
    }
});

export { router as currentUserRouter };
