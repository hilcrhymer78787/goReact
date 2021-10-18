<template>
    <div>
        <v-card>
            <v-toolbar color="info" dark style="box-shadow:none;">
                <v-toolbar-title>マイページ</v-toolbar-title>
            </v-toolbar>
            <v-card-text v-if="loginInfo" class="d-flex align-center">
                <div class="mx-auto" style="width:30%;">
                    <v-img :src="loginInfo.photoURL" aspect-ratio="1" class="rounded-circle main_img"></v-img>
                </div>
                <v-spacer></v-spacer>
                <div class="pt-2" style="width:65%;">
                    <v-text-field class="mt-5 mb-3" dense color="info" prepend-icon="mdi-account" readonly label="名前" v-model="loginInfo.displayName"></v-text-field>
                    <v-text-field dense color="info" prepend-icon="mdi-email" readonly label="メールアドレス" v-model="loginInfo.email"></v-text-field>
                </div>
            </v-card-text>
        </v-card>
    </div>
</template>

<script>
import { mapState } from "vuex";
export default {
    data() {
        return {};
    },
    computed: {
        ...mapState(["loginInfo"]),
    },
};
</script>

<style>
</style>