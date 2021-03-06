import React from 'react'
import { BrowserRouter, Redirect, Route, Switch } from 'react-router-dom'

import Home from './components/home'
import MainLayout from './components/main-layout'
import Table from './components/draggable-task-list'

import './index.scss'

export default function Routes() {
    return (
        <BrowserRouter>
            <Switch>
                <Route exact path="/" render={() => (<Redirect to="/home" />)}></Route>

                <Route exact path="/home">
                    <MainLayout>
                        <Home />
                    </MainLayout>
                </Route>

                <Route exact path="/tasks">
                    <MainLayout>
                        <div className="drag-list">
                            <Table />
                        </div>
                    </MainLayout>
                </Route>
            </Switch>
        </BrowserRouter>
    )
}