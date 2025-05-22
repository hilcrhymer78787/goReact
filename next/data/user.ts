import axios, { AxiosError, AxiosRequestConfig, AxiosResponse, Canceler } from "axios";

import { api } from "@/plugins/axios";
import { apiUserBasicAuthenticationRequestType } from "@/types/api/user/basicAuthentication/request";
import { apiUserBasicAuthenticationResponseType } from "@/types/api/user/basicAuthentication/response";
import { apiUserBearerAuthenticationResponseType } from "@/types/api/user/bearerAuthentication/response";
import { apiUserCreateResponseType } from "@/types/api/user/create/response";
import { atom } from "recoil";
import { errorType } from "@/types/api/error";
import { useState } from "react";

export const loginInfoAtom = atom<apiUserBearerAuthenticationResponseType | null>({
  key: "loginInfo",
  dangerouslyAllowMutability: true,
  default: null,
});
export const useUserApi = () => {

  // testAuthentication

  const [testAuthenticationLoading, setTestAuthenticationLoading] = useState<boolean>(false);
  const testAuthentication = async (): Promise<AxiosResponse> => {
    const requestConfig: AxiosRequestConfig = {
      url: "/api/user/test_authentication",
      method: "GET",
    };
    setTestAuthenticationLoading(true);
    return api(requestConfig)
      .finally(() => {
        setTestAuthenticationLoading(false);
      });
  };

  // basicAuthentication

  const [basicAuthenticationLoading, setBasicAuthenticationLoading] = useState<boolean>(false);
  const basicAuthentication = async (
    params: apiUserBasicAuthenticationRequestType
  ): Promise<AxiosResponse<apiUserBasicAuthenticationResponseType>> => {
    const requestConfig: AxiosRequestConfig = {
      url: "/api/user/basic_authentication",
      method: "POST",
      data: params
    };
    setBasicAuthenticationLoading(true);
    return api(requestConfig)
      .catch((err: AxiosError<errorType>) => {
        throw err;
      })
      .finally(() => {
        setBasicAuthenticationLoading(false);
      });
  };

  // createUser

  const [createUserLoading, setCreateUserLoading] = useState<boolean>(false);
  const createUser = async (
    params: FormData
  ): Promise<AxiosResponse<apiUserCreateResponseType>> => {
    const requestConfig: AxiosRequestConfig = {
      url: "/api/user/create",
      method: "POST",
      data: params
    };
    setCreateUserLoading(true);
    return api(requestConfig)
      .catch((err: AxiosError<errorType>) => {
        throw err;
      })
      .finally(() => {
        setCreateUserLoading(false);
      });
  };

  return {
    testAuthentication,
    testAuthenticationLoading,
    basicAuthentication,
    basicAuthenticationLoading,
    createUser,
    createUserLoading,
  };
};
