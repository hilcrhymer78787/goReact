<template>
    <div>
        <div class="d-flex">
            <v-text-field v-model="newTask" class="mr-2" @keydown.enter="createTask()" placeholder="newTask" prepend-inner-icon="mdi-check" background-color="white" color="info" outlined dense light clearable></v-text-field>
            <v-btn @click="createTask()" color="info">create</v-btn>
        </div>
        <v-simple-table>
            <thead>
                <tr>
                    <th class="text-left">id</th>
                    <th class="text-left">Name</th>
                    <th></th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="task in tasks" :key="task.id">
                    <td>{{ task.id }}</td>
                    <td>{{ task.name }}</td>
                    <td>
                        <div class="d-flex justify-end">
                            <v-btn color="orange" class="pa-0" style="margin-right:6px;">edit</v-btn>
                            <v-btn @click="deleteTask()" color="error" class="pa-0">delete</v-btn>
                        </div>
                    </td>
                </tr>
            </tbody>
        </v-simple-table>
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