import axios, { AxiosError } from 'axios'

axios.interceptors.request.use(config => {
    config.baseURL = window.baseUrl
    config.withCredentials = true
    return config
})

axios.interceptors.response.use(response => {
    if (response.data) {
        if (response.data.code) {
            if (response.data.code === "NO_LOGIN") {
                if (window.location.pathname !== '/login') {
                    window.location.href = "/login"
                }
            } else {
                window.$message.error(response.data.msg)
            }
        }
    }
    return response
})

export default function request(url: string, config?: object) {
    let userConfig = {
        url,
        ...config
    };
    return axios.request(userConfig).then(res => {
        return res.data
    }).catch((err: AxiosError) => {
        window.$message.error(err.message)
        return { code: err.code, msg: err.message }
    });
}