import { getAuth, onAuthStateChanged } from "firebase/auth";

export default async ({ store, route, redirect }) => {
    const auth = getAuth();
    onAuthStateChanged(auth, (user) => {
        if (user) {
            if ((route.name == 'login' || route.name == 'register')) {
                redirect({ name: 'index' })
            }
            store.commit('setLoginInfo', user)
        } else {
            if (!(route.name == 'login' || route.name == 'register')) {
                redirect({ name: 'login' })
            }
            store.commit('setLoginInfo', false)
        }
    })
}