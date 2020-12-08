import TodoItem from './TodoItem'

type Props = {
  tasks: Task[],
  toggleTaskCompleted: (task: Task) => Promise<void>,
  updateTask: (task: Task, params: TaskUpdateParams) => Promise<void>,
  deleteTask: (task: Task) => Promise<void>
}

export default function TodoList({
  tasks,
  toggleTaskCompleted,
  updateTask,
  deleteTask
}: Props) {
  return (
    <ul>
      {tasks.map(task => {
        return (
          <li key={task.id}>
            <TodoItem 
              task={task}
              updateTask={updateTask}
              toggleTaskCompleted={toggleTaskCompleted}
              deleteTask={deleteTask} />
          </li>
        )
      })}
    </ul>
  )
}
