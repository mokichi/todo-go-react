type Props = {
  task: Task,
  updateTask?: (task: Task, params: TaskUpdateParams) => void,
  toggleTaskCompleted: (task: Task) => void,
  deleteTask: (task: Task) => void
}

export default function TodoItem({
  task,
  updateTask,
  toggleTaskCompleted,
  deleteTask
}: Props) {
  return (
    <div>
      <label>
        <input
          type="checkbox"
          checked={task.completed}
          readOnly={true}
          onChange={() => toggleTaskCompleted(task)} />
        {task.title}
      </label>
      <button onClick={() => deleteTask(task)}>Delete</button>
    </div>
  )
}
