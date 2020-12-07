import TodoItem from './TodoItem'

type Props = {
  tasks: Task[],
  toggleTaskCompleted: (task: Task) => void
}

export default function TodoList({ tasks, toggleTaskCompleted }: Props) {
  return (
    <ul>
      {tasks.map(task => {
        return (
          <li key={task.id}>
            <TodoItem task={task} toggleTaskCompleted={toggleTaskCompleted} />
          </li>
        )
      })}
    </ul>
  )
}
