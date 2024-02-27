import request from "supertest";
import { app } from "../../app";

it("response with details about current user", async () => {
    const signin = await request(app)
        .post("/api/users/signup")
        .send({
            email: "test@test.com",
            password: "password",
        })
        .expect(201);

    const cookie = signin.get("Set-Cookie");

    const res = await request(app)
        .get("/api/users/currentuser")
        .set("Cookie", cookie)
        .send()
        .expect(200);

    expect(res.body.currentUser.email).toEqual("test@test.com");
    expect(res.body.currentUser.id).toBeDefined();
});
