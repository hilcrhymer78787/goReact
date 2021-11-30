<template>
    <div>
        <v-app v-if="loginInfo != null">
            <v-main>
                <v-container>
                    <Nuxt />
                </v-container>
            </v-main>
            <v-bottom-navigation app light fixed color="info">
                <v-btn v-for="(nav,index) in navs" :key="index" :style="`width:calc(100% / ${navs.length}); height:100%;background-color:white;`" :to="nav.to" router light exact>
                    <span>{{nav.ttl}}</span>
                    <v-icon>{{nav.icon}}</v-icon>
                </v-btn>
            </v-bottom-navigation>
        </v-app>
    </div>
</template>

<script>
import { mapState } from "vuex";
export default {
    data() {
        return {
            navs: [
                {
                    ttl: "index",
                    icon: "mdi-playlist-check",
                    to: "/",
                },
                {
                    ttl: "タスク",
                    icon: "mdi-playlist-check",
                    to: "/task",
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