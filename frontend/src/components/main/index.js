import React from "react";

import { useHistory } from 'react-router-dom'

export default function Main() {
    let history = useHistory();

    const changeRoute = (route) => {
        history.push(route)
    }

    return (
        <button onClick={() => changeRoute("/tasks")}>See tasks</button>
    )
}