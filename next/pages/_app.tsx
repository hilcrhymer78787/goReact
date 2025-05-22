import "@/styles/globals.scss";

import { Dispatch, SetStateAction, useEffect, useState } from "react";
import { RecoilRoot } from "recoil";
import { useBearerAuthentication } from "@/data/user/useBearerAuthentication";

import { AdapterDateFns } from "@mui/x-date-pickers/AdapterDateFns";
import type { AppProps } from "next/app";
import CssBaseline from "@mui/material/CssBaseline";
import { LocalizationProvider } from "@mui/x-date-pickers";
import { ThemeProvider } from "@mui/material/styles";
import ja from "date-fns/locale/ja";
import theme from "@/plugins/theme";

const AppInit = ({
  setIsAuth,
}: {
  setIsAuth: Dispatch<SetStateAction<boolean | null>>;
}) => {
  const { bearerAuthentication } = useBearerAuthentication();

  useEffect(() => {
    const mountedFunc = async () => {
      const res = await bearerAuthentication();
      setIsAuth(!res);
    };
    mountedFunc();
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, []);
  
  return <></>;
};

const MyApp = ({ Component, pageProps }: AppProps) => {
  const [isAuth, setIsAuth] = useState<boolean | null>(null);
  return (
    <RecoilRoot>
      <ThemeProvider theme={theme}>
        <LocalizationProvider dateAdapter={AdapterDateFns} adapterLocale={ja}>
          <CssBaseline />
          <AppInit setIsAuth={setIsAuth} />
          {isAuth !== null && <Component {...pageProps} />}
        </LocalizationProvider>
      </ThemeProvider>
    </RecoilRoot>
  );
};

export default MyApp;
