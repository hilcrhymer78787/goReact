import React, { useState, useEffect } from "react";
import { api } from "@/plugins/axios";
import { AxiosRequestConfig, AxiosResponse, AxiosError } from "axios";
import { apiTaskReadResponseType } from "@/types/api/task/read/response";
import { apiTaskReadResponseTaskType } from "@/types/api/task/read/response";
import { apiWorkDeleteRequestType } from "@/types/api/work/delete/request";
import { apiWorkCreateRequestType } from "@/types/api/work/create/request";
import CheckBoxOutlineBlankIcon from "@mui/icons-material/CheckBoxOutlineBlank";
import TaskOutlinedIcon from "@mui/icons-material/TaskOutlined";
import CheckBoxIcon from "@mui/icons-material/CheckBox";
import CreateTask from "@/components/task/CreateTask"
import {
    CardActionArea,
    Dialog,
    List,
    ListItem,
    ListItemButton,
    ListItemAvatar,
    ListItemText,
    Avatar,
    CircularProgress
} from "@mui/material";
type Props = {
    task: apiTaskReadResponseTaskType,
    taskRead: any
}
export default function TaskItem(props: Props) {
    const [createTaskDialog, setCreateTaskDialog] = useState(false as boolean);
    return (
        <ListItem sx={{ p: 0 }}>
            <ListItemButton sx={{ p: "8px 48px 8px 16px" }}>
                <ListItemAvatar onClick={() => { setCreateTaskDialog(true); }}>
                    <Avatar sx={{ bgcolor: "#1976d2" }}>
                        <TaskOutlinedIcon />
                    </Avatar>
                </ListItemAvatar>
                <ListItemText
                    onClick={() => { setCreateTaskDialog(true); }}
                    primary={props.task.name}
                    secondary={`想定:${0}分` + ` 実績:${0}分`}
                />
            </ListItemButton>
            <Dialog open={createTaskDialog} onClose={() => { setCreateTaskDialog(false); }}>
                {createTaskDialog &&
                    <CreateTask
                        onCloseMyself={() => {
                            setCreateTaskDialog(false);
                            props.taskRead();
                        }}
                        task={props.task}
                    />
                }
            </Dialog>
        </ListItem>
    );
}
