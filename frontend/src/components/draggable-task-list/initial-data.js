const initialData = {
    tasks: {
        'task-1': {id: 'task-1', description: 'Some content1'},
        'task-2': {id: 'task-2', description: 'Some content2'},
        'task-3': {id: 'task-3', description: 'Some content3'},
        'task-4': {id: 'task-4', description: 'Some content4'},
        'task-5': {id: 'task-5', description: 'Some content5'},
    },
    columns: {
        'column-1': {
            id: 'column-1',
            title: 'To do',
            taskIds: ['task-1', 'task-2', 'task-3', 'task-4', 'task-5'],
        },
    },
    columnOrder: ['column-1'],
};

export default initialData;