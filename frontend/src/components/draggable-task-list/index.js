import initialData from './initial-data';
import React from 'react'
import Column from './column';
import { DragDropContext } from 'react-beautiful-dnd'
import styled from 'styled-components'

const Container = styled.div`
    display: flex;
`;

const Button = styled.button`
  background: "lightblue";
  color: "black";

  font-size: 1em;
  margin: 1em;
  padding: 0.25em 1em;
  border: 1px solid;
  border-radius: 1px;
`;


export default class Table extends React.Component {
    state = initialData;

    addTask = () => {
        const newTask = {
            id: 'task',
            content: 'new content'
        };
        this.state.tasks['task'] = newTask;
        this.state.columns['column-1'].taskIds.push(newTask.id)
        
        this.setState(this.state);

    };

    onDragEnd = result => {
        const { destination, source, draggableId } = result;
        if (!destination) {
            return;
        }
        if (destination.droppableId === source.droppableId &&
            destination.index === source.index) {
            return;
        }
        const column = this.state.columns[source.droppableId];
        const newTaskIds = Array.from(column.taskIds);
        newTaskIds.splice(source.index, 1);
        newTaskIds.splice(destination.index, 0, draggableId);
        const newColumn = {
            ...column,
            taskIds: newTaskIds,
        };

        const newState = {
            ...this.state,
            columns: {
                ...this.columns,
                [newColumn.id]: newColumn,
            }
        }

        this.setState(newState);
    };

    render() {
        return (
            <DragDropContext onDragEnd={this.onDragEnd}>
                <Button onClick={this.addTask}>
                    Add task
                </Button>
                <Container>
                    {this.state.columnOrder.map(columnId => {
                        const column = this.state.columns[columnId];
                        const tasks = column.taskIds.map(taskId => this.state.tasks[taskId]);
                        return <Column key={column.id} column={column} tasks={tasks} />;
                    })}
                </Container>
            </DragDropContext>
        );
    }
}