import { createStore } from "redux";
import { AxiosRequestConfig, AxiosResponse, AxiosError } from "axios";
import { api } from '@/plugins/axios';
import axios from 'axios'
import { apiUserBearerAuthenticationResponseType } from "@/types/api/user/bearerAuthentication/response"
const CancelToken = axios.CancelToken;
let setLoginInfoByTokenCancel: any = null;

const initialState = {
    loginInfo: {
        email: "user1@gmail.com",
        id: 1,
        name: "user1",
        token: "user1@gmail.comsyflysHt4HkYDTVh8DjHEOf9jIwBSZ6IsCJxGmayDNXCshvkDrVB7KxwNLqTLdB3pbTx3lmTNe3Dt4WRkISGgQpZ3tgF5tRmrYVs",
        user_img: "https://picsum.photos/500/300?image=1",
    },
};

const reducer = (state = initialState, action) => {
    switch (action.type) {
        case 'setLoginInfo':
            return { ...state, loginInfo: action.value, };
        default:
            return state;
    }
};
const store = createStore(reducer);

export const bearerAuthentication = async () => {
    if (setLoginInfoByTokenCancel) {
        setLoginInfoByTokenCancel()
    }
    const requestConfig: AxiosRequestConfig = {
        url: `/user/bearer_authentication`,
        method: "GET",
        cancelToken: new CancelToken(c => {
            setLoginInfoByTokenCancel = c
        }),
    };
    api(requestConfig)
        .then((res: AxiosResponse<apiUserBearerAuthenticationResponseType>) => {
            // トークンが有効
            return
            if (localStorage.getItem("token")) {
                store.dispatch({ type: "setLoginInfo", value: res.data })
            }
        })
}
export default store;