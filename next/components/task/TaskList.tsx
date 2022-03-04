import React, { useState, useEffect } from "react";
import { api } from "@/plugins/axios";
import { AxiosRequestConfig, AxiosResponse, AxiosError } from "axios";
import CreateTask from "@/components/task/CreateTask";
import TaskItem from "@/components/task/TaskItem";
import { apiTaskReadResponseType } from "@/types/api/task/read/response";
import { apiTaskReadResponseTaskType } from "@/types/api/task/read/response";
import AddIcon from "@mui/icons-material/Add";
import {
    Card,
    CardHeader,
    CardContent,
    IconButton,
    Dialog,
    CircularProgress
} from "@mui/material";
type Props = {
}
export default function TaskList (props: Props) {
    const [createTaskDialog, setCreateTaskDialog] = useState(false as boolean);
    const [tasks, setTasks] = useState([] as apiTaskReadResponseTaskType[]);
    const [taskReadLoading, seTtaskReadLoading] = useState(false as boolean);

    const taskRead = () => {
        const requestConfig: AxiosRequestConfig = {
            url: "/task/read",
            method: "GET",
        };
        seTtaskReadLoading(true);
        api(requestConfig)
            .then((res: AxiosResponse<apiTaskReadResponseType>) => {
                setTasks(res.data.tasks);
            })
            .finally(() => {
                seTtaskReadLoading(false);
            });
    };

    useEffect(() => {
        taskRead();
    }, []);

    return (
        <>
            <Card>
                <CardHeader
                    action={
                        <IconButton onClick={() => { setCreateTaskDialog(true); }} component="span">
                            <AddIcon color="primary" />
                        </IconButton>
                    }
                    title="タスク"
                />
                {taskReadLoading && !Boolean(tasks.length) &&
                    <CardContent
                        sx={{
                            display: "flex",
                            justifyContent: "center",
                            p: "30px"
                        }}>
                        <CircularProgress />
                    </CardContent>
                }
                {!taskReadLoading && !Boolean(tasks.length) &&
                    <CardContent
                        sx={{
                            textAlign: "center",
                            p: "20px !important"
                        }}>
                        登録されているタスクはありません
                    </CardContent>
                }
                {Boolean(tasks.length) &&
                    <CardContent sx={{ p: "0 !important" }}>
                        {tasks.map((task: apiTaskReadResponseTaskType, index: number) => (
                            <TaskItem
                                task={task}
                                taskRead={taskRead}
                                key={index.toString()}
                            />
                        ))}
                    </CardContent>
                }
            </Card>

            <Dialog open={createTaskDialog} onClose={() => { setCreateTaskDialog(false); }}>
                {createTaskDialog &&
                    <CreateTask
                        onCloseMyself={() => {
                            setCreateTaskDialog(false);
                            taskRead();
                        }}
                        task={null}
                    />
                }
            </Dialog>
            {/* <pre>{JSON.stringify(tasks, null, 2)}</pre> */}
        </>
    );
}