import { NvimPlugin } from "neovim";
import fs from "fs";
import open from "open";
import { server } from "./server";
import os from "os";
import axios from "axios";
import { chat } from "./chat";

export interface AccessInfo {
    ok: boolean;
    app_id: string;
    authed_user: AuthedUser;
    team: Team;
    enterprise: null;
    is_enterprise_install: boolean;
}

export interface AuthedUser {
    id: string;
    scope: string;
    access_token: string;
    token_type: string;
}

export interface Team {
    id: string;
    name: string;
}

export interface User {
    ok: boolean;
    user: UserClass;
}

export interface UserClass {
    id: string;
    team_id: string;
    name: string;
    deleted: boolean;
    color: string;
    real_name: string;
    tz: string;
    tz_label: string;
    tz_offset: number;
    profile: Profile;
    is_admin: boolean;
    is_owner: boolean;
    is_primary_owner: boolean;
    is_restricted: boolean;
    is_ultra_restricted: boolean;
    is_bot: boolean;
    updated: number;
    is_app_user: boolean;
    has_2fa: boolean;
}

export interface Profile {
    avatar_hash: string;
    status_text: string;
    status_emoji: string;
    real_name: string;
    display_name: string;
    real_name_normalized: string;
    display_name_normalized: string;
    email: string;
    image_original: string;
    image_24: string;
    image_32: string;
    image_48: string;
    image_72: string;
    image_192: string;
    image_512: string;
    team: string;
}

type Method = "get" | "post";
export const call = async (endpoint: string, tok: string, params?: string) => {
    const url = "https://slack.com/api/" + endpoint + "?" + params;
    var config = {
        headers: {
            Accept: "application/json",
            Authorization: "Bearer " + tok,
            "Content-Type": "application/json",
            Cookie: "b=1ewceef27en5uy5tnd74sh0eh",
        },
    };

    return (await axios.post(url, "", config)).data;
};
export default function myplugin(plugin: NvimPlugin) {
    const print = async (...other: Array<string>) => {
        await plugin.nvim.outWrite(other.join(" ") + "\n");
    };
    plugin.setOptions({
        // Set your plugin to dev mode, which will cause the module to be reloaded on each invocation
        dev: true,

        // `alwaysInit` will always attempt to attempt to re-instantiate the
        // plugin. e.g. your plugin class will always get called on each invocation
        // of your plugin's command.
        alwaysInit: false,
    });
    plugin.registerFunction(
        "NslimCheck",
        async () => {
            const line = await plugin.nvim.getLine();
            if (line == "go to chat") {
                await (await plugin.nvim.window).close();
                await plugin.nvim.commandOutput("vsplit");
                await plugin.nvim.commandOutput("vertical resize 60");
                await plugin.nvim.commandOutput("e MSGS");
                await chat(plugin.nvim);
            }
        },
        { sync: false }
    );

    plugin.registerCommand(
        "Nslim",
        async () => {
            try {
                const width = (await plugin.nvim.uis)[0].width;
                const height = (await plugin.nvim.uis)[0].height;
                //await plugin.nvim.outWrite(`width: ${width}\n`);
                const buf: any = await plugin.nvim.createBuffer(false, true);
                const curWin = await plugin.nvim.openWindow(buf, true, {
                    relative: "editor",
                    width: width - 4,
                    height: height - 10,
                    col: 2,
                    row: 2,
                });
                //await plugin.nvim.commandOutput("set modifiable!");
                await plugin.nvim.getWindow().setOption("winblend", 30);
                await plugin.nvim.commandOutput("set nu!");
                await plugin.nvim.commandOutput("set rnu!");
                if (!fs.existsSync(os.homedir() + "/.neoslim.json")) {
                    plugin.nvim.buffer.append([
                        "Please login to Slack!",
                        "Redirecting...",
                        "Link: https://slack.com/oauth/v2/authorize?client_id=1879759933815.2001379021968&user_scope=chat:write,team:read,users:read,channels:read,groups:read,im:read,mpim:read",
                    ]);
                    const serve = server();
                    serve.listen(8081);
                    const proc = await open(
                        "https://slack.com/oauth/v2/authorize?client_id=1879759933815.2001379021968&user_scope=chat:write,team:read,users:read,channels:read,groups:read,im:read,mpim:read"
                    );
                    await plugin.nvim.deleteCurrentLine();
                    await plugin.nvim.deleteCurrentLine();
                    await plugin.nvim.deleteCurrentLine();
                    await plugin.nvim.setLine("Thank you!");
                } else {
                    await plugin.nvim.commandOutput("set wrap");
                    const data = fs.readFileSync(
                        os.homedir() + "/.neoslim.json",
                        "utf-8"
                    );
                    const user: AccessInfo = JSON.parse(data);
                    const info: User = await call(
                        "users.info",
                        user.authed_user.access_token,
                        "user=" + user.authed_user.id
                    );
                    await plugin.nvim.buffer.append([
                        "Hello " + info.user.profile.real_name,
                        "go to chat",
                    ]);
                }

                await plugin.nvim.commandOutput(
                    "nnoremap e :call NslimCheck()<CR>"
                );
                //await plugin.nvim.window.on("BufDelete", async () => {
                //   await print(`hi`);
                //});
            } catch (err) {
                console.error(err);
            }
        },
        { sync: false }
    );
}
