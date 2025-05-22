import { api } from "@/plugins/axios";
import { errHandler } from "@/data/common";
import { useState } from "react";
import axios, { Canceler } from "axios";
import { apiUserBearerAuthenticationResponseType } from "@/types/api/user/bearerAuthentication/response";
import { atom, useRecoilState } from "recoil";

const CancelToken = axios.CancelToken;
let setLoginInfoByTokenCancel: Canceler;

const loginInfoAtom = atom<apiUserBearerAuthenticationResponseType | null>({
  key: "loginInfo",
  dangerouslyAllowMutability: true,
  default: null,
});
export const useBearerAuthentication = () => {
  const [loginInfo, setLoginInfo] = useRecoilState(loginInfoAtom); 

  const [bearerAuthenticationLoading, setBearerAuthenticationLoading] =
    useState(false);
  const [bearerAuthenticationError, setBearerAuthenticationError] =
    useState("");
  const bearerAuthentication = async () => {
    setBearerAuthenticationError("");
    setBearerAuthenticationLoading(true);
    const requestConfig = {
      url: "/api/user/bearer_authentication",
      method: "GET",
      cancelToken: new CancelToken((c) => {
        setLoginInfoByTokenCancel = c;
      }),
    };
    return api(requestConfig as any)
      .then((res) => {
        setLoginInfo(res.data);
        return res;
      })
      .catch((err) => {
        errHandler(err, setBearerAuthenticationError);
      })
      .finally(() => {
        setBearerAuthenticationLoading(false);
      });
  };

  return {
    loginInfo,
    bearerAuthentication,
    bearerAuthenticationError,
    bearerAuthenticationLoading,
  };
};
