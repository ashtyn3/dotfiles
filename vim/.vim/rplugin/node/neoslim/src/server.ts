import fast from "fastify";
import fs from "fs";
import { AccessInfo } from ".";
import axios from "axios";
import os from "os";

export const server = () => {
    const server = fast({});
    server.route({
        method: "GET",
        url: "/",
        schema: {
            querystring: {
                code: { type: "string" },
            },
        },
        handler: async (req, res) => {
            const params: any = req.query;
            fs.writeFileSync("./exchange.json", JSON.stringify(params));
            const file = fs.readFileSync("./exchange.json", "utf-8");
            const code = JSON.parse(file).code;
            const response = await axios.post(
                "https://slack.com/api/oauth.v2.access?client_id=1879759933815.2001379021968&client_secret=5228b2fbd77bdb266e5277917a493363&code=" +
                    code
            );
            const data: AccessInfo = response.data;
            fs.writeFileSync(
                os.homedir() + "/.neoslim.json",
                JSON.stringify(data)
            );
            fs.unlinkSync("./exchange.json");
            res.send("Done");
        },
    });
    return server;
};
