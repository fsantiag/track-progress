import React from 'react'
import Column from './column';
import { DragDropContext } from 'react-beautiful-dnd'
import { Container } from '@material-ui/core';
import { v4 as uuidv4 } from 'uuid';


// const CustomContainer = withStyles({
//     root: {
//         background: 'linear-gradient(45deg, #FE6B8B 30%, #FF8E53 90%)',
//         height: '100%',
//         position: 'fixed',
//         display: 'flex',
//     }
// })(Container);

export default class Table extends React.Component {
    constructor(props) {
        super(props)
        this.state = {
            tasks: [],
            columns: [
                {
                    id: 'todo',
                    title: 'To do',
                    taskIds: []
                },
                {
                    id: 'in-progress',
                    title: 'In Progress',
                    taskIds: []
                },
                {
                    id: 'done',
                    title: 'Done',
                    taskIds: []
                }
            ]
        };
    }

    updateStateForColumns = (state, columns) => {
        const newColumns = Array.from(this.state.columns)
        columns.forEach(column => {
            const index = state.columns.map(c => c.id).indexOf(column.id);
            newColumns.splice(index, 1, column)
        })
        return {
            ...this.state,
            columns: newColumns,
        }
    };

    addTask = () => {
        const newTask = {
            id: uuidv4(),
            status: 'active',
            title: 'Awesome title',
            description: 'New content'
        };

        const columns = this.state.columns.slice()
        columns[0].taskIds.push(newTask.id)


        this.setState({ tasks: [...this.state.tasks, newTask], columns: columns });

        var params = { Action: 'SendMessage', Version: '2011-10-01', MessageBody: JSON.stringify(newTask) };
        const searchParams = Object.keys(params).map((key) => {
            return encodeURI(key) + '=' + encodeURI(params[key]);
        }).join('&');
        var request = {
            method: 'POST',
            headers: {
                'Content-Type': 'application/x-www-form-urlencoded'
            },
            body: searchParams
        };

        fetch('http://localhost:4576/queue/queue', request)

    };

    onDragEnd = result => {
        const { destination, source, draggableId } = result;
        if (!destination) {
            return;
        }
        if (destination.droppableId === source.droppableId) {
            if (destination.index === source.index) {
                return;
            }
            const column = this.state.columns.filter(column => column.id === source.droppableId)[0];
            const newTaskIds = Array.from(column.taskIds);
            newTaskIds.splice(source.index, 1);
            newTaskIds.splice(destination.index, 0, draggableId);
            const newColumn = {
                ...column,
                taskIds: newTaskIds,
            };

            const newState = this.updateStateForColumns(this.state, [newColumn])

            this.setState(newState);
        } else {
            const sourceColumn = this.state.columns.filter(column => column.id === source.droppableId)[0];
            const destinationColumn = this.state.columns.filter(column => column.id === destination.droppableId)[0];

            const sourceTaskIds = Array.from(sourceColumn.taskIds);
            const destinationTaskIds = Array.from(destinationColumn.taskIds);
            sourceTaskIds.splice(source.index, 1);
            destinationTaskIds.splice(destination.index, 0, draggableId);
            const sourceNewColumn = {
                ...sourceColumn,
                taskIds: sourceTaskIds,
            };
            const destinationNewColumn = {
                ...destinationColumn,
                taskIds: destinationTaskIds,
            };

            const newState = this.updateStateForColumns(this.state, [sourceNewColumn, destinationNewColumn])

            this.setState(newState);
        }

    };

    render() {
        return (
            <DragDropContext onDragEnd={this.onDragEnd}>
                <Container style={{ display: 'flex', position: 'fixed', height: '100%' }}>
                    {this.state.columns.map(column => {
                        const tasks = column.taskIds.map(taskId => {
                            return this.state.tasks.filter(task => task.id === taskId)[0]
                        })
                        return (<Container style={{padding: '2px'}}><Column key={column.id} column={column} tasks={tasks} addTask={this.addTask} /></Container>);
                    })}
                </Container >
            </DragDropContext>
        );
    }
}