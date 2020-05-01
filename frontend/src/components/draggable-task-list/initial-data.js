const initialData = {
    tasks: {
        'task-1': {id: 'task-1', content: 'Some content1'},
        'task-2': {id: 'task-2', content: 'Some content2'},
        'task-3': {id: 'task-3', content: 'Some content3'},
        'task-4': {id: 'task-4', content: 'Some content4'},
        'task-5': {id: 'task-5', content: 'Some content5'},
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