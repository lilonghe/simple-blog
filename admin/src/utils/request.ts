import axios, { AxiosError } from 'axios'
import { getRealPath } from '.'

axios.interceptors.request.use(config => {
    config.baseURL = import.meta.env.VITE_BASE_API_URL || '/api';
    config.withCredentials = true
    return config
})

axios.interceptors.response.use(response => {
    if (response.data) {
        if (response.data.code) {
            if (response.data.code === "NO_LOGIN") {
                if (window.location.pathname !== getRealPath('/login')) {
                    window.location.href = getRealPath("/login")
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