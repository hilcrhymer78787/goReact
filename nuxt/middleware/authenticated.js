import { getAuth, onAuthStateChanged } from "firebase/auth";

export default async ({ store, route, redirect }) => {
    const auth = getAuth();
    onAuthStateChanged(auth, (user) => {
        if (user) {
            if ((route.name == 'login' || route.name == 'register')) {
                redirect({ name: 'index' })
            }
            setTimeout(() => {
                store.commit('setLoginInfo', user)
            }, 500);
        } else {
            if (!(route.name == 'login' || route.name == 'register')) {
                redirect({ name: 'login' })
            }
            setTimeout(() => {
                store.commit('setLoginInfo', false)
            }, 500);
        }
    })
}