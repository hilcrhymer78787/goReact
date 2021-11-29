import {
    getAuth,
    createUserWithEmailAndPassword,
    signInWithEmailAndPassword,
    signInWithRedirect,
    GoogleAuthProvider,
    // TwitterAuthProvider,
    // FacebookAuthProvider,
    signOut
} from "firebase/auth";

export const strict = false

export const state = () => ({
    loginInfo: null,
})

export const mutations = {
    setLoginInfo(state, payload) {
        state.loginInfo = payload
        console.log(payload)
    }
}

export const actions = {
    async signUp({ commit }, { email, password }) {
        const auth = getAuth();
        return createUserWithEmailAndPassword(auth, email, password)
    },

    async signInWithEmail({ commit }, { email, password }) {
        const auth = getAuth();
        return signInWithEmailAndPassword(auth, email, password)
    },

    // signInWithTwitter() {
    //     const auth = getAuth();
    //     return signInWithRedirect(auth, new TwitterAuthProvider())
    // },

    // async signInWithFacebook() {
    //     const auth = getAuth();
    //     return signInWithRedirect(auth, new FacebookAuthProvider())
    // },

    async signInWithGoogle() {
        const auth = getAuth();
        return signInWithRedirect(auth, new GoogleAuthProvider())
    },
    async signOut() {
        const auth = getAuth();
        return signOut(auth)
    }
}