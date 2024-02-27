import { MongoMemoryServer } from "mongodb-memory-server";
import mongoose from "mongoose";
import request from "supertest";

import { app } from "../app";

let mongo: MongoMemoryServer;

declare global {
    var signup: () => Promise<string[]>;
}

beforeAll(async () => {
    process.env.JWT_KEY = "test";

    mongo = await MongoMemoryServer.create();
    const mongoUri = mongo.getUri();

    await mongoose.connect(mongoUri, {});
});

beforeEach(async () => {
    const collections = await mongoose.connection.db.collections();
    for (const collection of collections) {
        await collection.deleteMany({});
    }
});

afterAll(async () => {
    if (mongo) {
        await mongo.stop();
    }
    await mongoose.connection.close();
});

global.signup = async () => {
    const email = "test@test.com";
    const password = "password";

    const res = await request(app)
        .post("/api/users/signup")
        .send({
            email,
            password,
        })
        .expect(201);

    return res.get("Set-Cookie");
};
