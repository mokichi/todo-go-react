type Props = {
  task: Task,
  toggleTaskCompleted: (task: Task) => void
}

export default function TodoItem({ task, toggleTaskCompleted }: Props) {
  return (
    <div>
      <label>
        <input type="checkbox" checked={task.completed} onChange={() => toggleTaskCompleted(task)} readOnly={true} />
        {task.title}
      </label>
    </div>
  )
}
