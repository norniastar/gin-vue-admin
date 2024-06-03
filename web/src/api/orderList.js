import service from "@/utils/request";

export const getOrderList = (params) => {
    return service({
        url: '/order/list',
        method: 'get',
        params
    })
}

export const addOrSaveOrder = (data) => {
    return service({
        url: '/order/add-or-save',
        method: 'post',
        data
    })
}

export const removeOrder = (data) => {
    return service({
        url: '/order/remove',
        method: 'post',
        data
    })
}

