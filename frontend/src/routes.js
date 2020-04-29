import React from 'react'
import { BrowserRouter, Switch, Route } from 'react-router-dom'

import Main from './components/main'
import Task from './components/task'

export default function Routes() {
    return (
        <BrowserRouter>
            <Switch>
                <Route exact path="/">
                    <Main></Main>
                </Route>

                <Route exact path="/tasks">
                    <Task></Task>
                </Route>
            </Switch>
        </BrowserRouter>
    )
}