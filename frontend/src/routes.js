import React from 'react'
import { BrowserRouter, Redirect, Route, Switch } from 'react-router-dom'

import Home from './components/home'
import MainLayout from './components/main-layout'
import Task from './components/task'

export default function Routes() {
    return (
        <BrowserRouter>
            <Switch>
                <Route exact path="/" render={() => (<Redirect to="/home"/>)}></Route>

                <Route exact path="/home">
                    <MainLayout>
                        <Home></Home>
                    </MainLayout>
                </Route>

                <Route exact path="/tasks">
                    <MainLayout>
                        <Task></Task>
                    </MainLayout>
                </Route>
            </Switch>
        </BrowserRouter>
    )
}