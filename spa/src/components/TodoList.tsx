import TodoItem from './TodoItem'

type Props = {
  tasks: Task[],
  updateTask?: (task: Task, params: TaskUpdateParams) => void,
  toggleTaskCompleted: (task: Task) => void,
  deleteTask: (task: Task) => void
}

export default function TodoList({
  tasks,
  updateTask,
  toggleTaskCompleted,
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
