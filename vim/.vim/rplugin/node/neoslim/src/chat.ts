import { Neovim } from "neovim";
import { RTMClient } from "@slack/rtm-api";
import fs from "fs";
import os from "os";
import { AccessInfo, call } from ".";

export interface UrlInfo {
    ok: boolean;
    url: string;
}

export const chat = async (nvim: Neovim) => {
    const token: AccessInfo = JSON.parse(
        fs.readFileSync(os.homedir() + "/.neoslim.json", "utf-8")
    );
    const url: UrlInfo = await call(
        "apps.connections.open",
        "xapp-1-A0201B50MUG-1996172247377-6550d273e5926553ca766e71ad108072d3ef81200d38935c8b51650b18c38079"
    );
    nvim.buffer.append(`${url.ok}`);
    if (url.ok) {
        let wssUrl = url.url;
        let socket = new WebSocket(wssUrl+"&debug_reconnects=true");

        socket.onopen = function (e) {
            nvim.buffer.append("CONNECTED\n");
            // connection established
        };

        socket.onmessage = function (event) {
            nvim.outWriteLine(event.data);
        };
    }
};
