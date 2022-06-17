import request from "@/utils/request";

export async function loginReq(name: string, password: string) {
    return request("/admin/login", {
        method: "POST",
        data: {
            name,
            password
        }
    });
}

export interface ISystemInfoResponse {
    count: {
        posts: number;
        categories: number;
        tags: number;
    },
    versions: any;
}

export async function getSystemInfoReq() {
    return request("/admin/system");
}

export async function getUserReq() {
    return request("admin/user")
}