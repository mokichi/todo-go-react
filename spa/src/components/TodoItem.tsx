import React, { useState, useRef } from 'react'

type Props = {
  task: Task,
  toggleTaskCompleted: (task: Task) => Promise<void>,
  updateTask: (task: Task, params: TaskUpdateParams) => Promise<void>,
  deleteTask: (task: Task) => Promise<void>
}

export default function TodoItem({
  task,
  toggleTaskCompleted,
  updateTask,
  deleteTask
}: Props) {
  const [isEditing, setIsEditing] = useState<boolean>(false)

  const inputTitleEl = useRef<HTMLInputElement>(null)
  const handleSubmitTitleForm = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    const title = inputTitleEl.current?.value
    if (!title) {
      return
    }
    await updateTask(task, { title: title })
    setIsEditing(false)
  }
  const TitleForm = (
    <form style={{display: 'inline'}} onSubmit={handleSubmitTitleForm}>
      <input type="text" autoFocus={true} ref={inputTitleEl} defaultValue={task.title} />
    </form>
  )

  return (
    <div>
      <input
        type="checkbox"
        checked={task.completed}
        readOnly={true}
        onChange={() => toggleTaskCompleted(task)} />
      {isEditing ? TitleForm : task.title}
      <small style={{margin: '0 0.5em'}}>
        {new Date(task.updatedAt).toString()}
      </small>
      {isEditing
        ? <button onClick={() => setIsEditing(false)}>Cancel</button>
        : <button onClick={() => setIsEditing(true)}>Edit</button>}
      <button onClick={() => deleteTask(task)}>Delete</button>
    </div>
  )
}
