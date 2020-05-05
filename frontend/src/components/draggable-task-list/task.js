import React from 'react';
import { Draggable } from 'react-beautiful-dnd';
import ListItem from '@material-ui/core/ListItem';
import ListItemText from '@material-ui/core/ListItemText';

export default class ListItemTask extends React.Component {
    render() {
        return (
            <Draggable draggableId={this.props.task.id} index={this.props.index}>
                {(provided) => (
                    <ListItem button alignItems="flex-start" dense
                        {...provided.draggableProps}
                        {...provided.dragHandleProps}
                        ref={provided.innerRef}
                    >
                        <ListItemText primary={this.props.task.description} />
                    </ListItem>
                )}
            </Draggable>
        );
    }
}