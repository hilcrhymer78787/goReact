<template>
    <v-app>
        <v-app-bar app fixed>
            <div style="width:45px;">
                <v-img v-if="loginInfo" :src="loginInfo.photoURL" class="d-block rounded-circle"></v-img>
            </div>
            <div class="pl-3" v-if="loginInfo">{{loginInfo.displayName}}</div>
            <v-spacer></v-spacer>
            <v-btn @click="$store.dispatch('signOut')">ログアウト</v-btn>
        </v-app-bar>
        <v-main>
            <v-container>
                <Nuxt v-if="$store.state.loginInfo != null" />
            </v-container>
        </v-main>
        <v-bottom-navigation app light fixed color="info">
            <v-btn v-for="(nav,index) in navs" :key="index" :style="`width:calc(100% / ${navs.length}); height:100%;background-color:white;`" :to="nav.to" router light exact>
                <span>{{nav.ttl}}</span>
                <v-icon>{{nav.icon}}</v-icon>
            </v-btn>
        </v-bottom-navigation>
    </v-app>
</template>

<script>
import { mapState } from "vuex";
export default {
    data() {
        return {
            navs: [
                {
                    ttl: "タスク",
                    icon: "mdi-playlist-check",
                    to: "/",
                },
                {
                    ttl: "マイページ",
                    icon: "mdi-account",
                    to: "/mypage",
                },
            ],
        };
    },
    computed: {
        ...mapState(["loginInfo"]),
    },
};
</script>
<style lang="scss">
.v-application {
    // background-color: #e9e9e9 !important;
    ul {
        padding: 0;
    }
    li {
        list-style: none;
    }
}
</style>