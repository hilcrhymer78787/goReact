<template>
    <div class="crud">
        <form @submit.prevent="createTask">
            <input v-model="newTask" type="text">
            <v-btn @click="createTask()" color="orange" dark>create</v-btn>
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
                <v-btn @click="deleteTask()" color="error" dark>delete</v-btn>
            </li>
        </ul>
    </div>
</template>

<script>
export default {
    data() {
        return {
            newTask: "",
            tasks: [],
        };
    },
    methods: {
        readTask() {
            this.$axios
                .get(`/opening/allread`)
                .then((res) => {
                    this.tasks = res.data;
                })
                .catch((err) => {
                    alert("通信に失敗しました");
                });
        },
        createTask() {
            alert("要素が追加されました");
            this.newTask = "";
        },
        deleteTask() {
            if (confirm("「" + "内容" + "」を削除しますか？")) {
                alert("要素が削除されました");
            }
        },
    },
    mounted() {
        this.readTask();
    },
};
</script>
<style lang="scss" scoped>
hr {
    margin: 20px 0;
}
li {
    list-style: none;
}
input {
    border: 1px solid orange;
}
</style>