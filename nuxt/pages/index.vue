<template>
    <div class="crud">
        <form @submit.prevent="createTodo">
            <input v-model="newtodo" type="text">
            <v-btn @click="createTodo()" color="orange" dark>create</v-btn>
        </form>
        <hr>
        <ul>
            <li class="mb-3">
                <span class="pr-5">id</span>
                <span class="pr-5">name</span>
            </li>
            <li v-for="(task, index) in tasks" :key="index" class="mb-3">
                <span class="pr-5">{{task.id}}</span>
                <span class="pr-5">{{task.name}}</span>
                <v-btn @click="deleteTodo()" color="error" dark>delete</v-btn>
            </li>
        </ul>
    </div>
</template>

<script>
export default {
    data() {
        return {
            newtodo: "",
            tasks: [],
        };
    },
    methods: {
        createTodo() {
            alert("要素が追加されました");
            this.newtodo = "";
        },
        deleteTodo() {
            if (confirm("「" + '内容' + "」を削除しますか？")) {
                alert("要素が削除されました");
            }
        },
    },
    mounted() {
        this.$axios
            .get(`/opening/allread`)
            .then((res) => {
                this.tasks = res.data
            })
            .catch((err) => {
                alert("通信に失敗しました");
            });
    },
};
</script>
<style lang="scss" scoped>
hr {
    margin: 20px 0;
}
li{
    list-style: none;
}
</style>