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