import axios from 'axios'

axios.interceptors.request.use(config => {
    config.baseURL = "http://localhost:8080/api"
    return config
})

axios.interceptors.response.use(response => {
    if (response.data) {
        if (response.data.code) {
            window.$message.error(response.data.msg)
        }
    }
    return response
})

export default function request(url: string, config: object | null) {
    let userConfig = {
        url,
        ...config
    };
    return axios.request(userConfig).then(res => {
        return res.data
    });
}