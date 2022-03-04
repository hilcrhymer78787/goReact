import React, { useState, useEffect } from "react";
import { api } from "@/plugins/axios";
import { AxiosRequestConfig, AxiosResponse, AxiosError } from "axios";
import { apiTaskReadResponseType } from "@/types/api/task/read/response";
import { apiTaskReadResponseTaskType } from "@/types/api/task/read/response";
import { apiTaskCreateRequestType } from "@/types/api/task/create/request";
import { apiTaskDeleteRequestType } from "@/types/api/task/delete/request";
import SendIcon from "@material-ui/icons/Send";
import DeleteIcon from "@material-ui/icons/Delete";
import {
    TextField,
    Select,
    MenuItem,
    Box,
    Card,
    CardHeader,
    CardContent,
    CardActionArea,
    CardActions,
    IconButton,
    Dialog,
    ListItem,
    ListItemAvatar,
    ListItemText,
    Avatar,
    CircularProgress
} from "@mui/material";
import { LoadingButton } from "@mui/lab";

type Props = {
    task: apiTaskReadResponseTaskType
    onCloseMyself: any
}
export default function CreateTask (props: Props) {
    const [taskCreateLoading, setTaskCreateLoading] = useState(false as boolean);
    const [taskDeleteLoading, setTaskDeleteLoading] = useState(false as boolean);
    const [formTask, setFormTask] = useState({
        id: 0 as number,
        name: "" as string,
    });
    const [nameError, setNameError] = useState("" as string);
    const taskDelete = () => {
        if (!confirm(`「${props.task.name}」を削除しますか？`)) {
            return;
        }
        const apiParam: apiTaskDeleteRequestType = {
            task_id: formTask.id
        };
        const requestConfig: AxiosRequestConfig = {
            url: `/task/delete?id=${formTask.id}`,
            method: "DELETE",
            data: apiParam
        };
        setTaskDeleteLoading(true);
        api(requestConfig)
            .then((res: AxiosResponse<apiTaskReadResponseType>) => {
                props.onCloseMyself();
            })
            .finally(() => {
                setTaskDeleteLoading(false);
            });
    };
    const taskCreate = () => {
        if (validation()) {
            return;
        }
        const apiParam: apiTaskCreateRequestType = {
            id: Number(formTask.id),
            name: formTask.name,
        };
        const requestConfig: AxiosRequestConfig = {
            url: "/task/create",
            method: "POST",
            data: apiParam
        };
        setTaskCreateLoading(true);
        api(requestConfig)
            .then((res: AxiosResponse<apiTaskReadResponseType>) => {
                props.onCloseMyself();
            })
            .finally(() => {
                setTaskCreateLoading(false);
            });
    };
    const validation = (): boolean => {
        let isError: boolean = false;
        setNameError("");
        if (formTask.name == "") {
            setNameError("タスクの名前は必須です");
            isError = true;
        }
        return isError;
    };
    useEffect(() => {
        if (props.task) {
            setFormTask({
                id: props.task.id,
                name: props.task.name,
            });
        }
    }, []);
    return (
        <Card>
            <CardHeader title={props.task ? props.task.name : "新規タスク登録"} />
            <CardContent>
                <ul>
                    <li>
                        <Box sx={{ mb: "15px" }}>
                            <TextField
                                error={Boolean(nameError)}
                                helperText={nameError}
                                value={formTask.name}
                                onChange={(e) => { setFormTask({ ...formTask, name: e.currentTarget.value }); }}
                                label="タスクの名前" variant="outlined" color="primary" />
                        </Box>
                    </li>
                </ul>
            </CardContent>
            <CardActions>
                <LoadingButton
                    color="error"
                    variant="contained"
                    onClick={taskDelete}
                    loading={taskDeleteLoading}
                    disabled={taskCreateLoading}>
                    削除<DeleteIcon />
                </LoadingButton>
                <LoadingButton
                    color="primary"
                    variant="contained"
                    onClick={taskCreate}
                    loading={taskCreateLoading}
                    disabled={taskDeleteLoading}>
                    登録<SendIcon />
                </LoadingButton>
            </CardActions>
        </Card>
    );
}