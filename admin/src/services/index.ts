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
    return request("/admin/user")
}

export async function getPostListReq(params: any) {
    return request("/admin/post/list", { params })
}

export async function getOptionsReq() {
    return request("/admin/options")
}

export async function getCateListReq() {
    return request("/admin/cate/list")
}

export async function getTagListReq() {
    return request("/admin/tag/list")
}

export async function createPostReq(params:any) {
    return request("/admin/post", {
        method: "POST",
        data: params,
    })
}

export async function getEditPostReq(id: number) {
    return request(`/admin/post/${id}`)
}