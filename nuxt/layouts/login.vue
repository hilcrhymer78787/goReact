<template>
    <v-app>
        <v-main>
            <v-container>
                <Nuxt v-if="$store.state.loginInfo != null" />
            </v-container>
        </v-main>
    </v-app>
</template>

<script>
export default {
    data() {
        return {};
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