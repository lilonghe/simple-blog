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

export async function deletePostReq(id:number) {
    return request(`/admin/post/${id}`, {
        method: "DELETE",
    })
}

export async function deleteCateReq(id:number) {
    return request(`/admin/cate/${id}`, {
        method: "DELETE",
    })
}

export async function createOrEditCateReq(params:any) {
    return request(`/admin/cate`, {
        method: "POST",
        data: params,
    })
}

export async function getCateReq(id: number) {
    return request(`/admin/cate/${id}`)
}

export async function deleteTagReq(id:number) {
    return request(`/admin/tag/${id}`, {
        method: "DELETE",
    })
}

export async function createOrEditTagReq(params:any) {
    return request(`/admin/tag`, {
        method: "POST",
        data: params,
    })
}

export async function getTagReq(id: number) {
    return request(`/admin/tag/${id}`)
}

export async function updateOptionsReq(params:any) {
    return request(`/admin/options`, {
        method: "POST",
        data: params,
    })
}

export async function logoutReq() {
    return request(`/admin/logout`, {
        method: "POST",
    })
}

export async function uploadFileReq(params:any) {
    return request(`/admin/upload`, {
        method: 'POST',
        data: params,
        headers: { 'Content-Type': 'multipart/form-data' },
    })
}

export async function getVisitListReq(params: any) {
    return request("/admin/visit/list", { params })
}

export async function getWhisperListReq(params: any) {
    return request("/admin/whisper", { params })
}

export async function createWhisperReq(params: any) {
    return request("/admin/whisper", {
        method: "POST",
        data: params,
    })
}

export async function updateWhisperReq(id: number, params: any) {
    return request(`/admin/whisper/${id}`, {
        method: "PUT",
        data: params,
    })
}

export async function deleteWhisperReq(id: number) {
    return request(`/admin/whisper/${id}`, {
        method: "DELETE",
    })
}

export async function updateWhisperVisibilityReq(id: number, is_public: boolean) {
    return request(`/admin/whisper/${id}/visibility`, {
        method: "PUT",
        data: { is_public },
    })
}


