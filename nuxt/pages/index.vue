<template>
    <div>
        <div class="d-flex">
            <v-text-field v-model="newTask" class="mr-2" @keydown.enter="createTask()" placeholder="newTask" prepend-inner-icon="mdi-check" background-color="white" color="info" outlined dense light clearable></v-text-field>
            <v-btn dark @click="createTask()" color="orange lighten-1">create</v-btn>
        </div>
        <v-card>
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
                                <v-btn dark @click="dialogOpen(task.id,task.name)" color="orange lighten-1" class="pa-0" style="margin-right:6px;">edit</v-btn>
                                <v-btn dark @click="deleteTask(task.id,task.name)" color="error lighten-1" class="pa-0">delete</v-btn>
                            </div>
                        </td>
                    </tr>
                </tbody>
            </v-simple-table>
        </v-card>

        <v-dialog v-model="dialog" width="500">
            <v-card>
                <v-card-title class="text-h5">update task</v-card-title>
                <v-card-text>
                    <div class="d-flex pt-5">
                        <v-text-field v-model="focusTaskName" class="mr-2" @keydown.enter="updateTask()" placeholder="focusTaskName" prepend-inner-icon="mdi-check" background-color="white" color="info" outlined dense light clearable></v-text-field>
                        <v-btn dark @click="updateTask()" color="orange lighten-1">update</v-btn>
                    </div>
                </v-card-text>
                <v-divider></v-divider>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn dark @click="dialog = false">close</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>

    </div>
</template>

<script>
export default {
    data() {
        return {
            dialog: false,
            newTask: "",
            focusTaskId: "",
            focusTaskName: "",
            tasks: [],
        };
    },
    methods: {
        dialogOpen(id, name) {
            this.dialog = true;
            this.focusTaskName = name;
            this.focusTaskId = id;
        },
        createTask() {
            this.$axios
                .post(`/tasks/create?name=${this.newTask}`)
                .then((res) => {
                    this.newTask = "";
                    this.readTask();
                })
                .catch((err) => {
                    alert("通信に失敗しました");
                });
        },
        readTask() {
            this.$axios
                .get(`/tasks/allread`)
                .then((res) => {
                    this.tasks = res.data;
                })
                .catch((err) => {
                    alert("通信に失敗しました");
                });
        },
        updateTask() {
            this.$axios
                .put(
                    `/tasks/update?id=${this.focusTaskId}&name=${this.focusTaskName}`
                )
                .then((res) => {
                    this.readTask();
                    this.dialog = false;
                })
                .catch((err) => {
                    alert("通信に失敗しました");
                });
        },
        deleteTask(id, name) {
            if (confirm(`「${name}」を削除しますか？`)) {
                this.$axios
                    .delete(`/tasks/delete?id=${id}`)
                    .then((res) => {
                        this.readTask();
                    })
                    .catch((err) => {
                        alert("通信に失敗しました");
                    });
            }
        },
    },
    mounted() {
        this.readTask();
    },
};
</script>