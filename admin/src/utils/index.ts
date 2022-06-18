import dayjs from "dayjs";

export function formatTime(time: string | Date) {
    return dayjs(time).format('YYYY-MM-DD HH:mm:ss')
}

export function debounce(func: Function, wait = 300) {
    if (window.timeout) clearTimeout(window.timeout)
    window.timeout = setTimeout(() => {
        func()
    }, wait)
}