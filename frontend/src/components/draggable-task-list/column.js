import React from 'react';
import styled from 'styled-components';
import ListItemTask from './task';
import { Droppable } from 'react-beautiful-dnd';
import { Container, IconButton } from '@material-ui/core';
import { withStyles } from '@material-ui/core/styles';
import Typography from '@material-ui/core/Typography';
import List from '@material-ui/core/List';
import Add from '@material-ui/icons/Add';

const CustomContainer = withStyles({
    root: {
        border: '1px',
        borderStyle: 'ridge',
        borderRadius: '10px',
        display: 'flex', height: '100%',
    }
})(Container);

const Title = withStyles({
    root: {
        width: '100%',
        padding: '8px',
        textAlign: 'center',
    }
})(Typography);

export default class Column extends React.Component {
    render() {
        return (
            <CustomContainer>
                <Title>
                    {this.props.column.title}
                    <IconButton onClick={this.props.addTask}>
                        <Add />
                    </IconButton>
                </Title>
                <Droppable droppableId={this.props.column.id}>
                    {(provided) => (
                        <List
                            ref={provided.innerRef}
                            {...provided.droppableProps}
                        >
                            {this.props.tasks.map((task, index) => <Container>
                                <ListItemTask key={task.id} task={task} index={index} />
                                </Container>)}
                            {provided.placeholder}
                        </List>
                    )}
                </Droppable>
            </CustomContainer>
        );
    }
}